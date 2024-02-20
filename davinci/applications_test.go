package davinci

// import (
// 	"fmt"
// 	"strings"
// 	"testing"
// )

// var testDataApps = map[string]interface{}{
// 	"params": map[string]Params{
// 		"a":    {"0", "1", nil},
// 		"bNeg": {"100", "10", nil},
// 		"c":    {},
// 	},
// 	"appsCreate": map[string]App{
// 		"aCreate": {
// 			Name: "aCreate",
// 		},
// 		"bCreate": {
// 			Name: "B CREATE",
// 		},
// 		"cCreateNeg": {
// 			Name: "",
// 		},
// 	},
// 	"appsRead": map[string]App{
// 		"aCreate": {
// 			Name: "aCreate",
// 		},
// 		"bCreate": {
// 			Name: "B CREATE",
// 		},
// 		"cCreateNeg": {
// 			Name: "",
// 		},
// 	},
// 	"appsUpdate": map[string]AppUpdate{
// 		"aCreate": {
// 			Name: "aCreate",
// 			Oauth: &Oauth{
// 				Enabled: true,
// 				Values: &OauthValues{
// 					Enabled:       true,
// 					AllowedScopes: []string{"openid", "profile", "flow_analytics"},
// 				},
// 			},
// 			Policies: []Policy{{
// 				PolicyFlows: []PolicyFlow{{
// 					FlowID:    "1764a19731a067d8b56f0c2d250cd9ea",
// 					VersionID: -1,
// 					Weight:    100,
// 				}},
// 				Name: "aCreatePolicy",
// 			}},
// 			UserPortal: &UserPortal{
// 				Values: &UserPortalValues{
// 					UpTitle: "thisIsTitle",
// 				},
// 			},
// 		},
// 		"bCreate": {
// 			Name: "BCREATE",
// 			Saml: &Saml{
// 				Values: &SamlValues{
// 					Enabled:     true,
// 					RedirectURI: "https://example.com",
// 				},
// 			},
// 		},
// 		"cCreateNeg": {
// 			Name: "",
// 		},
// 	},
// }

// func TestReadApplications(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	args, _ := testDataApps["params"].(map[string]Params)
// 	for i, thisArg := range args {
// 		testName := i
// 		t.Run(testName, func(t *testing.T) {
// 			msg := ""
// 			resp, err := c.ReadApplications(c.CompanyID, &thisArg)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				msg = fmt.Sprint("Failed Successfully\n")
// 				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 					msg = fmt.Sprintf("failed to get flows with params: %v \n Error is: %v", args, err)
// 					t.Fail()
// 				}
// 			}
// 			if resp != nil {
// 				msg = fmt.Sprintf("Apps Returned Successfully\n appId is: %+v \n", resp[0].AppID)
// 			}
// 			fmt.Println(msg)
// 		})
// 	}
// }

// func TestCreateApplication(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	args, _ := testDataApps["appsCreate"].(map[string]App)
// 	for i, thisArg := range args {
// 		testName := i
// 		t.Run(testName, func(t *testing.T) {
// 			msg := ""
// 			resp, err := c.CreateApplication(c.CompanyID, thisArg.Name)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				msg = fmt.Sprint("Failed Successfully\n")
// 				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 					msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
// 					t.Fail()
// 				}
// 			}
// 			if resp != nil {
// 				msg = fmt.Sprintf("Apps Returned Successfully\n appId is: %+v \n", resp.AppID)
// 			}
// 			fmt.Println(msg)
// 		})
// 	}
// }

// func TestReadApplication(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	args, _ := testDataApps["appsRead"].(map[string]App)
// 	for i, thisArg := range args {
// 		testName := i
// 		t.Run(testName, func(t *testing.T) {
// 			msg := ""
// 			resp, err := c.CreateApplication(c.CompanyID, thisArg.Name)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				msg = fmt.Sprint("Failed Successfully\n")
// 				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 					msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
// 					t.Fail()
// 				}
// 			}
// 			if resp != nil {
// 				res, err := c.ReadApplication(c.CompanyID, resp.AppID)
// 				if err != nil {
// 					fmt.Println(err.Error())
// 					msg = fmt.Sprint("Failed Successfully\n")
// 					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 						msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
// 						t.Fail()
// 					}
// 				}
// 				if res != nil {
// 					msg = fmt.Sprintf("Apps Returned Successfully\n appId is: %+v \n", res.AppID)
// 				}
// 			}
// 			fmt.Println(msg)
// 		})
// 	}
// }

// func TestUpdateApplication(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	args, _ := testDataApps["appsUpdate"].(map[string]AppUpdate)
// 	for i, thisArg := range args {
// 		testName := i
// 		t.Run(testName, func(t *testing.T) {
// 			msg := ""
// 			resp, err := c.CreateApplication(c.CompanyID, thisArg.Name)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				msg = fmt.Sprint("Failed Successfully\n")
// 				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 					msg = fmt.Sprintf("App Create failed with params: %v \n Error is: %v", args, err)
// 					t.Fail()
// 				}
// 			}
// 			if resp != nil {
// 				thisArg.AppID = resp.AppID
// 				res, err := c.UpdateApplication(c.CompanyID, &thisArg)
// 				if err != nil {
// 					fmt.Println(err.Error())
// 					msg = fmt.Sprint("Failed Successfully\n")
// 					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 						msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
// 						t.Fail()
// 					}
// 				}
// 				if res != nil {
// 					msg = fmt.Sprintf("Apps Returned Successfully\n appId is: %+v \n", resp.AppID)
// 				}
// 			}
// 			fmt.Println(msg)
// 		})
// 	}
// }

// func TestCreateInitializedApplication(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	args, _ := testDataApps["appsUpdate"].(map[string]AppUpdate)
// 	for i, thisArg := range args {
// 		testName := i
// 		t.Run(testName, func(t *testing.T) {
// 			msg := ""
// 			resp, err := c.CreateInitializedApplication(c.CompanyID, &thisArg)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				msg = fmt.Sprint("Failed Successfully\n")
// 				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 					msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
// 					t.Fail()
// 				}
// 			}
// 			if resp != nil {
// 				msg = fmt.Sprintf("Apps Created Successfully\n appId is: %+v \n", resp.AppID)
// 				// Payload Validation:
// 				if resp.Oauth.Values.ClientSecret == "" {
// 					msg = fmt.Sprintf("failed with params: %v \n No Client Secret found. Error is: %v", args, err)
// 					t.Fail()
// 				}
// 			}
// 			fmt.Println(msg)
// 		})
// 	}
// }

// func TestDeleteApplication(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	args, _ := testDataApps["appsUpdate"].(map[string]AppUpdate)
// 	for i, thisArg := range args {
// 		testName := i
// 		t.Run(testName, func(t *testing.T) {
// 			msg := ""
// 			resp, err := c.CreateApplication(c.CompanyID, thisArg.Name)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				msg = fmt.Sprint("Failed Successfully\n")
// 				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 					msg = fmt.Sprintf("App Create failed with params: %v \n Error is: %v", args, err)
// 					t.Fail()
// 				}
// 			}
// 			if resp != nil {
// 				thisArg.AppID = resp.AppID
// 				res, err := c.DeleteApplication(c.CompanyID, thisArg.AppID)
// 				fmt.Println(res)
// 				if err != nil && res.Message != "App successfully removed" {
// 					fmt.Println(err.Error())
// 					msg = fmt.Sprint("Failed Successfully\n")
// 					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 						msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
// 						t.Fail()
// 					}
// 				}
// 				if res != nil {
// 					msg = fmt.Sprintf("App Deleted Successfully\n message is: %v \n", res)
// 				}
// 			}
// 			fmt.Println(msg)
// 		})
// 	}
// }
