package davinci

import (
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"
	"strings"
)

// // Sign up - Create new user, return user token upon successful creation
// func (c *Client) SignUp(auth AuthStruct) (*AuthResponse, error) {
// 	if auth.Username == "" || auth.Password == "" {
// 		return nil, fmt.Errorf("define username and password")
// 	}
// 	rb, err := json.Marshal(auth)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", c.HostURL), strings.NewReader(string(rb)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ar := AuthResponse{}
// 	err = json.Unmarshal(body, &ar)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ar, nil
// }

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	// For Prod Getting an Access Token takes multiple steps:
	// 1. Login with User/PW - get access_token
	// 2. Start Auth Flow - Get json response
	// 3. Post response to skCallback

	// Login
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("define username and password")
	}
	lReqBody, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}

	lreq, err := http.NewRequest("POST", fmt.Sprintf("%s/customers/login", c.HostURL), strings.NewReader(string(lReqBody)))
	if err != nil {
		return nil, err
	}

	lbody, err := c.doRequest(lreq, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("Error on User Login, got: %v", err)
	}

	lr := LoginResponse{}
	err = json.Unmarshal(lbody, &lr)
	if err != nil {
		return nil, err
	}

	// Start Auth
	var sreq *http.Request
	if c.HostURL == "https://orchestrate-api.pingone.com/v1" {
		sreq, err = http.NewRequest("POST", fmt.Sprintf("https://auth.pingone.com/%s/davinci/policy/%s/start", lr.CompanyID, lr.FlowPolicyID), nil)
	} else {
		sreq, err = http.NewRequest("POST", fmt.Sprintf("%s/auth/%s/policy/%s/start", c.HostURL, lr.CompanyID, lr.FlowPolicyID), nil)
	}
	if err != nil {
		return nil, err
	}

	sbody, err := c.doRequest(sreq, &lr.SkSdkToken.AccessToken, nil)
	if err != nil {
		return nil, fmt.Errorf("Error on Start Auth, got: %v", err)
	}

	sr := Callback{}
	err = json.Unmarshal(sbody, &sr)
	if err != nil {
		return nil, err
	}

	// Callback
	cReqBody, err := json.Marshal(sr)
	if err != nil {
		return nil, err
	}
	areq, err := http.NewRequest("POST", fmt.Sprintf("%s/customers/skcallback", c.HostURL), strings.NewReader(string(cReqBody)))
	if err != nil {
		return nil, err
	}
	abody, err := c.doRequest(areq, &lr.AccessToken, nil)
	if err != nil {
		return nil, fmt.Errorf("Error on Callback, got: %v", err)
	}

	ar := AuthResponse{}
	err = json.Unmarshal(abody, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

// SignIn - Get a new token for user
// func (c *Client) GetUserTokenSignIn(auth AuthStruct) (*AuthResponse, error) {
// 	if auth.Username == "" || auth.Password == "" {
// 		return nil, fmt.Errorf("define username and password")
// 	}
// 	rb, err := json.Marshal(auth)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/customers/login", c.HostURL), strings.NewReader(string(rb)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, nil)
// 	if err != nil {
// 		return nil, errors.New("Unable to login")
// 	}

// 	ar := AuthResponse{}
// 	err = json.Unmarshal(body, &ar)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ar, nil
// }

// // SignOut - Revoke the token for a user
// func (c *Client) SignOut(authToken *string) error {
// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signout", c.HostURL), strings.NewReader(string("")))
// 	if err != nil {
// 		return err
// 	}

// 	body, err := c.doRequest(req, authToken)
// 	if err != nil {
// 		return err
// 	}

// 	if string(body) != "Signed out user" {
// 		return errors.New(string(body))
// 	}

// 	return nil
// }
