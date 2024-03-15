package test

import (
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

var testDataApps = map[string]dv.AppUpdate{
	"defaultFlow": {
		Name: "goclient-base-app",
		Oauth: &dv.Oauth{
			Enabled: true,
			Values: &dv.OauthValues{
				Enabled:       true,
				AllowedScopes: []string{"openid", "profile", "flow_analytics"},
			},
		},
		Policies: []dv.Policy{{
			PolicyFlows: []dv.PolicyFlow{{
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
		UserPortal: &dv.UserPortal{
			Values: &dv.UserPortalValues{
				UpTitle: "thisIsTitle",
			},
		},
	},
	"plain": {
		Name: "goclient-plain-app",
		Oauth: &dv.Oauth{
			Enabled: true,
			Values: &dv.OauthValues{
				Enabled:       true,
				AllowedScopes: []string{"openid", "profile", "flow_analytics"},
			},
		},
		Policies: []dv.Policy{},
		UserPortal: &dv.UserPortal{
			Values: &dv.UserPortalValues{
				UpTitle: "thisIsTitle",
			},
		},
	},
}

var testDataAppPolicy = map[string]dv.Policy{
	"basePolicy": {
		PolicyFlows: []dv.PolicyFlow{{
			FlowID:    "",
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
	},
}

func makeTestApps(identifier string) map[string]dv.AppUpdate {
	for key, value := range testDataApps {
		value.Name = value.Name + "-" + identifier
		testDataApps[key] = value
	}
	return testDataApps
}
