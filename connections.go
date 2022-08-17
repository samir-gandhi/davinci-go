package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strings"
)

// Gets an array of all roles for a company
func (c *Client) GetConnections(companyId *string, args *Params) ([]Connection, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/company/%s/roles", c.HostURL, cIdString), nil)
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

// func UnmarshalUnstructuredJSON(map[string]interface{}) (map[string]interface{})
