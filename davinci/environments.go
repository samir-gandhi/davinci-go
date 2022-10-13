package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strings"
)

// Returns list of Environments (auth required)
func (c *Client) ReadEnvironments() (*Environments, error) {
	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/customers/me", c.HostURL),
	}
	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	environments := Environments{}
	err = json.Unmarshal(body, &environments)
	if err != nil {
		return nil, err
	}

	return &environments, nil
}

func (c *Client) ReadEnvironment(companyId *string) (*Environment, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer
	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/company/%s", c.HostURL, cIdString),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	environment := Environment{}
	err = json.Unmarshal(body, &environment)
	if err != nil {
		return nil, err
	}

	return &environment, nil
}

func (c *Client) ReadEnvironmentstats(companyId *string) (*EnvironmentStats, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/company/%s/stats", c.HostURL, cIdString),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	environment := EnvironmentStats{}
	err = json.Unmarshal(body, &environment)
	if err != nil {
		return nil, err
	}

	return &environment, nil
}

func (c *Client) SetEnvironment(companyId *string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/company/%s/switch", c.HostURL, cIdString), nil)
	if err != nil {
		return nil, err
	}
	// req.Close = true
	body, _, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	msg := Message{}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
