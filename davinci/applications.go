package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ReadFlows only accepts Limit as a param
func (c *Client) ReadApplications(companyId *string, args *Params) ([]App, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apps", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, args)
	if err != nil {
		return nil, err
	}

	// Returned flows are an array in top level flowsInfo key
	resp := Apps{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	if len(resp.Apps) == 0 {
		return nil, fmt.Errorf("No applications found with given params")
	}
	return resp.Apps, nil
}

func (c *Client) CreateApplication(companyId *string, appName string) (*App, error) {
	if appName == "" {
		return nil, fmt.Errorf("Must provide appName")
	}
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}

	_, err := c.SetEnvironment(cIdPointer)

	if err != nil {
		return nil, err
	}

	p := App{
		Name: appName,
	}

	payload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apps", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	r := CreatedApp{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	if r.App.Name == "" {
		return nil, fmt.Errorf("Unable to create app")
	}
	return &r.App, nil
}

func (c *Client) UpdateApplication(companyId *string, payload *AppUpdate) (*App, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	if payload == nil || payload.Name == "" || payload.AppID == "" {
		return nil, fmt.Errorf("App Name and ID required in payload")
	}
	appId := payload.AppID
	payload.AppID = ""

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apps/%s", c.HostURL, appId), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	res := App{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadApplication(companyId *string, appId string) (*App, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}

	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	if appId == "" {
		return nil, fmt.Errorf("AppId not provided")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apps/%s", c.HostURL, appId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	resp := App{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CreateInitializedApplication(companyId *string, payload *AppUpdate) (*App, error) {
	resp, err := c.CreateApplication(companyId, payload.Name)
	if err != nil {
		err = fmt.Errorf("Unable to create connection. Error: %v", err)
		return nil, err
	}
	payload.AppID = resp.AppID
	resp, err = c.UpdateApplication(companyId, payload)
	if err != nil {
		err = fmt.Errorf("Unable to create connection. Error: %v", err)
		return nil, err
	}
	return resp, nil
}
