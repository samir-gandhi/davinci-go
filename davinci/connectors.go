package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Only reading connectors is supported

// Gets array of all connectors for the provided company
func (c *APIClient) ReadConnectors(companyId *string, args *Params) ([]Connector, error) {
	r, _, err := c.ReadConnectorsWithResponse(companyId, args)
	return r, err
}

func (c *APIClient) ReadConnectorsWithResponse(companyId *string, args *Params) ([]Connector, *http.Response, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, res, err := c.SetEnvironmentWithResponse(cIdPointer)
	if err != nil {
		return nil, res, err
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/connectors", c.HostURL),
	}

	body, res, err := c.doRequestRetryable(companyId, req, args)
	if err != nil {
		return nil, res, err
	}

	connectionsLoose := []ConnectorLoose{}
	err = json.Unmarshal(body, &connectionsLoose)
	if err != nil {
		return nil, res, err
	}

	connections, err := connectorTypeAssign(connectionsLoose)
	if err != nil {
		return nil, res, fmt.Errorf("error casting connectors: %s", err)
	}

	return connections, res, nil
}

func connectorTypeAssign(cLoose []ConnectorLoose) ([]Connector, error) {
	cs := []Connector{}
	for i, v := range cLoose {
		cl := cLoose[i]
		for k, prop := range v.Properties {
			p := v.Properties[k]
			if info, ok := prop.Info.(string); ok {
				p.Info = info
			} else {
				p.Info = ""
			}
			cl.Properties[k] = p
		}
		vMarshal, err := json.Marshal(cl)
		if err != nil {
			return nil, err
		}
		c := Connector{}
		err = json.Unmarshal(vMarshal, &c)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return cs, nil
}

// Gets single connections based on ConnectionId
func (c *APIClient) ReadConnector(companyId *string, connectorId string) (*Connector, error) {
	r, _, err := c.ReadConnectorWithResponse(companyId, connectorId)
	return r, err
}

func (c *APIClient) ReadConnectorWithResponse(companyId *string, connectorId string) (*Connector, *http.Response, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}

	_, res, err := c.SetEnvironmentWithResponse(cIdPointer)
	if err != nil {
		return nil, res, err
	}
	if connectorId == "" {
		return nil, nil, fmt.Errorf("connectorId not provided")
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/connectors/%s", c.HostURL, connectorId),
	}

	body, res, err := c.doRequestRetryable(companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	connector := Connector{}
	err = json.Unmarshal(body, &connector)
	if err != nil {
		return nil, res, err
	}

	return &connector, res, nil
}
