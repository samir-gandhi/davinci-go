package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetCustomers(companyId, page, limit *string) (*Customers, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/company/%s/customers", c.HostURL, cIdString), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token)
	if err != nil {
		return nil, err
	}

	customers := Customers{}
	err = json.Unmarshal(body, &customers)
	if err != nil {
		return nil, err
	}

	return &customers, nil
}

func (c *Client) GetCustomer(companyId, customerId *string) (*Customer, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer

	if customerId == nil {
		return nil, fmt.Errorf("customerId not provided")
	}

	uIdString := *customerId
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/company/%s/customers/%s", c.HostURL, cIdString, uIdString), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token)
	if err != nil {
		return nil, err
	}

	customer := Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (c *Client) UpdateCustomer(companyId *string, customerId *string, payload *CustomerUpdate) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer

	if customerId == nil {
		return nil, fmt.Errorf("customerId not provided")
	}

	uIdString := *customerId

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/company/%s/customers/%s", c.HostURL, cIdString, uIdString), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token)
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

func (c *Client) CreateCustomer(companyId *string, payload *CustomerCreate) (*Customer, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer

	if payload == nil {
		return nil, fmt.Errorf("payload not provided")
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/company/%s/customers", c.HostURL, cIdString), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token)
	if err != nil {
		return nil, err
	}

	cus := Customer{}
	err = json.Unmarshal(body, &cus)
	if err != nil {
		return nil, err
	}

	return &cus, nil
}

func (c *Client) DeleteCustomer(companyId, customerId *string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer

	if customerId == nil {
		return nil, fmt.Errorf("customerId not provided")
	}

	uIdString := *customerId
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/company/%s/customers/%s", c.HostURL, cIdString, uIdString), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token)
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