package davinci

// import (
// 	// "encoding/json"
// 	"encoding/json"
// 	"fmt"

// 	// "strconv"
// 	"math/rand"
// 	"strings"
// 	"testing"
// 	"time"
// )

// var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func RandString(n int) string {
// 	rand.Seed(time.Now().UnixNano())
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letters[rand.Intn(len(letters))]
// 	}
// 	return string(b)
// }

// var randString = RandString(6)

// // type property struct {
// // 	EnvId struct {
// // 		Value string `json:"value`
// // 	} `json:"envId`
// // }

// var testDataConnections = map[string]interface{}{
// 	"params": map[string]Params{
// 		"a": {"1", "10", nil},
// 		"b": {"1000", "10", nil},
// 		"c": {},
// 	},
// 	"connectionsCreate": map[string]Connection{
// 		"aCreate": {
// 			Name:        "connectionACreate",
// 			ConnectorID: "pingOneMfaConnector",
// 		},
// 		"bCreate": {
// 			Name:        "connectionBCreate",
// 			ConnectorID: "pingOneMfaConnector",
// 			CustomerID:  "1234",
// 		},
// 		"cCreateneg": {
// 			Name:        "",
// 			ConnectorID: "pingOneMfaConnector",
// 		},
// 		"dCreateneg": {
// 			Name:        "connectionDCreate",
// 			ConnectorID: "",
// 		},
// 	},
// 	"connectionsRead": map[string]Connection{
// 		"aRead": {
// 			Name:        "connectionARead",
// 			ConnectorID: "pingOneMfaConnector",
// 		},
// 		"bRead": {
// 			Name:        "connectionBRead",
// 			ConnectorID: "pingOneMfaConnector",
// 			CustomerID:  "1234",
// 		},
// 		"cReadneg": {
// 			Name:        "",
// 			ConnectorID: "pingOneMfaConnector",
// 		},
// 		"dReadneg": {
// 			Name:        "connectionDRead",
// 			ConnectorID: "",
// 		},
// 		"eRead": {
// 			Name:        "connectionERead" + randString,
// 			ConnectorID: "genericConnector",
// 		},
// 	},
// 	"connectionsUpdate": map[string]Connection{
// 		"aUpdate": {
// 			Name:        "connectionAUpdate",
// 			ConnectorID: "pingOneMfaConnector",
// 			Properties: func() *Properties {
// 				return &Properties{
// 					AdditionalProperties: map[string]interface{}{
// 						"envId": struct {
// 							Value string `json:"value"`
// 						}{"1234"},
// 						"policyId": struct {
// 							Value string `json:"value"`
// 						}{"1234"},
// 					},
// 				}
// 			}(),
// 		},
// 		"bUpdate": {
// 			Name:        "connectionBUpdate",
// 			ConnectorID: "pingOneSSOConnector",
// 			Properties: func() *Properties {
// 				return &Properties{
// 					AdditionalProperties: map[string]interface{}{
// 						"envId": struct {
// 							Value string `json:"value"`
// 						}{"1234"},
// 					},
// 				}
// 			}(),
// 		},
// 		"cUpdate": {
// 			Name:        "connectionCUpdate",
// 			ConnectorID: "pingOneMfaConnector",
// 			Properties: func() *Properties {
// 				return &Properties{
// 					AdditionalProperties: map[string]interface{}{},
// 				}
// 			}(),
// 		},
// 		"dUpdateNeg": {
// 			Name:        "connectionCUpdateNeg",
// 			ConnectorID: "pingOneMfaConnector",
// 			Properties: func() *Properties {
// 				return &Properties{
// 					AdditionalProperties: map[string]interface{}{
// 						"InvalidProperty": struct {
// 							Value string `json:"value"`
// 						}{"Foo"},
// 					},
// 				}
// 			}(),
// 		},
// 	},
// 	"connectionsCreateInitialized": map[string]Connection{
// 		// "aCreateInitialized": {
// 		// 	Name:        "connectionACreateInitialized",
// 		// 	ConnectorID: "pingOneMfaConnector",
// 		// 	Properties: Properties{
// 		// 		"envId": struct {
// 		// 			Value string `json:"value"`
// 		// 		}{"1234"},
// 		// 		"policyId": struct {
// 		// 			Value string `json:"value"`
// 		// 		}{"1234"},
// 		// 	},
// 		// },
// 		// "bCreateInitialized": {
// 		// 	Name:        "connectionBCreateInitialized",
// 		// 	ConnectorID: "pingOneSSOConnector",
// 		// 	Properties: Properties{
// 		// 		"envId": struct {
// 		// 			Value string `json:"value"`
// 		// 		}{"1234"},
// 		// 	},
// 		// },
// 		// "cCreateInitialized": {
// 		// 	Name:        "connectionCCreateInitialized",
// 		// 	ConnectorID: "pingOneMfaConnector",
// 		// 	Properties:  Properties{},
// 		// },
// 		// "dCreateInitializedNeg": {
// 		// 	Name:        "connectionDCreateInitialized",
// 		// 	ConnectorID: "pingOneMfaConnector",
// 		// 	Properties: Properties{
// 		// 		"InvalidProperty": struct {
// 		// 			Value string `json:"value"`
// 		// 		}{"Foo"},
// 		// 	},
// 		// },
// 		"eCreateInitializedOidc": {
// 			Name:        "connectionECreateInitializedOidc" + randString,
// 			ConnectorID: "genericConnector",
// 			Properties: func() *Properties {
// 				return &Properties{
// 					AdditionalProperties: map[string]interface{}{
// 						"customAuth": CustomAuth{
// 							Properties: CustomAuthProperties{
// 								ProviderName: ProviderName{
// 									Value: "foooidc",
// 								},
// 							},
// 						},
// 					},
// 				}
// 			}(),
// 		},
// 	},
// }

// func TestReadConnections(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	args, _ := testDataConnections["params"].(map[string]Params)
// 	for i := range args {
// 		thisArgs := args[i]
// 		fmt.Printf("args[i] is %q\n", thisArgs)
// 		resp, err := c.ReadConnections(c.CompanyID, &thisArgs)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 				t.Fail()
// 			}
// 		}
// 		fmt.Printf("resp is: %v\n", resp)
// 	}
// }

// func TestCreateConnection(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	if args, ok := testDataConnections["connectionsCreate"].(map[string]Connection); ok {
// 		for i := range args {
// 			thisArg := args[i]
// 			fmt.Printf("args[i] is %#v\n", thisArg)
// 			resp, err := c.CreateConnection(c.CompanyID, &thisArg)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				// if it's not a negative test, consider it an actual failure.
// 				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 					t.Fail()
// 				}
// 				fmt.Printf("Connection from key %v failed, continuing to next iteration \n", i)
// 				continue
// 			}
// 			thisArg.ConnectionID = resp.ConnectionID
// 			args[i] = thisArg
// 			fmt.Printf("resp is: %#v\n", resp)
// 		}
// 	}
// }

// // DELETE
// func TestReadConnection_Oidc(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := c.ReadConnection(c.CompanyID, "3b51289bf0126ac190d61284920d99e4")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		t.Fail()
// 	}
// 	a, _ := json.Marshal(resp.Properties["customAuth"])
// 	var customAuth CustomAuth
// 	json.Unmarshal(a, &customAuth)
// 	if customAuth.Properties.ClientID.Value == "" {
// 		t.Fail()
// 	}
// }

// //TODO - instead of creating connection for test, read all connections for connectionId

// func TestReadConnection(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	if args, ok := testDataConnections["connectionsRead"].(map[string]Connection); ok {
// 		for i := range args {
// 			thisArg := args[i]
// 			fmt.Printf("thisArg is %#v\n", thisArg)
// 			// Reading connection requires generated ConnectionID
// 			resp, err := c.CreateConnection(c.CompanyID, &thisArg)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				// if it's not a negative test, consider it an actual failure.
// 				if !(strings.Contains(i, "neg")) {
// 					t.Fail()
// 				}
// 				// stop iteration of loop if connection create failed.
// 				fmt.Printf("Connection from key %v creation failed, continuing to next iteration \n", i)
// 				continue
// 			}
// 			thisArg.ConnectionID = resp.ConnectionID
// 			args[i] = thisArg
// 			fmt.Println(args[i])
// 			fmt.Printf("resp is: %#v\n", resp)
// 			res, err := c.ReadConnection(c.CompanyID, resp.ConnectionID)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				// if it's not a negative test, consider it an actual failure.
// 				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 					t.Fail()
// 				}
// 			}
// 			fmt.Printf("res is: %v\n", res)
// 			if res.Properties["customAuth"] != nil {
// 				fmt.Println("customAuth is not nil, will attempt to unmarshal")
// 				a, _ := json.Marshal(res.Properties["customAuth"])
// 				var customAuth CustomAuth
// 				json.Unmarshal(a, &customAuth)
// 				if customAuth.Properties.ClientID.Value == "" {
// 					fmt.Println("customAuth.Properties.ClientID.Value is empty after unmarshal")
// 					t.Fail()
// 				}
// 			}
// 		}
// 	}
// }

// func TestUpdateConnection(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	if args, ok := testDataConnections["connectionsUpdate"].(map[string]Connection); ok {
// 		for i := range args {
// 			thisArg := args[i]
// 			fmt.Printf("args[i] is %#v\n", thisArg)
// 			// Reading connection requires generated ConnectionID
// 			// For testing, the connection is created to get ConnectionID from response
// 			resp, err := c.CreateConnection(c.CompanyID, &thisArg)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				// if it's not a negative test, consider it an actual failure.
// 				if !(strings.Contains(i, "neg")) {
// 					t.Fail()
// 				}
// 				// stop iteration of loop if connection create failed.
// 				fmt.Printf("Connection from key %v creation failed, continuing to next iteration \n", i)
// 				continue
// 			}
// 			thisArg.ConnectionID = resp.ConnectionID
// 			args[i] = thisArg
// 			fmt.Printf("resp is: %#v\n", resp)
// 			//upon success, pull properties from test data
// 			resp.Properties = thisArg.Properties
// 			res, err := c.UpdateConnection(c.CompanyID, resp)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				// if it's not a negative test, consider it an actual failure.
// 				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 					t.Fail()
// 				}
// 			}
// 			fmt.Printf("res is: %v\n", res)
// 		}
// 	}
// }

// func TestCreateInitializedConnection(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	if args, ok := testDataConnections["connectionsCreateInitialized"].(map[string]Connection); ok {
// 		for i := range args {
// 			thisArg := args[i]
// 			fmt.Printf("args[i] is %#v\n", thisArg)
// 			// Reading connection requires generated ConnectionID
// 			// For testing, the connection is created to get ConnectionID from response
// 			resp, err := c.CreateInitializedConnection(c.CompanyID, &thisArg)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				// if it's not a negative test, consider it an actual failure.
// 				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 					t.Fail()
// 				}
// 				// stop iteration of loop if connection create failed.
// 				fmt.Printf("Initialized Connection from key %v creation failed, continuing to next iteration \n", i)
// 				continue
// 			}
// 			thisArg.ConnectionID = resp.ConnectionID
// 			args[i] = thisArg
// 			fmt.Printf("resp is: %#v\n", resp)
// 			if resp.Properties["customAuth"] != nil {
// 				fmt.Println("customAuth is not nil, will attempt to unmarshal")
// 				a, _ := json.Marshal(resp.Properties["customAuth"])
// 				var customAuth CustomAuth
// 				json.Unmarshal(a, &customAuth)
// 				if customAuth.Properties.ProviderName.Value == "" {
// 					fmt.Println("customAuth.Properties.ProviderName.Value is empty after unmarshal")
// 					t.Fail()
// 				}
// 			}
// 		}
// 	}
// }

// func TestDeleteConnection(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	for i := range testDataConnections {
// 		if args, ok := testDataConnections[i].(map[string]Connection); ok {
// 			for j := range args {
// 				thisArg := args[j]
// 				fmt.Printf("thisArg is %#v\n", thisArg)
// 				// Reading connection requires generated ConnectionID
// 				res, err := c.ReadConnection(c.CompanyID, thisArg.ConnectionID)
// 				if err != nil {
// 					fmt.Println(err.Error())
// 					// if it's not a negative test, consider it an actual failure.
// 					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 						t.Fail()
// 					}
// 					fmt.Printf("Connection from key %v failed, continuing to next iteration \n", j)
// 					continue
// 				}
// 				fmt.Printf("res is: %v\n", res)
// 				resp, err := c.DeleteConnection(c.CompanyID, thisArg.ConnectionID)
// 				if err != nil {
// 					fmt.Println(err.Error())
// 					// if it's not a negative test, consider it an actual failure.
// 					if !(strings.Contains(j, "neg")) && !(strings.Contains(j, "Neg")) {
// 						fmt.Println("Failed Successfully")
// 						t.Fail()
// 					}
// 					// stop iteration of loop if connection create failed.
// 					fmt.Printf("Connection from key %v failed, continuing to next iteration \n", j)
// 					continue
// 				}
// 				fmt.Printf("resp is: %q\n", resp)
// 			}
// 		}
// 	}
// }
