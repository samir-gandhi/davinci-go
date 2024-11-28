package davinci

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

// Gets array of all connections for the provided company
func (c *APIClient) ReadConnections(companyId string, args *Params) ([]Connection, error) {
	r, _, err := c.ReadConnectionsWithResponse(companyId, args)
	return r, err
}

func (c *APIClient) ReadConnectionsWithResponse(companyId string, args *Params) ([]Connection, *http.Response, error) {
	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/connections", c.HostURL),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, args)
	if err != nil {
		return nil, res, err
	}

	connections := []Connection{}
	err = json.Unmarshal(body, &connections)
	if err != nil {
		return nil, nil, err
	}

	return connections, res, nil
}

// Gets single connections based on ConnectionId
func (c *APIClient) ReadConnection(companyId string, connectionId string) (*Connection, error) {
	r, _, err := c.ReadConnectionWithResponse(companyId, connectionId)
	return r, err
}

func (c *APIClient) ReadConnectionWithResponse(companyId string, connectionId string) (*Connection, *http.Response, error) {
	if connectionId == "" {
		return nil, nil, fmt.Errorf("connectionId not provided")
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/connections/%s", c.HostURL, connectionId),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	connection := Connection{}
	err = json.Unmarshal(body, &connection)
	if err != nil {
		return nil, res, err
	}

	return &connection, res, nil
}

// Create a bare connection, properties can be added _after_ creation
func (c *APIClient) CreateConnection(companyId string, payload *Connection) (*Connection, error) {
	r, _, err := c.CreateConnectionWithResponse(companyId, payload)
	return r, err
}

func (c *APIClient) CreateConnectionWithResponse(companyId string, payload *Connection) (*Connection, *http.Response, error) {
	if payload == nil || payload.Name == nil || *payload.Name == "" || payload.ConnectorID == nil || *payload.ConnectorID == "" {
		return nil, nil, fmt.Errorf("Empty or invalid payload")
	}
	connectionCreateBody := Connection{
		Name:        payload.Name,
		ConnectorID: payload.ConnectorID,
	}
	reqBody, err := json.Marshal(connectionCreateBody)
	if err != nil {
		return nil, nil, err
	}
	req := DvHttpRequest{
		Method: "POST",
		Url:    fmt.Sprintf("%s/connections", c.HostURL),
		Body:   string(reqBody),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	connResponse := Connection{}
	err = json.Unmarshal(body, &connResponse)
	if err != nil {
		return nil, res, err
	}

	if *connResponse.CompanyID != companyId {
		slog.Error("Connection created with wrong companyId", "Company ID (Response)", *connResponse.CompanyID, "Company ID (Intended)", companyId)
		return nil, res, fmt.Errorf("Connection created with wrong companyId")
	}

	return &connResponse, res, nil
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
func (c *APIClient) UpdateConnection(companyId string, payload *Connection) (*Connection, error) {
	r, _, err := c.UpdateConnectionWithResponse(companyId, payload)
	return r, err
}

func (c *APIClient) UpdateConnectionWithResponse(companyId string, payload *Connection) (*Connection, *http.Response, error) {
	if payload == nil || payload.Name == nil || *payload.Name == "" || payload.ConnectorID == nil || *payload.ConnectorID == "" {
		return nil, nil, fmt.Errorf("Empty or invalid payload")
	}

	//Update connection ONLY allows properties
	propsOnly := Connection{
		Properties: payload.Properties,
	}

	reqBody, err := json.Marshal(propsOnly)
	if err != nil {
		return nil, nil, err
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/connections/%v", c.HostURL, *payload.ConnectionID),
		Body:   string(reqBody),
	}
	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	connection := Connection{}
	err = json.Unmarshal(body, &connection)
	if err != nil {
		return nil, res, err
	}

	return &connection, res, nil
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
func (c *APIClient) CreateInitializedConnection(companyId string, payload *Connection) (*Connection, error) {
	r, _, err := c.CreateInitializedConnectionWithResponse(companyId, payload)
	return r, err
}

func (c *APIClient) CreateInitializedConnectionWithResponse(companyId string, payload *Connection) (*Connection, *http.Response, error) {
	connCreatePayload := Connection{
		Name:        payload.Name,
		ConnectorID: payload.ConnectorID,
	}

	resp, res, err := c.CreateConnectionWithResponse(companyId, &connCreatePayload)
	if err != nil {
		return nil, res, err
	}
	payload.ConnectionID = resp.ConnectionID

	if payload.Properties != nil {
		resp, res, err = c.UpdateConnectionWithResponse(companyId, payload)
		if err != nil {
			return nil, res, err
		}
	}

	return resp, res, nil
}

// Deletes a connection based on ConnectionId
func (c *APIClient) DeleteConnection(companyId string, connectionId string) (*Message, error) {
	r, _, err := c.DeleteConnectionWithResponse(companyId, connectionId)
	return r, err
}

func (c *APIClient) DeleteConnectionWithResponse(companyId string, connectionId string) (*Message, *http.Response, error) {
	req := DvHttpRequest{
		Method: "DELETE",
		Url:    fmt.Sprintf("%s/connections/%s", c.HostURL, connectionId),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	connection := Message{}
	err = json.Unmarshal(body, &connection)
	if err != nil {
		return nil, res, err
	}

	return &connection, res, nil
}
