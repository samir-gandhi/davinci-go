package davinci

import (
	"encoding/json"

	// "errors"
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func (c *APIClient) SignInSSO(targetEnvironmentID *string) (*AuthResponse, error) {
	r, _, err := c.SignInSSOWithResponse(targetEnvironmentID)
	return r, err
}

func (c *APIClient) SignInSSOWithResponse(targetEnvironmentID *string) (*AuthResponse, *http.Response, error) {
	return c.doAuthRequestRetryable(targetEnvironmentID)
}

func (c *APIClient) doAuthRequestRetryable(targetEnvironmentID *string) (*AuthResponse, *http.Response, error) {
	body, res, err := exponentialBackOffRetry(func() (any, *http.Response, error) {
		return c.doAuthRequest(targetEnvironmentID)
	})
	if err != nil {
		return nil, res, err
	}

	var returnVar *AuthResponse = nil
	var ok bool
	if returnVar, ok = body.(*AuthResponse); !ok {
		return nil, res, fmt.Errorf("Unable to cast variable type to response from Davinci API for variable with name: %s", returnVar)
	}

	return returnVar, res, nil
}

func (c *APIClient) doAuthRequest(targetEnvironmentID *string) (*AuthResponse, *http.Response, error) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	// For Prod an accessToken is aquired by providing an authToken takes multiple steps:
	// 1. Generate SSO Url and refresh state (a)
	// 2. Authorize - provides get code or FlowId (b)
	// 3a. Got FlowId, Log in with admin (c)
	// 3b. Refresh FlowId for Code (d)
	// 4. Send SSO code or state for callback to get authToken (e)
	// 5. Use authToken for DV accessToken (f)

	// Login
	var dvSsoCode, dvFlowId, dvSsoState, dvSsoAuthToken string
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, nil, fmt.Errorf("define username and password")
	}
	if c.PingOneSSOEnvId == "" {
		return nil, nil, fmt.Errorf("define PingOne Admin and Target EnvId")
	}

	// step 1 - Start SSO, refresh state and generate callback
	areq, err := http.NewRequest("GET", fmt.Sprintf("%s/customers/pingone/sso", c.HostURL), nil)
	if err != nil {
		return nil, nil, err
	}

	paramsMap := map[string]string{
		"env": c.PingOneSSOEnvId,
	}

	if targetEnvironmentID != nil && *targetEnvironmentID != "" {
		paramsMap["target"] = *targetEnvironmentID
	}

	aParams := Params{
		"", "", paramsMap,
	}
	ares, res, err := c.doRequestVerbose(areq, nil, &aParams)
	if err != nil || ares.StatusCode != 302 {
		return nil, res, fmt.Errorf("Error getting SSO callback, got err: %v\n", err)
	}
	if ares.StatusCode != 302 {
		return nil, res, fmt.Errorf("Error getting SSO callback, got err: %v", string(ares.Body))
	}
	if ares.LocationParams.Get("state") == "" {
		return nil, res, fmt.Errorf("Error Parsing SSO State not found, got: %s", ares.Location)
	}
	dvSsoState = ares.LocationParams["state"][0]

	//step 2 - Directly Execute Callback from step 1
	// Receive cookies and code or flowid in response
	breq, err := http.NewRequest("GET", fmt.Sprintf("%s://%s%s", ares.Location.Scheme, ares.Location.Host, ares.Location.Path), nil)
	if err != nil {
		return nil, nil, err
	}
	bParams := Params{
		"", "", map[string]string{},
	}
	for i, v := range ares.LocationParams {
		bParams.ExtraParams[i] = v[0]
	}

	bres, res, err := c.doRequestVerbose(breq, nil, &bParams)
	if err != nil {
		return nil, res, fmt.Errorf("Error following SSO callback, got error: %v\n", err)
	}
	if bres.StatusCode != 302 {
		return nil, res, fmt.Errorf("Error following SSO callback, got: %v\n", string(bres.Body))
	}
	if bres.LocationParams.Get("flowId") != "" {
		dvFlowId = bres.LocationParams["flowId"][0]
	}
	if bres.LocationParams.Get("code") != "" {
		dvSsoCode = bres.LocationParams["code"][0]
	}
	if dvFlowId == "" && dvSsoCode == "" {
		return nil, res, fmt.Errorf("Error: SSO Location header did not provide Code or FlowId: %s", bres.Location)
	}

	if dvFlowId != "" {
		// step 3 Refresh FlowID to retrieve dvSsoCode

		//step 3a Log in Admin
		// Assumption that this refreshes backend SSO state..
		crb := map[string]string{
			"username": c.Auth.Username,
			"password": c.Auth.Password}
		cReqBody, err := json.Marshal(crb)
		creq, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/%s/flows/%s", ares.Location.Scheme, ares.Location.Host, c.PingOneSSOEnvId, dvFlowId), bytes.NewBuffer(cReqBody))
		if err != nil {
			return nil, nil, err
		}
		// PingOne Auth Specific Header
		creq.Header.Set("Content-Type", "application/vnd.pingidentity.usernamePassword.check+json; charset=UTF-8")

		cres, res, err := c.doRequestVerbose(creq, nil, nil)
		if err != nil {
			return nil, res, fmt.Errorf("Error Authenticating PingOne Admin: %v", err)
		}
		cResBody := SSOAuthenticationResponse{}
		json.Unmarshal(cres.Body, &cResBody)
		if cResBody.Status != "COMPLETED" {
			return nil, res, fmt.Errorf("Authentication during SSO failed with result: %v", string(cres.Body))
		}
		//step 3b Retrieve dvSsoCode with refreshed Auth
		dreq, err := http.NewRequest("GET", fmt.Sprintf("%s://%s/%s/as/resume", ares.Location.Scheme, ares.Location.Host, c.PingOneSSOEnvId), nil)
		if err != nil {
			return nil, res, err
		}
		dParams := Params{
			"", "", map[string]string{
				"flowId": dvFlowId,
			},
		}
		dres, res, err := c.doRequestVerbose(dreq, nil, &dParams)
		if err != nil {
			return nil, res, fmt.Errorf("Error resuming auth, got error: %v\n", err)
		}
		if dres.StatusCode != 302 {
			return nil, res, fmt.Errorf("Error resuming auth, got: %v\n", string(dres.Body))
		}
		if dres.LocationParams.Get("code") == "" {
			return nil, res, fmt.Errorf("Error Parsing SSO Location, dvSsoCode not found: %v", dres.Location)
		}
		dvSsoCode = dres.LocationParams["code"][0]
	}
	//step 4 use dvSsoCode and dvSsoState to get dvAuthToken
	ereq, err := http.NewRequest("GET", fmt.Sprintf("%s/customers/pingone/callback", c.HostURL), nil)
	if err != nil {
		return nil, nil, err
	}
	eParams := Params{
		"", "", map[string]string{},
	}
	if dvSsoCode != "" {
		eParams.ExtraParams["code"] = dvSsoCode
	}
	if dvSsoState != "" {
		eParams.ExtraParams["state"] = dvSsoState
	}
	eres, res, err := c.doRequestVerbose(ereq, nil, &eParams)
	if err != nil {
		return nil, res, fmt.Errorf("Error getting admin callback, got: %v\n", err)
	}
	if eres.StatusCode != 302 {
		return nil, res, fmt.Errorf("Error getting admin callback, got: %v\n", string(eres.Body))
	}
	if eres.LocationParams.Get("authToken") == "" {
		return nil, res, fmt.Errorf("Auth Token not found, unsuccessful login, got: %v", string(eres.Body))
	}
	dvSsoAuthToken = eres.LocationParams["authToken"][0]

	//step 5 Swap dvSsoAuthToken for access_token
	frb := map[string]string{
		"authToken": dvSsoAuthToken}
	fReqBody, err := json.Marshal(frb)

	freq, err := http.NewRequest("POST", fmt.Sprintf("%s/customers/sso/auth", c.HostURL), strings.NewReader(string(fReqBody)))
	if err != nil {
		return nil, nil, err
	}
	fres, res, err := c.doRequestVerbose(freq, nil, nil)
	if err != nil {
		return nil, res, fmt.Errorf("Error getting admin callback, got: %v", err)
	}

	ar := AuthResponse{}
	err = json.Unmarshal(fres.Body, &ar)
	if err != nil {
		return nil, res, err
	}

	c.CompanyID = "newauth"

	return &ar, res, nil
}
