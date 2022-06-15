package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Returns list of Environments (auth required)
func (c *Client) GetEnvironments() ([]Environments, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/customers/me", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token)
	if err != nil {
		return nil, err
	}

	environments := []Environments{}
	err = json.Unmarshal(body, &environments)
	if err != nil {
		return nil, err
	}

	return environments, nil
}
