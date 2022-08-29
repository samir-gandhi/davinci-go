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

// Gets single connections based on ConnectionId
func (c *Client) ReadConnection(companyId *string, connectionId string) (*Connection, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}

	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	if connectionId == "" {
		return nil, fmt.Errorf("connectionId not provided")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/connections/%s", c.HostURL, connectionId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	connection := Connection{}
	err = json.Unmarshal(body, &connection)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

// Create a bare connection, properties can be added _after_ creation
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
		return nil, fmt.Errorf("Empty or invalid payload")
	}
	connectionCreateBody := Connection{
		Name:        payload.Name,
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

	connResponse := Connection{}
	err = json.Unmarshal(body, &connResponse)
	if err != nil {
		return nil, err
	}

	return &connResponse, nil
}

// Update existing connection properties.
//
/* Sample minimal payload:
&Connection{
	ConnectionID: "foo-123"
	Properties: Properties{
		"foo": struct {
			Value string `json:"value"`
		}{"bar"}
	}
}
*/
func (c *Client) UpdateConnection(companyId *string, payload *Connection) (*Connection, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	if payload == nil || payload.Name == "" || payload.ConnectorID == "" {
		return nil, fmt.Errorf("Empty or invalid payload")
	}

	//Update connection ONLY allows properties
	propsOnly := Connection{
		Properties: payload.Properties,
	}

	reqBody, err := json.Marshal(propsOnly)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/connections/%s", c.HostURL, payload.ConnectionID), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	res := Connection{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Create a connection and fill connection properties
/* Sample minimal payload:
&Connection{
	ConnectorID: "fooConnector"
	Name: "Foo Connector"
	Properties: Properties{
		"foo": struct {
			Value string `json:"value"`
		}{"bar"}
	}
}
*/
func (c *Client) CreateInitializedConnection(companyId *string, payload *Connection) (*Connection, error) {
	connCreatePayload := Connection{
		Name:        payload.Name,
		ConnectorID: payload.ConnectorID,
	}
	resp, err := c.CreateConnection(companyId, &connCreatePayload)
	if err != nil {
		err = fmt.Errorf("Unable to create connection. Error: %v", err)
		return nil, err
	}

	//Update connection ONLY allows properties
	propsOnly := Connection{
		Properties: payload.Properties,
	}

	reqBody, err := json.Marshal(propsOnly)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/connections/%s", c.HostURL, resp.ConnectionID), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	res := Connection{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Deletes a connection based on ConnectionId
func (c *Client) DeleteConnection(companyId *string, connectionId string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/connections/%s", c.HostURL, connectionId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	connection := Message{}
	err = json.Unmarshal(body, &connection)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}
