package davinci

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Gets array of all connections for the provided company
func (c *Client) ReadConnections(companyId *string, args *Params) ([]Connection, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	msg, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Set Env tp: %s\n", msg.Message)

	cIdString := *cIdPointer
	log.Print(cIdString)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/connections", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, args)
	if err != nil {
		return nil, err
	}

	connections := []Connection{}
	err = json.Unmarshal(body, &connections)
	if err != nil {
		return nil, err
	}

	return connections, nil
}

// Create a bare role, policies can be added _after_ creation
func (c *Client) CreateConnection(companyId *string, payload *Connection) (*Connection, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	msg, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Set Env tp: %s\n", msg.Message)

	if payload == nil || payload.Name == "" || payload.ConnectorID == "" {
		return nil, fmt.Errorf("empty or invalid payload")
	}
	connectionCreateBody := Connection{
		Name: payload.Name,
		ConnectorID: payload.ConnectorID,
	}
	reqBody, err := json.Marshal(connectionCreateBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/connections", c.HostURL), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	conectionResponse := Connection{}
	err = json.Unmarshal(body, &conectionResponse)
	if err != nil {
		return nil, err
	}

	return &conectionResponse, nil
}