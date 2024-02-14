package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *APIClient) CreateFlowPolicy(companyId string, appId string, policy Policy) (*App, error) {
	r, _, err := c.CreateFlowPolicyWithResponse(companyId, appId, policy)
	return r, err
}

func (c *APIClient) CreateFlowPolicyWithResponse(companyId string, appId string, policy Policy) (*App, *http.Response, error) {
	if appId == "" {
		return nil, nil, fmt.Errorf("Must provide appName")
	}

	payload := policy
	payload.PolicyID = nil
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	req := DvHttpRequest{
		Method: "POST",
		Url:    fmt.Sprintf("%s/apps/%s/policy", c.HostURL, appId),
		Body:   strings.NewReader(string(reqBody)),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	r := App{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, res, err
	}

	if len(r.Policies) == 0 {
		return nil, res, fmt.Errorf("Unable to create FlowPolicy")
	}

	if *r.CompanyID != companyId {
		return nil, res, fmt.Errorf("Application flow policy created with wrong companyId")
	}

	return &r, res, nil
}

func (c *APIClient) UpdateFlowPolicy(companyId string, appId string, policy Policy) (*App, error) {
	r, _, err := c.UpdateFlowPolicyWithResponse(companyId, appId, policy)
	return r, err
}

func (c *APIClient) UpdateFlowPolicyWithResponse(companyId string, appId string, policy Policy) (*App, *http.Response, error) {
	if appId == "" || policy.PolicyID == nil || *policy.PolicyID == "" {
		return nil, nil, fmt.Errorf("Missing appId or policy.PolicyID")
	}
	payload := policy
	payload.PolicyID = nil
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/apps/%s/policy/%v", c.HostURL, appId, *policy.PolicyID),
		Body:   strings.NewReader(string(reqBody)),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	application := App{}
	err = json.Unmarshal(body, &application)
	if err != nil {
		return nil, res, err
	}

	return &application, res, nil
}

// Deletes an application based on applicationId
func (c *APIClient) DeleteFlowPolicy(companyId string, appId string, policyId string) (*Message, error) {
	r, _, err := c.DeleteFlowPolicyWithResponse(companyId, appId, policyId)
	return r, err
}

func (c *APIClient) DeleteFlowPolicyWithResponse(companyId string, appId string, policyId string) (*Message, *http.Response, error) {
	req := DvHttpRequest{
		Method: "DELETE",
		Url:    fmt.Sprintf("%s/apps/%s/policy/%s", c.HostURL, appId, policyId),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
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
