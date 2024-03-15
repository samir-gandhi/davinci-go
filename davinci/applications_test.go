package davinci_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

var testDataApps = map[string]interface{}{
	"params": map[string]davinci.Params{
		"a": {
			Page:  "0",
			Limit: "1",
		},
		"bNeg": {
			Page:  "100",
			Limit: "10",
		},
		"c": {},
	},
	"appsCreate": map[string]davinci.App{
		"aCreate": {
			Name: "aCreate",
		},
		"bCreate": {
			Name: "B CREATE",
		},
		"cCreateNeg": {
			Name: "",
		},
	},
	"appsRead": map[string]davinci.App{
		"aCreate": {
			Name: "aCreate",
		},
		"bCreate": {
			Name: "B CREATE",
		},
		"cCreateNeg": {
			Name: "",
		},
	},
	"appsUpdate": map[string]davinci.AppUpdate{
		"aCreate": {
			Name: "aCreate",
			Oauth: &davinci.Oauth{
				Enabled: true,
				Values: &davinci.OauthValues{
					Enabled:       true,
					AllowedScopes: []string{"openid", "profile", "flow_analytics"},
				},
			},
			Policies: []davinci.Policy{{
				PolicyFlows: []davinci.PolicyFlow{{
					FlowID:    "1764a19731a067d8b56f0c2d250cd9ea",
					VersionID: -1,
					Weight: func() *int {
						i := 100
						return &i
					}(),
				}},
				Name: func() *string {
					s := "aCreatePolicy"
					return &s
				}(),
			}},
			UserPortal: &davinci.UserPortal{
				Values: &davinci.UserPortalValues{
					UpTitle: "thisIsTitle",
				},
			},
		},
		"bCreate": {
			Name: "BCREATE",
			Saml: &davinci.Saml{
				Values: &davinci.SamlValues{
					Enabled: true,
					RedirectURI: func() *string {
						s := "https://example.com"
						return &s
					}(),
				},
			},
		},
		"cCreateNeg": {
			Name: "",
		},
	},
}

func TestApplications_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataApps["params"].(map[string]davinci.Params)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			msg := ""
			resp, err := c.ReadApplications(c.CompanyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					msg = fmt.Sprintf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				//msg = fmt.Sprintf("Apps Returned Successfully\n appId is: %+v \n", *resp[0].AppID)
				msg = fmt.Sprintf("Apps Returned Successfully\n appId is: %#v \n", resp)
			}
			fmt.Println(msg)
		})
	}
}

func TestApplication_Create(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataApps["appsCreate"].(map[string]davinci.App)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			msg := ""
			resp, err := c.CreateApplication(c.CompanyID, thisArg.Name)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				msg = fmt.Sprintf("Apps Returned Successfully\n appId is: %+v \n", resp.AppID)
			}
			fmt.Println(msg)
		})
	}
}

func TestApplication_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataApps["appsRead"].(map[string]davinci.App)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			msg := ""
			resp, err := c.CreateApplication(c.CompanyID, thisArg.Name)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				res, err := c.ReadApplication(c.CompanyID, *resp.AppID)
				if err != nil {
					fmt.Println(err.Error())
					msg = fmt.Sprint("Failed Successfully\n")
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
						t.Fail()
					}
				}
				if res != nil {
					msg = fmt.Sprintf("Apps Returned Successfully\n appId is: %+v \n", res.AppID)
				}
			}
			fmt.Println(msg)
		})
	}
}

func TestApplication_Update(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataApps["appsUpdate"].(map[string]davinci.AppUpdate)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			msg := ""
			resp, err := c.CreateApplication(c.CompanyID, thisArg.Name)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					msg = fmt.Sprintf("App Create failed with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				thisArg.AppID = resp.AppID
				res, err := c.UpdateApplication(c.CompanyID, &thisArg)
				if err != nil {
					fmt.Println(err.Error())
					msg = fmt.Sprint("Failed Successfully\n")
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
						t.Fail()
					}
				}
				if res != nil {
					msg = fmt.Sprintf("Apps Returned Successfully\n appId is: %+v \n", resp.AppID)
				}
			}
			fmt.Println(msg)
		})
	}
}

func TestApplication_Delete(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataApps["appsUpdate"].(map[string]davinci.AppUpdate)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			msg := ""
			resp, err := c.CreateApplication(c.CompanyID, thisArg.Name)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					msg = fmt.Sprintf("App Create failed with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				thisArg.AppID = resp.AppID
				res, err := c.DeleteApplication(c.CompanyID, *thisArg.AppID)
				fmt.Println(res)
				if err != nil && *res.Message != "App successfully removed" {
					fmt.Println(err.Error())
					msg = fmt.Sprint("Failed Successfully\n")
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						msg = fmt.Sprintf("failed with params: %v \n Error is: %v", args, err)
						t.Fail()
					}
				}
				if res != nil {
					msg = fmt.Sprintf("App Deleted Successfully\n message is: %v \n", res)
				}
			}
			fmt.Println(msg)
		})
	}
}
