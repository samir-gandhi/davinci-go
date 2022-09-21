package davinci

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

/* TEST PLAN:
* - Read all
* - Create new
* - Read new
* - Update new
* - Delete new
 */

// testData for Roles functions
var testDataFlows = map[string]interface{}{
	"params": map[string]Params{
		"limitTen": {Limit: "10"},
		"limitTwo": {Limit: "2"},
		"limitNil": {},
		// Flows doesn't allow page arg
		"pageNeg": {Page: "2"},
	},
	"flowsCreateJson": map[string]interface{}{
		"properImport":    `{"name":"tf testing","description":"","flowInfo":{"customerId":"dc7918cfa4b50966f8508072c2755c2c","name":"tf testing","description":"","flowStatus":"enabled","createdDate":1662960509175,"updatedDate":1662961640567,"currentVersion":0,"authTokenExpireIds":[],"deployedDate":1662960510106,"functionConnectionId":null,"isOutputSchemaSaved":false,"outputSchemaCompiled":null,"publishedVersion":1,"timeouts":null,"flowId":"bb45eb4a0e8a5c9d6a21c7ac2d1b3faa","companyId":"c431739a-29cd-4d9e-b465-0b37b2c235b1","versionId":0,"graphData":{"elements":{"nodes":[{"data":{"id":"pptape4ac2","nodeType":"CONNECTION","connectionId":"867ed4363b2bc21c860085ad2baa817d","connectorId":"httpConnector","name":"Http","label":"Http","status":"configured","capabilityName":"customHtmlMessage","type":"trigger","properties":{"message":{"value":"[\n  {\n    \"children\": [\n      {\n        \"text\": \"hello foobar\"\n      }\n    ]\n  }\n]"}}},"position":{"x":570,"y":240},"group":"nodes","removed":false,"selected":false,"selectable":true,"locked":false,"grabbable":true,"pannable":false,"classes":""}]},"data":{},"zoomingEnabled":true,"userZoomingEnabled":true,"zoom":1,"minZoom":1e-50,"maxZoom":1e+50,"panningEnabled":true,"userPanningEnabled":true,"pan":{"x":0,"y":0},"boxSelectionEnabled":true,"renderer":{"name":"null"}},"flowColor":"#AFD5FF","connectorIds":["httpConnector"],"savedDate":1662961640542,"variables":[]},"flowNameMapping":{"bb45eb4a0e8a5c9d6a21c7ac2d1b3faa":"tf testing"}}`,
		"directExport":    `{"customerId":"dc7918cfa4b50966f8508072c2755c2c","name":"tf testing","description":"","flowStatus":"enabled","createdDate":1662960509175,"updatedDate":1662961640567,"currentVersion":0,"authTokenExpireIds":[],"deployedDate":1662960510106,"functionConnectionId":null,"isOutputSchemaSaved":false,"outputSchemaCompiled":null,"publishedVersion":1,"timeouts":null,"flowId":"bb45eb4a0e8a5c9d6a21c7ac2d1b3faa","companyId":"c431739a-29cd-4d9e-b465-0b37b2c235b1","versionId":0,"graphData":{"elements":{"nodes":[{"data":{"id":"pptape4ac2","nodeType":"CONNECTION","connectionId":"867ed4363b2bc21c860085ad2baa817d","connectorId":"httpConnector","name":"Http","label":"Http","status":"configured","capabilityName":"customHtmlMessage","type":"trigger","properties":{"message":{"value":"[\n  {\n    \"children\": [\n      {\n        \"text\": \"hello foobar\"\n      }\n    ]\n  }\n]"}}},"position":{"x":570,"y":240},"group":"nodes","removed":false,"selected":false,"selectable":true,"locked":false,"grabbable":true,"pannable":false,"classes":""}]},"data":{},"zoomingEnabled":true,"userZoomingEnabled":true,"zoom":1,"minZoom":1e-50,"maxZoom":1e+50,"panningEnabled":true,"userPanningEnabled":true,"pan":{"x":0,"y":0},"boxSelectionEnabled":true,"renderer":{"name":"null"}},"flowColor":"#AFD5FF","connectorIds":["httpConnector"],"savedDate":1662961640542,"variables":[]}`,
		"directExportNeg": `{"customerId":"dc7918cfa4b50966f8508072c2755c2c","description":"","flowStatus":"enabled","createdDate":1662960509175,"updatedDate":1662961640567,"currentVersion":0,"authTokenExpireIds":[],"deployedDate":1662960510106,"functionConnectionId":null,"isOutputSchemaSaved":false,"outputSchemaCompiled":null,"publishedVersion":1,"timeouts":null,"flowId":"bb45eb4a0sdfsadf8a5c9d6a21c7ac2d1b3faa","companyId":"c431739a-asdfa29cd-4d9e-b465-0b37b2c235b1","versionId":0,"gaphData":{"s":[{"data":{"id":"pptaasfape4ac2","nodeType":"CONNECTION","connectionId":"867ed4363b2bc21c860085ad2baa817d","connectorId":"httpConnectorNO","name":"Http","label":"Http","status":"configured","capabilityName":"customHtmlMessage","type":"trigger","properties":{"message":{"value":"[\n  {\n    \"children\": [\n      {\n        \"text\": \"hello foobar\"\n      }\n    ]\n  }\n]"}}},"position":{"x":570,"y":240},"group":"nodes","removed":false,"selected":false,"selectable":true,"locked":false,"grabbable":true,"pannable":false,"classes":""}]},"data":{},"zoomingEnabled":true,"userZoomingEnabled":true,"zoom":1,"minZoom":1e-50,"maxZoom":1e+50,"panningEnabled":true,"userPanningEnabled":true,"pan":{"x":0,"y":0},"boxSelectionEnabled":true,"renderer":{"name":"null"}},"flowColor":"#AFD5FF","connectorIds":["httpConnector"],"savedDate":1662961640542,"variables":[]}`,
	},
}

// Gets an array of all roles for a company
func TestReadFlows(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataFlows["params"].(map[string]Params)
	for i := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			thisArgs := args[i]
			fmt.Printf("Test Args are: %q\n", thisArgs)
			msg := ""
			resp, err := c.ReadFlows(&c.CompanyID, &thisArgs)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					msg = fmt.Sprintf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				msg = fmt.Sprintf("Flows Returned Successfully\n resp[0].FlowId is: %+v \n", resp[0])
			}
			// Too verbose to print all Flows.
			fmt.Println(msg)
		})
	}
}

func TestCreateFlowWithJson(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	if args, ok := testDataFlows["flowsCreateJson"].(map[string]interface{}); ok {
		for i, thisArg := range args {
			testName := i
			t.Run(testName, func(t *testing.T) {
				fmt.Printf("thisArg is %q\n", thisArg)
				msg := ""
				if payloadJson, ok := thisArg.(string); ok {
					resp, err := c.CreateFlowWithJson(&c.CompanyID, &payloadJson)
					if err != nil {
						fmt.Println(err.Error())
						msg = fmt.Sprint("Failed Successfully\n")
						// if it's not a negative test, consider it an actual failure.
						if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
							msg = fmt.Sprintf("Failed with params: %v \n Error is: %v", args, err)
							t.Fail()
						}
					}
					if resp != nil {
						msg = fmt.Sprintf("Flows Created Successfully\n resp.FlowId is: %v \n", resp.FlowID)
					}
					fmt.Println(msg)
				}
			})
		}
	}
}

func TestReadFlow(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	flows, err := c.ReadFlows(&c.CompanyID, &Params{Limit: "3"})
	if err != nil {
		t.Fail()
	}
	for _, testVal := range flows {
		testName := testVal.Name
		t.Run(testName, func(t *testing.T) {
			msg := ""
			resp, err := c.ReadFlow(&c.CompanyID, testVal.FlowID)
			if err != nil {
				fmt.Println(err.Error())
				// msg = fmt.Sprint("Failed Successfully\n")
				// if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
				// 	msg = fmt.Sprintf("failed to get flows with param: %v \n Error is: %v", err)
				// 	t.Fail()
				// }
			}
			if resp != nil {
				msg = fmt.Sprintf("Flow Returned Successfully\n resp.FlowName is: %+v \n", resp.Flow.Name)
			}
			fmt.Println(msg)
		})
	}
}

//TODO - Use data specific to update for this.

func TestUpdateFlowWithJson(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	if args, ok := testDataFlows["flowsCreateJson"].(map[string]interface{}); ok {
		for i, thisArg := range args {
			testName := i
			t.Run(testName, func(t *testing.T) {
				fmt.Printf("thisArg is %q\n", thisArg)
				msg := ""
				if payloadJson, ok := thisArg.(string); ok {
					resp, err := c.CreateFlowWithJson(&c.CompanyID, &payloadJson)
					if err != nil {
						fmt.Println(err.Error())
						msg = fmt.Sprint("Failed Successfully\n")
						// if it's not a negative test, consider it an actual failure.
						if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
							msg = fmt.Sprintf("Failed with params: %v \n Error is: %v", args, err)
							t.Fail()
						}
					}
					if resp != nil {
						fmt.Printf("Flows Created Successfully\n resp.FlowId is: %v \n", resp.FlowID)
						resp.Name = fmt.Sprintf("%v-UPDATED", resp.Name)
						payload, err := json.Marshal(resp)
						if err != nil {
							t.Fail()
							return
						}
						payloadString := string(payload)
						resp, err := c.UpdateFlowWithJson(&c.CompanyID, &payloadString, resp.FlowID)
						if err != nil {
							msg = fmt.Sprintf("Update failed")
							t.Fail()
						}
						fmt.Printf("Update resp.Name: %v", resp.Name)
					}
					fmt.Println(msg)
				}
			})
		}
	}
}

func TestDeleteFlow(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	if args, ok := testDataFlows["flowsCreateJson"].(map[string]interface{}); ok {
		for i, thisArg := range args {
			testName := i
			t.Run(testName, func(t *testing.T) {
				fmt.Printf("thisArg is %q\n", thisArg)
				msg := ""
				if payloadJson, ok := thisArg.(string); ok {
					resp, err := c.CreateFlowWithJson(&c.CompanyID, &payloadJson)
					if err != nil {
						fmt.Println(err.Error())
						msg = fmt.Sprint("Failed Successfully\n")
						// if it's not a negative test, consider it an actual failure.
						if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
							msg = fmt.Sprintf("Failed with params: %v \n Error is: %v", args, err)
							t.Fail()
							return
						}
					}
					if resp != nil {
						fmt.Printf("Flows Created Successfully\n resp.FlowId is: %v \n", resp.FlowID)
						resp, err := c.DeleteFlow(&c.CompanyID, resp.FlowID)
						if err != nil {
							msg = fmt.Sprintf("Delete failed")
							t.Fail()
						}
						fmt.Printf("Deleted Successfully: %v", resp)
					}
					fmt.Println(msg)
				}
			})
		}
	}
}

func TestDeployFlow(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	if args, ok := testDataFlows["flowsCreateJson"].(map[string]interface{}); ok {
		for i, thisArg := range args {
			testName := i
			t.Run(testName, func(t *testing.T) {
				fmt.Printf("thisArg is %q\n", thisArg)
				msg := ""
				if payloadJson, ok := thisArg.(string); ok {
					resp, err := c.CreateFlowWithJson(&c.CompanyID, &payloadJson)
					if err != nil {
						fmt.Println(err.Error())
						msg = fmt.Sprint("Failed Successfully\n")
						// if it's not a negative test, consider it an actual failure.
						if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
							msg = fmt.Sprintf("Failed with params: %v \n Error is: %v", args, err)
							t.Fail()
							return
						}
					}
					if resp != nil {
						fmt.Printf("Flows Created Successfully\n resp.FlowId is: %v \n", resp.FlowID)
						resp, err := c.DeployFlow(&c.CompanyID, resp.FlowID)
						if err != nil {
							msg = fmt.Sprintf("Delete failed")
							t.Fail()
						}
						fmt.Printf("Deleted Successfully: %v", resp)
					}
					fmt.Println(msg)
				}
			})
		}
	}
}
