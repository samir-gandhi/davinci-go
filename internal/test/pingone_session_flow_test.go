package test

import (
	"testing"
)

func TestPingOneSessionFlowApp(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	apps := makeTestApps(RandString(10))
	app := apps["noPolicy"]
	policy := testDataAppPolicy["basePolicy"]
	app.Policies = append(app.Policies, policy)
	// using flow that is manually created in feature flag env.
	app.Policies[0].PolicyFlows[0].FlowID = "5c32a89d4093b0eba7292ecafdd6b0e9"
	c.CreateInitializedApplication(&c.CompanyID, &app)
}
func TestNoPolicyApp(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	apps := makeTestApps(RandString(10))
	app := apps["noPolicy"]
	c.CreateInitializedApplication(&c.CompanyID, &app)
}
