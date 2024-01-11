package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ReadFlows only accepts Limit as a param
func (c *APIClient) ReadApplications(companyId *string, args *Params) ([]App, error) {
	r, _, err := c.ReadApplicationsWithResponse(companyId, args)
	return r, err
}

func (c *APIClient) ReadApplicationsWithResponse(companyId *string, args *Params) ([]App, *http.Response, error) {
	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/apps", c.HostURL),
	}
	body, res, err := c.doRequestRetryable(companyId, req, args)
	if err != nil {
		return nil, res, err
	}

	// Returned flows are an array in top level flowsInfo key
	resp := Apps{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, res, err
	}
	// Leaving in case of revert - but this shouldn't be treated as an error.
	// if len(resp.Apps) == 0 {
	// 	return nil, fmt.Errorf("No applications found with given params")
	// }
	return resp.Apps, res, nil
}

func (c *APIClient) CreateApplication(companyId *string, appName string) (*App, error) {
	r, _, err := c.CreateApplicationWithResponse(companyId, appName)
	return r, err
}

func (c *APIClient) CreateApplicationWithResponse(companyId *string, appName string) (*App, *http.Response, error) {
	if appName == "" {
		return nil, nil, fmt.Errorf("Must provide appName")
	}

	p := App{
		Name: appName,
	}

	payload, err := json.Marshal(p)
	if err != nil {
		return nil, nil, err
	}

	req := DvHttpRequest{
		Method: "POST",
		Url:    fmt.Sprintf("%s/apps", c.HostURL),
		Body:   strings.NewReader(string(payload)),
	}

	body, res, err := c.doRequestRetryable(companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	r := ReadApp{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, res, err
	}

	if r.App.Name == "" {
		return nil, res, fmt.Errorf("Unable to create app")
	}
	return &r.App, res, nil
}

// UpdateApplication - Update all fields of an application besides Policies. Policies should be updated via UpdatePolicy
func (c *APIClient) UpdateApplication(companyId *string, payload *AppUpdate) (*App, error) {
	r, _, err := c.UpdateApplicationWithResponse(companyId, payload)
	return r, err
}

func (c *APIClient) UpdateApplicationWithResponse(companyId *string, payload *AppUpdate) (*App, *http.Response, error) {
	if payload == nil || payload.Name == "" || payload.AppID == "" {
		return nil, nil, fmt.Errorf("App Name and ID required in payload")
	}

	appId := payload.AppID
	payloadFormatted := *payload
	payloadFormatted.AppID = ""

	reqBody, err := json.Marshal(payloadFormatted)
	if err != nil {
		return nil, nil, err
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/apps/%s", c.HostURL, appId),
		Body:   strings.NewReader(string(reqBody)),
	}

	body, res, err := c.doRequestRetryable(companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	appRes := ReadApp{}
	err = json.Unmarshal(body, &appRes)
	if err != nil {
		return nil, res, err
	}

	return &appRes.App, res, nil
}

func (c *APIClient) ReadApplication(companyId *string, appId string) (*App, error) {
	r, _, err := c.ReadApplicationWithResponse(companyId, appId)
	return r, err
}

func (c *APIClient) ReadApplicationWithResponse(companyId *string, appId string) (*App, *http.Response, error) {
	if appId == "" {
		return nil, nil, fmt.Errorf("AppId not provided")
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/apps/%s", c.HostURL, appId),
	}

	body, res, err := c.doRequestRetryable(companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	appRes := ReadApp{}
	err = json.Unmarshal(body, &appRes)
	if err != nil {
		return nil, res, err
	}

	return &appRes.App, res, nil
}

// CreateInitializedApplication is useful when creating an application with flow policy.
// Takes an app payload and calls:
// - CreateApplication
// - UpdateApplication
// - CreateFlowPolicy
func (c *APIClient) CreateInitializedApplication(companyId *string, payload *AppUpdate) (*App, error) {
	r, _, err := c.CreateInitializedApplicationWithResponse(companyId, payload)
	return r, err
}

func (c *APIClient) CreateInitializedApplicationWithResponse(companyId *string, payload *AppUpdate) (*App, *http.Response, error) {

	//Create Application
	resp, res, err := c.CreateApplicationWithResponse(companyId, payload.Name)
	if err != nil {
		return nil, res, err
	}

	//Remove Policies from initial update payload as they must be created separately
	policies := payload.Policies
	payload.Policies = nil

	//Merge computed fields from Created application with update payload.	(e.g. client secret)
	if payload.Oauth == nil {
		payload.Oauth = resp.Oauth
	}
	if payload.Saml == nil {
		payload.Saml = resp.Saml
	}
	payload.Oauth.Values.ClientSecret = resp.Oauth.Values.ClientSecret

	payload.AppID = resp.AppID

	//Update Application
	resp, res, err = c.UpdateApplicationWithResponse(companyId, payload)
	if err != nil {
		return nil, res, err
	}

	//Create Flow Policies if exist
	if len(policies) != 0 {

		for _, v := range policies {
			_, res, err := c.CreateFlowPolicyWithResponse(companyId, resp.AppID, v)
			if err != nil {
				return nil, res, err
			}
			polRead, res, err := c.ReadApplicationWithResponse(companyId, resp.AppID)
			if err != nil {
				return nil, res, err
			}
			payload.Policies = polRead.Policies
		}

		//Update Application with final payload
		resp, res, err = c.UpdateApplicationWithResponse(companyId, payload)
		if err != nil {
			return nil, res, err
		}
	}

	return resp, res, nil
}

// Deletes an application based on applicationId
func (c *APIClient) DeleteApplication(companyId *string, appId string) (*Message, error) {
	r, _, err := c.DeleteApplicationWithResponse(companyId, appId)
	return r, err
}

func (c *APIClient) DeleteApplicationWithResponse(companyId *string, appId string) (*Message, *http.Response, error) {
	req := DvHttpRequest{
		Method: "DELETE",
		Url:    fmt.Sprintf("%s/apps/%s", c.HostURL, appId),
	}

	body, res, err := c.doRequestRetryable(companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	resp := Message{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, res, err
	}

	return &resp, res, nil
}
