package davinci

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ReadFlows only accepts Limit as a param
func (c *Client) ReadFlows(companyId *string, args *Params) ([]Flow, error) {
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

	cIdString := *cIdPointer
	log.Print(cIdString)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/flows", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, args)
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

// Create a bare connection, properties can be added _after_ creation
// func (c *Client) CreateFlow(companyId *string, payload *Flow) (*Flow, error) {
// 	cIdPointer := &c.CompanyID
// 	if companyId != nil {
// 		cIdPointer = companyId
// 	}
// 	msg, err := c.SetEnvironment(cIdPointer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("Set Env to: %s\n", msg.Message)

// 	if payload == nil || payload.Name == "" {
// 		return nil, fmt.Errorf("Empty or invalid payload")
// 	}

// 	reqBody, err := json.Marshal(payload)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/flows/import", c.HostURL), strings.NewReader(string(reqBody)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, &c.Token, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp := Flow{}
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &resp, nil
// }

func (c *Client) CreateFlowWithJson(companyId *string,
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
	pfi := FlowImport{}
	pf := Flow{}

	err = json.Unmarshal([]byte(*payloadJson), &pfi)
	if err != nil || pfi.FlowNameMapping == nil {
		log.Printf("Unable to unmarshal json to type FlowImport.\n Will try to unmarshal to type Flow")
		err = json.Unmarshal([]byte(*payloadJson), &pf)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type Flow.")
		}
		pfi = FlowImport{
			Name:            pf.Name,
			Description:     pf.Description,
			FlowNameMapping: map[string]interface{}{pf.FlowID: pf.Name},
			FlowInfo:        pf,
		}
	}
	payload, err := json.Marshal(pfi)
	// reqBody, err := json.Marshal(payloadJson)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/flows/import", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	resp := FlowInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Flow, nil
}

// ReadFlows only accepts Limit as a param
func (c *Client) ReadFlow(companyId *string, flowId string) (*FlowInfo, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	cIdString := *cIdPointer
	log.Print(cIdString)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/flows/%s", c.HostURL, flowId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
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
func (c *Client) UpdateFlowWithJson(companyId *string,
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
	// fmt.Printf("Set Env to: %s\n", msg.Message)
	pfi := FlowImport{}
	pf := Flow{}

	//handle incoming type Flow or Flow Import
	err = json.Unmarshal([]byte(*payloadJson), &pf)
	if err != nil {
		log.Printf("Unable to unmarshal json to type FlowImport.\n Will try to unmarshal to type Flow")
		err = json.Unmarshal([]byte(*payloadJson), &pfi)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type FlowImport.")
		}
		pf = pfi.FlowInfo
	}

	pAllowedProps := Flow{
		GraphData:      pf.GraphData,
		InputSchema:    pf.InputSchema,
		CurrentVersion: pf.CurrentVersion,
		Name:           pf.Name,
	}
	payload, err := json.Marshal(pAllowedProps)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/flows/%s", c.HostURL, pf.FlowID), strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	resp := FlowInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Flow, nil
}

// ReadFlows only accepts Limit as a param
func (c *Client) DeleteFlow(companyId *string, flowId string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	cIdString := *cIdPointer
	log.Print(cIdString)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/flows/%s", c.HostURL, flowId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
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
func (c *Client) DeployFlow(companyId *string, flowId string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	cIdString := *cIdPointer
	log.Print(cIdString)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/flows/%s/deploy", c.HostURL, flowId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
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
