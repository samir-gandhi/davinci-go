package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (c *APIClient) ReadVariables(companyId string, args *Params) (map[string]Variable, error) {
	r, _, err := c.ReadVariablesWithResponse(companyId, args)
	return r, err
}

func (c *APIClient) ReadVariablesWithResponse(companyId string, args *Params) (map[string]Variable, *http.Response, error) {
	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/constructs", c.HostURL),
	}
	body, res, err := c.doRequestRetryable(&companyId, req, args)
	if err != nil {
		return nil, res, err
	}

	// Vars are returned as map
	resp := map[string]Variable{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, res, err
	}

	return resp, res, nil
}

func (c *APIClient) ReadVariable(companyId string, variableName string) (map[string]Variable, error) {
	r, _, err := c.ReadVariableWithResponse(companyId, variableName)
	return r, err
}

func (c *APIClient) ReadVariableWithResponse(companyId string, variableName string) (map[string]Variable, *http.Response, error) {
	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/constructs/%s", c.HostURL, url.PathEscape(variableName)),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	// Vars are returned as map
	resp := map[string]Variable{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, res, err
	}
	// removed, allow return of empty payload instead. Handling will be done in the caller
	// if len(resp) != 1 {
	// 	return nil, fmt.Errorf("status: 404, body: Variable not found or invalid data returned")
	// }

	return resp, res, nil
}

func (c *APIClient) CreateVariable(companyId string, variable *VariablePayload) (map[string]Variable, error) {
	r, _, err := c.CreateVariableWithResponse(companyId, variable)
	return r, err
}

func (c *APIClient) CreateVariableWithResponse(companyId string, variable *VariablePayload) (map[string]Variable, *http.Response, error) {
	validate := validator.New()
	if err := validate.Struct(variable); err != nil {
		return nil, nil, err
	}

	reqBody, err := json.Marshal(variable)
	if err != nil {
		return nil, nil, err
	}

	req := DvHttpRequest{
		Method: "POST",
		Url:    fmt.Sprintf("%s/constructs", c.HostURL),
		Body:   strings.NewReader(string(reqBody)),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}
	var resp map[string]Variable
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, res, err
	}

	return resp, res, nil
}

// UpdateVariable can update fields besides Name and Context
func (c *APIClient) UpdateVariable(companyId string, variable *VariablePayload) (map[string]Variable, error) {
	r, _, err := c.UpdateVariableWithResponse(companyId, variable)
	return r, err
}

func (c *APIClient) UpdateVariableWithResponse(companyId string, variable *VariablePayload) (map[string]Variable, *http.Response, error) {
	validate := validator.New()
	if err := validate.Struct(variable); err != nil {
		return nil, nil, err
	}

	// Identify if variable name in payload computed name or simple name
	vName := variable.Name
	computedName := url.PathEscape(*vName)
	if variable.Context != "flow" {
		regex := regexp.MustCompile(`^[a-zA-Z0-9\s]+##SK##[a-zA-Z0-9\s]+$`)
		if !regex.MatchString(*variable.Name) {
			computedName = url.PathEscape(fmt.Sprintf(`%v##SK##%s`, *vName, variable.Context))
		}
	}
	if variable.Context == "flow" {
		regex := regexp.MustCompile(`^[a-zA-Z0-9\s]+##SK##flow+##SK##[a-zA-Z0-9\s]+$`)
		if !regex.MatchString(*variable.Name) {
			computedName = url.PathEscape(fmt.Sprintf(`%v##SK##%s##SK##%v`, *vName, variable.Context, *variable.FlowId))
		}
	}
	//Variable Name should not be in the payload, it is not an updatable field
	variable.Name = nil

	reqBody, err := json.Marshal(variable)
	if err != nil {
		return nil, nil, err
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/constructs/%s", c.HostURL, computedName),
		Body:   strings.NewReader(string(reqBody)),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}
	var resp map[string]Variable
	err = json.Unmarshal(body, &resp)
	if err != nil {
		// Only account for Variable values of type string.
		// Force string if bool or int
		var variable map[string]VariablesValueInterface
		err = json.Unmarshal(body, &variable)
		if err != nil {
			return nil, res, err
		}
		if len(variable) != 1 {
			return nil, res, fmt.Errorf("status: 404, body: Variable not found or invalid data returned")
		}
		for i, v := range variable {
			if v.Value != nil {
				switch v.Value.(type) {
				case bool:
					v.Value = strconv.FormatBool(v.Value.(bool))
				case int:
					v.Value = strconv.Itoa(v.Value.(int))
				}
				variable[i] = v
			}
		}
		// Marshal back to json
		body, err = json.Marshal(variable)
		if err != nil {
			return nil, res, err
		}
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, res, err
		}
	}
	return resp, res, nil
}

func (c *APIClient) DeleteVariable(companyId string, variableName string) (*Message, error) {
	r, _, err := c.DeleteVariableWithResponse(companyId, variableName)
	return r, err
}

func (c *APIClient) DeleteVariableWithResponse(companyId string, variableName string) (*Message, *http.Response, error) {

	req := DvHttpRequest{
		Method: "DELETE",
		Url:    fmt.Sprintf("%s/constructs/%s", c.HostURL, url.PathEscape(variableName)),
	}

	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	// Vars are returned as map
	resp := Message{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, res, err
	}

	return &resp, res, nil
}
