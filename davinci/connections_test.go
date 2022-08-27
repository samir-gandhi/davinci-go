package davinci

import (
	"fmt"
	"strings"
	"testing"
)

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
		},
		"bUpdate": {
			Name:        "connectionBUpdate",
			ConnectorID: "pingOneMfaConnector",
			CustomerID:  "1234",
		},
		"cUpdateneg": {
			Name:        "",
			ConnectorID: "pingOneMfaConnector",
		},
		"dUpdateneg": {
			Name:        "connectionDUpdate",
			ConnectorID: "",
		},
	},
	// "connectionsUpdateProperties": Properties{
	// 	"aUpdate": {
	// 		Name:        "connectionAUpdate",
	// 		ConnectorID: "pingOneMfaConnector",
	// 	},
	// 	"bUpdate": {
	// 		Name:        "connectionBUpdate",
	// 		ConnectorID: "pingOneMfaConnector",
	// 		CustomerID:  "1234",
	// 	},
	// 	"cUpdateneg": {
	// 		Name:        "",
	// 		ConnectorID: "pingOneMfaConnector",
	// 	},
	// 	"dUpdateneg": {
	// 		Name:        "connectionDUpdate",
	// 		ConnectorID: "",
	// 	},
	// },
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
			t.Fail()
		}
		fmt.Printf("resp is: %v\n", resp)
	}
}

func TestCreateConnection(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataConnections["connectionsCreate"].(map[string]Connection)
	for i := range args {
		thisArg := args[i]
		fmt.Printf("args[i] is %q\n", thisArg)
		resp, err := c.CreateConnection(&c.CompanyID, &thisArg)
		if err != nil {
			fmt.Println(err.Error())
			// if it's not a negative test, consider it an actual failure.
			if !(strings.Contains(i, "neg")) {
				t.Fail()
			}
		}
		fmt.Printf("resp is: %q\n", resp)
	}
}

func TestReadConnection(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataConnections["connectionsRead"].(map[string]Connection)
	for i := range args {
		thisArg := args[i]
		fmt.Printf("args[i] is %q\n", thisArg)
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
		fmt.Printf("resp is: %q\n", resp)
		res, err := c.ReadConnection(&c.CompanyID, resp.ConnectionID)
		if err != nil {
			fmt.Println(err.Error())
			// if it's not a negative test, consider it an actual failure.
			if !(strings.Contains(i, "neg")) {
				t.Fail()
			}
		}
		fmt.Printf("res is: %v\n", res)
	}
}

func TestUpdateConnection(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataConnections["connectionsRead"].(map[string]Connection)
	for i := range args {
		thisArg := args[i]
		fmt.Printf("args[i] is %q\n", thisArg)
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
		fmt.Printf("resp is: %q\n", resp)
		res, err := c.ReadConnection(&c.CompanyID, resp.ConnectionID)
		if err != nil {
			fmt.Println(err.Error())
			// if it's not a negative test, consider it an actual failure.
			if !(strings.Contains(i, "neg")) {
				t.Fail()
			}
		}
		fmt.Printf("res is: %v\n", res)
	}
}
