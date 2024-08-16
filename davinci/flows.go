package davinci

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ReadFlows only accepts Limit as a param
func (c *APIClient) ReadFlows(companyId string, args *Params) ([]Flow, error) {
	r, _, err := c.ReadFlowsWithResponse(companyId, args)
	return r, err
}

func (c *APIClient) ReadFlowsWithResponse(companyId string, args *Params) ([]Flow, *http.Response, error) {
	if args != nil && args.Page != "" {
		log.Println("Param.Page found, not allowed, removing.")
		args.Page = ""
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows", c.HostURL),
	}
	body, res, err := c.doRequestRetryable(&companyId, req, args)
	if err != nil {
		return nil, res, err
	}

	// Returned flows are an array in top level flowsInfo key
	resp := FlowsInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, res, err
	}

	return resp.Flow, res, nil
}

func (c *APIClient) CreateFlow(companyId string, payload interface{}) (*Flow, error) {
	r, _, err := c.CreateFlowWithResponse(companyId, payload)
	return r, err
}

// CreateFlowWithResponse creates a flow with the given payload and returns the created flow.
// The payload can be a string (or *string) of JSON, or a FlowsImport, FlowImport, or Flow object
func (c *APIClient) CreateFlowWithResponse(companyId string, payload interface{}) (*Flow, *http.Response, error) {

	var payloadString string
	var err error

	switch v := payload.(type) {
	case FlowsImport, FlowImport:
		payloadBytes, err := json.Marshal(v)
		if err != nil {
			return nil, nil, err
		}
		payloadString = string(payloadBytes[:])
	case Flow:
		data := FlowImport{
			Name:        &v.Name,
			Description: v.Description,
			FlowInfo:    &v,
			FlowNameMapping: map[string]string{
				v.FlowID: v.Name,
			},
		}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			return nil, nil, err
		}
		payloadString = string(payloadBytes[:])
	case string:
		compiledPayload, err := makeFlowPayload(v, nil)
		if err != nil {
			return nil, nil, err
		}
		payloadString = *compiledPayload
	case *string:
		compiledPayload, err := makeFlowPayload(*v, nil)
		if err != nil {
			return nil, nil, err
		}
		payloadString = *compiledPayload
	default:
		return nil, nil, fmt.Errorf("Payload must be one of the following types: string, *string (where string types are valid JSON), FlowsImport, FlowImport, or Flow, got %T.", v)
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/flows/import", c.HostURL),
		Body:   payloadString,
	}

	body, resFlow, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, resFlow, err
	}

	resp := FlowInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, resFlow, err
	}

	// Call orx, this replicates the GET made from UI which seems to trigger some database function:
	reqOrx := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, resp.Flow.FlowID),
	}
	params := Params{
		ExtraParams: map[string]string{
			"attributes": "orx",
		},
	}
	_, res, err := c.doRequestRetryable(&companyId, reqOrx, &params)
	if err != nil {
		return nil, res, err
	}

	// Call apps, this replicates the GET made from UI which seems to trigger some database function:
	reqApps := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, resp.Flow.FlowID),
	}
	_, res, err = c.doRequestRetryable(&companyId, reqApps, nil)
	if err != nil {
		return nil, res, err
	}

	if resp.Flow.CompanyID != companyId {
		return nil, res, fmt.Errorf("Flow created with wrong companyId")
	}

	return &resp.Flow, resFlow, nil
}

// Deprecated: Use CreateFlowWithResponse(companyId string, payload interface{}) instead
func (c *APIClient) CreateFlowWithJson(companyId string, payloadJson *string) (*Flow, error) {
	r, _, err := c.CreateFlowWithJsonWithResponse(companyId, payloadJson)
	return r, err
}

// Deprecated: Use CreateFlowWithResponse(companyId string, payload interface{}) instead
func (c *APIClient) CreateFlowWithJsonWithResponse(companyId string, payloadJson *string) (*Flow, *http.Response, error) {
	return c.CreateFlowWithResponse(companyId, payloadJson)
}

// ReadFlowVersion is like ReadFlow, but appends a version query parameter.
// When called with no version, this returns what a flow export produces.
// version should be a string version of the version number, or nil for latest.
func (c *APIClient) ReadFlowVersion(companyId string, flowId string, flowVersion *string) (*FlowInfo, error) {
	r, _, err := c.ReadFlowVersionWithResponse(companyId, flowId, flowVersion)
	return r, err
}

func (c *APIClient) ReadFlowVersionWithResponse(companyId string, flowId string, flowVersion *string) (*FlowInfo, *http.Response, error) {
	// prior to the new feature flows always returned variable values. So we default to true
	return c.ReadFlowVersionOptionalVariableWithResponse(companyId, flowId, flowVersion, true)
}

// ReadFlowVersionOptionalVariableWithResponse is like ReadFlowVersionWithResponse
// but also accepts option to include variable values
func (c *APIClient) ReadFlowVersionOptionalVariableWithResponse(companyId string, flowId string, flowVersion *string, includeVariableValues bool) (*FlowInfo, *http.Response, error) {
	if flowVersion == nil {
		flow, res, err := c.ReadFlowWithResponse(companyId, flowId)
		if err != nil || flow == nil || flow.Flow.CurrentVersion == nil {
			return nil, res, err
		}

		fv := fmt.Sprint(*flow.Flow.CurrentVersion)
		flowVersion = &fv
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows/%s/versions/%s?includeSubflows=false&includeVariableValues=%s", c.HostURL, flowId, *flowVersion, fmt.Sprint(includeVariableValues)),
	}
	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}
	resp := FlowInfo{}
	err = json.Unmarshal(body, &resp.Flow)
	if err != nil {
		return nil, res, err
	}

	return &resp, res, nil
}

// ReadFlow performs a GET with no other parameters to get the latest version of the flow
func (c *APIClient) ReadFlow(companyId string, flowId string) (*FlowInfo, error) {
	r, _, err := c.ReadFlowWithResponse(companyId, flowId)
	return r, err
}

func (c *APIClient) ReadFlowWithResponse(companyId string, flowId string) (*FlowInfo, *http.Response, error) {

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, flowId),
	}
	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	// Returned flows are an array in top level flowsInfo key
	resp := FlowInfo{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, res, err
	}

	return &resp, res, nil
}

// Only specific fields are supported during update:
// - GraphData
// - InputSchema
// - CurrentVersion
// - Name
func (c *APIClient) UpdateFlow(companyId string, flowId string, payload interface{}) (*Flow, error) {
	r, _, err := c.UpdateFlowWithResponse(companyId, flowId, payload)
	return r, err
}

func (c *APIClient) UpdateFlowWithResponse(companyId string, flowId string, payload interface{}) (*Flow, *http.Response, error) {

	if companyId == "" {
		return nil, nil, fmt.Errorf("Must provide companyId.")
	}
	if flowId == "" {
		return nil, nil, fmt.Errorf("Must provide flowId.")
	}

	var payloadString string
	var err error

	switch v := payload.(type) {
	case Flow:

		flowBytes, err := json.Marshal(v)
		if err != nil {
			return nil, nil, err
		}

		pAllowedProps := FlowUpdate{}
		err = json.Unmarshal(flowBytes, &pAllowedProps)
		if err != nil {
			return nil, nil, err
		}

		if pAllowedProps.InputSchema == nil {
			pAllowedProps.InputSchema = make([]interface{}, 0)
		}

		currentFlow, res, err := c.ReadFlowWithResponse(companyId, flowId)
		if err != nil {
			return nil, res, err
		}

		pAllowedProps.CurrentVersion = currentFlow.Flow.CurrentVersion

		payloadBytes, err := json.Marshal(pAllowedProps)
		if err != nil {
			return nil, nil, err
		}
		payloadString = string(payloadBytes[:])
	case FlowUpdate:

		if v.InputSchema == nil {
			v.InputSchema = make([]interface{}, 0)
		}

		if v.CurrentVersion == nil {
			currentFlow, res, err := c.ReadFlowWithResponse(companyId, flowId)
			if err != nil {
				return nil, res, err
			}

			v.CurrentVersion = currentFlow.Flow.CurrentVersion
		}

		payloadBytes, err := json.Marshal(v)
		if err != nil {
			return nil, nil, err
		}
		payloadString = string(payloadBytes[:])
	case string:
		payloadString = v
	case *string:
		payloadString = *v
	default:
		return nil, nil, fmt.Errorf("Payload must be one of the following types: string, *string (where string types are valid JSON), Flow, FlowUpdate, got %T.", v)
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, flowId),
		Body:   payloadString,
	}

	body, resFlow, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, resFlow, err
	}

	resp := Flow{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, resFlow, err
	}

	return &resp, resFlow, nil
}

// Deprecated: Use UpdateFlowWithResponse(companyId string, flowId string, payload interface{}) instead
func (c *APIClient) UpdateFlowWithJson(companyId string, payloadJson *string, flowId string) (*Flow, error) {
	r, _, err := c.UpdateFlowWithJsonWithResponse(companyId, payloadJson, flowId)
	return r, err
}

// Deprecated: Use UpdateFlowWithResponse(companyId string, flowId string, payload interface{}) instead
func (c *APIClient) UpdateFlowWithJsonWithResponse(companyId string, payloadJson *string, flowId string) (*Flow, *http.Response, error) {
	return c.UpdateFlowWithResponse(companyId, flowId, payloadJson)
}

// ReadFlows only accepts Limit as a param
func (c *APIClient) DeleteFlow(companyId string, flowId string) (*Message, error) {
	r, _, err := c.DeleteFlowWithResponse(companyId, flowId)
	return r, err
}

func (c *APIClient) DeleteFlowWithResponse(companyId string, flowId string) (*Message, *http.Response, error) {

	req := DvHttpRequest{
		Method: "DELETE",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, flowId),
	}
	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	resp := Message{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, res, err
	}

	return &resp, res, nil
}

// ReadFlows only accepts Limit as a param
func (c *APIClient) DeployFlow(companyId string, flowId string) (*Message, error) {
	r, _, err := c.DeployFlowWithResponse(companyId, flowId)
	return r, err
}

func (c *APIClient) DeployFlowWithResponse(companyId string, flowId string) (*Message, *http.Response, error) {

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/flows/%s/deploy", c.HostURL, flowId),
	}
	body, res, err := c.doRequestRetryable(&companyId, req, nil)
	if err != nil {
		return nil, res, err
	}

	resp := Message{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, res, err
	}

	return &resp, res, nil
}

func makeFlowPayload(payload string, output interface{}) (*string, error) {
	if payload == "" {
		return nil, fmt.Errorf("Payload must be a non-empty string.")
	}

	if ok := json.Valid([]byte(payload[:])); !ok {
		return nil, fmt.Errorf("Payload must be valid JSON.")
	}

	if output == nil {
		output = FlowImport{}
		fis := Flows{}
		err := json.Unmarshal([]byte(payload), &fis)
		if err == nil && len(fis.Flow) > 0 {
			output = FlowsImport{}
		}
	}

	switch v := output.(type) {
	case FlowsImport:
		pfis, err := ParseFlowsImportJson(payload)
		if err != nil {
			return nil, err
		}
		for _, v := range pfis.FlowInfo.Flow {
			cleanseVariables(&v)
		}
		fjBytes, err := json.Marshal(pfis)
		if err != nil {
			return nil, fmt.Errorf("Unable to marshal payload.")
		}
		fjString := string(fjBytes)
		return &fjString, nil
	case Flow:
		pfi, err := parseFlowJson(payload)
		if err != nil {
			return nil, err
		}
		cleanseVariables(pfi)
		fjBytes, err := json.Marshal(pfi)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type Flow.")
		}
		fjString := string(fjBytes)
		return &fjString, nil
	case FlowImport:
		pfi, err := parseFlowImportJson(payload)
		if err != nil {
			return nil, err
		}
		cleanseVariables(pfi.FlowInfo)
		fjBytes, err := json.Marshal(pfi)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type FlowImport.")
		}
		fjString := string(fjBytes)
		return &fjString, nil
	default:
		return nil, fmt.Errorf("Output must be one of: FlowsImport{}, FlowImport{}, or Flow{}, got %s.", v)
	}
}

func parseFlowImportJson(payload string) (*FlowImport, error) {
	fi := FlowImport{}
	flow := Flow{}
	// is FlowImport or Flow
	err := json.Unmarshal([]byte(payload), &fi)
	if err == nil && fi.FlowNameMapping != nil {
		forceEmptyEvalProps(fi.FlowInfo)
		return &fi, nil
	}
	err = json.Unmarshal([]byte(payload), &flow)
	if err == nil {
		pfi := FlowImport{
			Name:            &flow.Name,
			Description:     flow.Description,
			FlowNameMapping: map[string]string{flow.FlowID: flow.Name},
			FlowInfo:        &flow,
		}
		forceEmptyEvalProps(pfi.FlowInfo)
		return &pfi, nil
	}
	return nil, fmt.Errorf("Unable to parse payload to type FlowImport")
}

func parseFlowJson(payload string) (*Flow, error) {
	fi := FlowImport{}
	flow := Flow{}
	// is FlowImport or Flow
	err := json.Unmarshal([]byte(payload), &flow)
	if err == nil && flow.GraphData.Elements.Nodes != nil {
		forceEmptyEvalProps(&flow)
		return &flow, nil
	}
	err = json.Unmarshal([]byte(payload), &fi)
	if err == nil {
		pfi := fi.FlowInfo
		forceEmptyEvalProps(pfi)
		return pfi, nil
	}
	return nil, fmt.Errorf("Unable to parse payload to type FlowImport")
}

func ParseFlowsImportJson(payload string) (*FlowsImport, error) {
	fis := Flows{}
	//is Flows
	err := json.Unmarshal([]byte(payload), &fis)
	if err == nil && len(fis.Flow) > 0 {
		pfis := FlowsImport{
			Name:            func() *string { s := ""; return &s }(),
			Description:     func() *string { s := ""; return &s }(),
			FlowInfo:        &fis,
			FlowNameMapping: map[string]string{},
		}
		for _, v := range fis.Flow {
			pfis.FlowNameMapping[v.FlowID] = v.Name
			forceEmptyEvalProps(&v)
		}
		return &pfis, nil
	}
	return nil, fmt.Errorf("Unable parse payload to type Flows")
}

func forceEmptyEvalProps(flow *Flow) {
	if flow != nil && flow.GraphData != nil && flow.GraphData.Elements != nil && flow.GraphData.Elements.Nodes != nil {
		for i, nodeData := range flow.GraphData.Elements.Nodes {
			nodeEval := "EVAL"
			if nodeData.Data.NodeType == &nodeEval {
				if flow.GraphData.Elements.Nodes[i].Data.Properties == nil {
					flow.GraphData.Elements.Nodes[i].Data.Properties = &Properties{}
				}
			}

			nodeConnection := "CONNECTION"
			if nodeData.Data.NodeType == &nodeConnection {
				if flow.GraphData.Elements.Nodes[i].Data.Properties == nil {
					flow.GraphData.Elements.Nodes[i].Data.Properties = &Properties{}
				}
			}
		}
	}
}

func cleanseVariables(flow *Flow) {
	fv := flow.Variables
	fvNew := []FlowVariable{}
	if len(fv) > 0 {
		for i, flowVar := range fv {
			flowTitle := "flow"
			if flowVar.Context == &flowTitle {
				fvNew = append(fvNew, fv[i])
			}
		}
	}
	flow.Variables = fvNew
}
