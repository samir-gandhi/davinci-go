package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strings"
)

// Returns list of Environments (auth required)
func (c *APIClient) ReadEnvironments() (*Environments, error) {
	r, _, err := c.ReadEnvironmentsWithResponse()
	return r, err
}

func (c *APIClient) ReadEnvironmentsWithResponse() (*Environments, *http.Response, error) {
	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/customers/me", c.HostURL),
	}
	body, res, err := c.doRequestRetryable(nil, req, nil)
	if err != nil {
		return nil, res, err
	}

	environments := Environments{}
	err = json.Unmarshal(body, &environments)
	if err != nil {
		return nil, res, err
	}

	return &environments, res, nil
}

func (c *APIClient) ReadEnvironment(companyId string) (*Environment, error) {
	r, _, err := c.ReadEnvironmentWithResponse(companyId)
	return r, err
}

func (c *APIClient) ReadEnvironmentWithResponse(companyId string) (*Environment, *http.Response, error) {

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/company/%s", c.HostURL, companyId),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	environment := Environment{}
	err = json.Unmarshal(body, &environment)
	if err != nil {
		return nil, res, err
	}

	return &environment, res, nil
}

func (c *APIClient) ReadEnvironmentstats(companyId string) (*EnvironmentStats, error) {
	r, _, err := c.ReadEnvironmentstatsWithResponse(companyId)
	return r, err
}

func (c *APIClient) ReadEnvironmentstatsWithResponse(companyId string) (*EnvironmentStats, *http.Response, error) {

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/company/%s/stats", c.HostURL, companyId),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	environment := EnvironmentStats{}
	err = json.Unmarshal(body, &environment)
	if err != nil {
		return nil, res, err
	}

	return &environment, res, nil
}

func (c *APIClient) SetEnvironment(companyId string) (*Message, error) {
	r, _, err := c.SetEnvironmentWithResponse(companyId)
	return r, err
}

func (c *APIClient) SetEnvironmentWithResponse(companyId string) (*Message, *http.Response, error) {

	if companyId == "" {
		return nil, nil, fmt.Errorf("companyId not provided")
	}

	// req.Close = true
	body, res, err := c.exponentialBackOffRetry(func() (any, *http.Response, error) {
		return c.doRequest(DvHttpRequest{
			Method: "PUT",
			Url:    fmt.Sprintf("%s/company/%s/switch", c.HostURL, companyId),
		}, nil)
	}, false)
	if err != nil {
		return nil, res, err
	}

	msg := Message{}
	err = json.Unmarshal(body.([]byte), &msg)
	if err != nil {
		return nil, res, err
	}

	c.companyID = companyId

	return &msg, res, nil
}
