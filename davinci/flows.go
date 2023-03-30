package davinci

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// ReadFlows only accepts Limit as a param
func (c *APIClient) ReadFlows(companyId *string, args *Params) ([]Flow, error) {
	if args.Page != "" {
		log.Println("Param.Page found, not allowed, removing.")
		args.Page = ""
	}
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	// req, err := http.NewRequest("GET", fmt.Sprintf("%s/flows", c.HostURL), nil)
	// if err != nil {
	// 	return nil, err
	// }
	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows", c.HostURL),
	}
	body, err := c.doRequestRetryable(req, &c.Token, args)
	if err != nil {
		return nil, err
	}

	// Returned flows are an array in top level flowsInfo key
	resp := FlowsInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Flow, nil
}

type flowJson struct {
	payload *string
}

func ParseFlowImportJson(payload *string) (*FlowImport, error) {
	fi := FlowImport{}
	flow := Flow{}
	// is FlowImport or Flow
	err := json.Unmarshal([]byte(*payload), &fi)
	if err == nil && fi.FlowNameMapping != nil {
		forceEmptyEvalProps(&fi.FlowInfo)
		return &fi, nil
	}
	err = json.Unmarshal([]byte(*payload), &flow)
	if err == nil {
		pfi := FlowImport{
			Name:            flow.Name,
			Description:     flow.Description,
			FlowNameMapping: map[string]string{flow.FlowID: flow.Name},
			FlowInfo:        flow,
		}
		forceEmptyEvalProps(&pfi.FlowInfo)
		return &pfi, nil
	}
	return nil, fmt.Errorf("Unable to parse payload to type FlowImport")
}

func ParseFlowJson(payload *string) (*Flow, error) {
	fi := FlowImport{}
	flow := Flow{}
	// is FlowImport or Flow
	err := json.Unmarshal([]byte(*payload), &flow)
	if err == nil && flow.GraphData.Elements.Nodes != nil {
		forceEmptyEvalProps(&flow)
		return &flow, nil
	}
	err = json.Unmarshal([]byte(*payload), &fi)
	if err == nil {
		pfi := fi.FlowInfo
		forceEmptyEvalProps(&pfi)
		return &pfi, nil
	}
	return nil, fmt.Errorf("Unable to parse payload to type FlowImport")
}

func ParseFlowsImportJson(payload *string) (*FlowsImport, error) {
	fis := Flows{}
	//is Flows
	err := json.Unmarshal([]byte(*payload), &fis)
	if err == nil && len(fis.Flow) > 0 {
		pfis := FlowsImport{
			Name:            "",
			Description:     "",
			FlowInfo:        fis,
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

// Force properties on EVAL nodes to be empty object instead of null.
// Empty object should be harmless for EVAL nodes, but is necessary for CONNECTION nodes
func forceEmptyEvalProps(flow *Flow) {
	for i, nodeData := range flow.GraphData.Elements.Nodes {
		if nodeData.Data.NodeType == "EVAL" {
			if flow.GraphData.Elements.Nodes[i].Data.Properties == nil {
				flow.GraphData.Elements.Nodes[i].Data.Properties = map[string]interface{}{}
			}
		}
	}
}

// cleanseVariables removes all variables that are not flow context
// this is to ensure the flow ONLY created variables that singularly attached to the flow
func cleanseVariables(flow *Flow) {
	fv := flow.Variables
	if len(fv) > 0 {
		for i, flowVar := range fv {
			if flowVar.Context != "flow" {
				// if on last element, nullify the slice instead of removing the element
				if len(fv) == 1 {
					fv = []FlowVariable{}
					break
				}
				fv[i] = fv[len(fv)-1]
				fv = fv[:len(fv)-1]
			}
		}
	}
}

// MakeFlowPayload accepts
// payload: string of format Flows, FlowImport, or Flow
// output: desired type of FlowsImport, FlowImport, or Flow
// Payloads can only be converted to matching plurality
func MakeFlowPayload(payload *string, output string) (*string, error) {
	if output == "" {
		output = "FlowImport"
		fis := Flows{}
		err := json.Unmarshal([]byte(*payload), &fis)
		if err == nil && len(fis.Flow) > 0 {
			output = "FlowsImport"
		}
	}
	switch output {
	case "FlowsImport":
		pfis, _ := ParseFlowsImportJson(payload)
		for _, v := range pfis.FlowInfo.Flow {
			cleanseVariables(&v)
		}
		fjBytes, err := json.Marshal(pfis)
		if err != nil {
			return nil, fmt.Errorf("Unable to marshal payload.")
		}
		fjString := string(fjBytes)
		payloadString := &fjString
		return payloadString, nil
	case "Flow":
		pfi, _ := ParseFlowJson(payload)
		cleanseVariables(pfi)
		fjBytes, err := json.Marshal(pfi)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type Flow.")
		}
		fjString := string(fjBytes)
		payload = &fjString
		return payload, nil
	case "FlowImport":
		pfi, _ := ParseFlowImportJson(payload)
		cleanseVariables(&pfi.FlowInfo)
		fjBytes, err := json.Marshal(pfi)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type FlowImport.")
		}
		fjString := string(fjBytes)
		payload = &fjString
		return payload, nil
	default:
		return nil, fmt.Errorf("Output must be one of: FlowsImport, FlowImport, or Flow.")
	}
}

func (c *APIClient) CreateFlowWithJson(companyId *string,
	payloadJson *string) (*Flow, error) {
	if payloadJson == nil {
		return nil, fmt.Errorf("Must provide payloadJson.")
	}
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	payload, err := MakeFlowPayload(payloadJson, "")
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/flows/import", c.HostURL),
		Body:   strings.NewReader(*payload),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	resp := FlowInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
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
	_, err = c.doRequestRetryable(reqOrx, &c.Token, &params)
	if err != nil {
		return nil, err
	}

	// Call apps, this replicates the GET made from UI which seems to trigger some database function:
	reqApps := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, resp.Flow.FlowID),
	}
	_, err = c.doRequestRetryable(reqApps, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	return &resp.Flow, nil
}

// ReadFlowVersion is like ReadFlow, but appends a version query parameter.
// When called with no version, this returns what a flow export produces.
// version should be a string version of the version number, or nil for latest.
func (c *APIClient) ReadFlowVersion(companyId *string, flowId string, flowVersion *string) (*FlowInfo, error) {
	if flowVersion == nil {
		flow, err := c.ReadFlow(companyId, flowId)
		if err != nil {
			return nil, err
		}
		fv := strconv.Itoa(flow.Flow.CurrentVersion)
		flowVersion = &fv
	}
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	//sample version endpoint:
	//Request URL: https://orchestrate-api.pingone.com/v1/flows/ea578b4b66ff8cb4f015e4e1109dc872/versions/14?includeSubFlows=false
	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows/%s/versions/%s?includeSubflows=false", c.HostURL, flowId, *flowVersion),
	}
	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}
	resp := FlowInfo{}
	err = json.Unmarshal(body, &resp.Flow)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ReadFlow performs a GET with no other parameters to get the latest version of the flow
func (c *APIClient) ReadFlow(companyId *string, flowId string) (*FlowInfo, error) {
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
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, flowId),
	}
	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	// Returned flows are an array in top level flowsInfo key
	resp := FlowInfo{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Only specific fields are supported during update:
// - GraphData
// - InputSchema
// - CurrentVersion
// - Name
func (c *APIClient) UpdateFlowWithJson(companyId *string, payloadJson *string, flowId string) (*Flow, error) {
	if payloadJson == nil {
		return nil, fmt.Errorf("Must provide payloadJson.")
	}
	if flowId == "" {
		return nil, fmt.Errorf("Must provide flowId.")
	}
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	pf := Flow{}
	flow, err := MakeFlowPayload(payloadJson, "Flow")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(*flow), &pf)

	currentFlow, err := c.ReadFlow(cIdPointer, flowId)
	if err != nil {
		return nil, err
	}

	// since InputSchema is []interface, have to make a slice to ensure InputSchema is empty array if nil
	if pf.InputSchema == nil {
		pf.InputSchema = make([]interface{}, 0)
	}

	pAllowedProps := FlowUpdate{
		CurrentVersion: currentFlow.Flow.CurrentVersion,
		Name:           pf.Name,
		Description:    pf.Description,
		Settings:       pf.Settings,
		Trigger:        pf.Trigger,
		GraphData:      pf.GraphData,
		InputSchema:    pf.InputSchema,
		// not sure if it's used
		// InputSchemaCompiled: pf.InputSchemaCompiled,
		// not allowed
		// IsInputSchemaSaved:  pf.IsInputSchemaSaved,
		OutputSchema: pf.OutputSchemaCompiled,
	}

	payload, err := json.Marshal(pAllowedProps)
	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, flowId),
		Body:   strings.NewReader(string(payload)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	resp := Flow{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ReadFlows only accepts Limit as a param
func (c *APIClient) DeleteFlow(companyId *string, flowId string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "DELETE",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, flowId),
	}
	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	resp := Message{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ReadFlows only accepts Limit as a param
func (c *APIClient) DeployFlow(companyId *string, flowId string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	// req, err := http.NewRequest("PUT", fmt.Sprintf("%s/flows/%s/deploy", c.HostURL, flowId), nil)
	// if err != nil {
	// 	return nil, err
	// }
	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/flows/%s/deploy", c.HostURL, flowId),
	}
	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	resp := Message{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
