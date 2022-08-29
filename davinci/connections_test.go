package davinci

import (
	// "encoding/json"
	"fmt"
	// "strconv"
	"strings"
	"testing"
)

// type property struct {
// 	EnvId struct {
// 		Value string `json:"value`
// 	} `json:"envId`
// }

var testDataConnections = map[string]interface{}{
	"params": map[string]Params{
		"a": {"1", "10"},
		"b": {"1000", "10"},
		"c": {},
	},
	"connectionsCreate": map[string]Connection{
		"aCreate": {
			Name:        "connectionACreate",
			ConnectorID: "pingOneMfaConnector",
		},
		"bCreate": {
			Name:        "connectionBCreate",
			ConnectorID: "pingOneMfaConnector",
			CustomerID:  "1234",
		},
		"cCreateneg": {
			Name:        "",
			ConnectorID: "pingOneMfaConnector",
		},
		"dCreateneg": {
			Name:        "connectionDCreate",
			ConnectorID: "",
		},
	},
	"connectionsRead": map[string]Connection{
		"aRead": {
			Name:        "connectionARead",
			ConnectorID: "pingOneMfaConnector",
		},
		"bRead": {
			Name:        "connectionBRead",
			ConnectorID: "pingOneMfaConnector",
			CustomerID:  "1234",
		},
		"cReadneg": {
			Name:        "",
			ConnectorID: "pingOneMfaConnector",
		},
		"dReadneg": {
			Name:        "connectionDRead",
			ConnectorID: "",
		},
	},
	"connectionsUpdate": map[string]Connection{
		"aUpdate": {
			Name:        "connectionAUpdate",
			ConnectorID: "pingOneMfaConnector",
			Properties: Properties{
				"envId": struct {
					Value string `json:"value"`
				}{"1234"},
				"policyId": struct {
					Value string `json:"value"`
				}{"1234"},
			},
		},
		"bUpdate": {
			Name:        "connectionBUpdate",
			ConnectorID: "pingOneSSOConnector",
			Properties: Properties{
				"envId": struct {
					Value string `json:"value"`
				}{"1234"},
			},
		},
		"cUpdate": {
			Name:        "connectionCUpdate",
			ConnectorID: "pingOneMfaConnector",
			Properties:  Properties{},
		},
		"dUpdateNeg": {
			Name:        "connectionCUpdateNeg",
			ConnectorID: "pingOneMfaConnector",
			Properties: Properties{
				"InvalidProperty": struct {
					Value string `json:"value"`
				}{"Foo"},
			},
		},
	},
	"connectionsCreateInitialized": map[string]Connection{
		"aCreateInitialized": {
			Name:        "connectionACreateInitialized",
			ConnectorID: "pingOneMfaConnector",
			Properties: Properties{
				"envId": struct {
					Value string `json:"value"`
				}{"1234"},
				"policyId": struct {
					Value string `json:"value"`
				}{"1234"},
			},
		},
		"bCreateInitialized": {
			Name:        "connectionBCreateInitialized",
			ConnectorID: "pingOneSSOConnector",
			Properties: Properties{
				"envId": struct {
					Value string `json:"value"`
				}{"1234"},
			},
		},
		"cCreateInitialized": {
			Name:        "connectionCCreateInitialized",
			ConnectorID: "pingOneMfaConnector",
			Properties:  Properties{},
		},
		"dCreateInitializedNeg": {
			Name:        "connectionDCreateInitialized",
			ConnectorID: "pingOneMfaConnector",
			Properties: Properties{
				"InvalidProperty": struct {
					Value string `json:"value"`
				}{"Foo"},
			},
		},
	},
}

func TestReadConnections(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataConnections["params"].(map[string]Params)
	for i := range args {
		thisArgs := args[i]
		fmt.Printf("args[i] is %q\n", thisArgs)
		resp, err := c.ReadConnections(&c.CompanyID, &thisArgs)
		if err != nil {
			fmt.Println(err.Error())
			if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
				t.Fail()
			}
		}
		fmt.Printf("resp is: %v\n", resp)
	}
}

func TestCreateConnection(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	if args, ok := testDataConnections["connectionsCreate"].(map[string]Connection); ok {
		for i := range args {
			thisArg := args[i]
			fmt.Printf("args[i] is %q\n", thisArg)
			resp, err := c.CreateConnection(&c.CompanyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					t.Fail()
				}
				fmt.Printf("Connection from key %v failed, continuing to next iteration \n", i)
				continue
			}
			thisArg.ConnectionID = resp.ConnectionID
			args[i] = thisArg
			fmt.Printf("resp is: %q\n", resp)
		}
	}
}

func TestReadConnection(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	if args, ok := testDataConnections["connectionsRead"].(map[string]Connection); ok {
		for i := range args {
			thisArg := args[i]
			fmt.Printf("thisArg is %q\n", thisArg)
			// Reading connection requires generated ConnectionID
			resp, err := c.CreateConnection(&c.CompanyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) {
					t.Fail()
				}
				// stop iteration of loop if connection create failed.
				fmt.Printf("Connection from key %v creation failed, continuing to next iteration \n", i)
				continue
			}
			thisArg.ConnectionID = resp.ConnectionID
			args[i] = thisArg
			fmt.Println(args[i])
			fmt.Printf("resp is: %q\n", resp)
			res, err := c.ReadConnection(&c.CompanyID, resp.ConnectionID)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					t.Fail()
				}
			}
			fmt.Printf("res is: %v\n", res)
		}
	}
}

func TestUpdateConnection(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	if args, ok := testDataConnections["connectionsUpdate"].(map[string]Connection); ok {
		for i := range args {
			thisArg := args[i]
			fmt.Printf("args[i] is %q\n", thisArg)
			// Reading connection requires generated ConnectionID
			// For testing, the connection is created to get ConnectionID from response
			resp, err := c.CreateConnection(&c.CompanyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) {
					t.Fail()
				}
				// stop iteration of loop if connection create failed.
				fmt.Printf("Connection from key %v creation failed, continuing to next iteration \n", i)
				continue
			}
			thisArg.ConnectionID = resp.ConnectionID
			args[i] = thisArg
			fmt.Printf("resp is: %q\n", resp)
			//upon success, pull properties from test data
			resp.Properties = thisArg.Properties
			res, err := c.UpdateConnection(&c.CompanyID, resp)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					t.Fail()
				}
			}
			fmt.Printf("res is: %v\n", res)
		}
	}
}

func TestCreateInitializedConnection(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	if args, ok := testDataConnections["connectionsCreateInitialized"].(map[string]Connection); ok {
		for i := range args {
			thisArg := args[i]
			fmt.Printf("args[i] is %q\n", thisArg)
			// Reading connection requires generated ConnectionID
			// For testing, the connection is created to get ConnectionID from response
			resp, err := c.CreateInitializedConnection(&c.CompanyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					t.Fail()
				}
				// stop iteration of loop if connection create failed.
				fmt.Printf("Initialized Connection from key %v creation failed, continuing to next iteration \n", i)
				continue
			}
			thisArg.ConnectionID = resp.ConnectionID
			args[i] = thisArg
			fmt.Printf("resp is: %q\n", resp)
		}
	}
}

func TestDeleteConnection(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	for i := range testDataConnections {
		if args, ok := testDataConnections[i].(map[string]Connection); ok {
			for j := range args {
				thisArg := args[j]
				fmt.Printf("thisArg is %q\n", thisArg)
				// Reading connection requires generated ConnectionID
				res, err := c.ReadConnection(&c.CompanyID, thisArg.ConnectionID)
				if err != nil {
					fmt.Println(err.Error())
					// if it's not a negative test, consider it an actual failure.
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						t.Fail()
					}
					fmt.Printf("Connection from key %v failed, continuing to next iteration \n", j)
					continue
				}
				fmt.Printf("res is: %v\n", res)
				resp, err := c.DeleteConnection(&c.CompanyID, thisArg.ConnectionID)
				if err != nil {
					fmt.Println(err.Error())
					// if it's not a negative test, consider it an actual failure.
					if !(strings.Contains(j, "neg")) && !(strings.Contains(j, "Neg")) {
						fmt.Println("Failed Successfully")
						t.Fail()
					}
					// stop iteration of loop if connection create failed.
					fmt.Printf("Connection from key %v failed, continuing to next iteration \n", j)
					continue
				}
				fmt.Printf("resp is: %q\n", resp)
			}
		}
	}
}
