package davinci

import (
	"encoding/json"
	"fmt"
)

// Only reading connectors is supported

// Gets array of all connectors for the provided company
func (c *APIClient) ReadConnectors(companyId *string, args *Params) ([]Connector, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/connectors", c.HostURL),
	}

	body, err := c.doRequestRetryable(req, &c.Token, args)
	if err != nil {
		return nil, err
	}

	connectionsLoose := []ConnectorLoose{}
	err = json.Unmarshal(body, &connectionsLoose)
	if err != nil {
		return nil, err
	}

	connections, err := connectorTypeAssign(connectionsLoose)
	if err != nil {
		return nil, fmt.Errorf("error casting connectors: %s", err)
	}

	return connections, nil
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
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}

	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	if connectorId == "" {
		return nil, fmt.Errorf("connectorId not provided")
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/connectors/%s", c.HostURL, connectorId),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	connector := Connector{}
	err = json.Unmarshal(body, &connector)
	if err != nil {
		return nil, err
	}

	return &connector, nil
}
