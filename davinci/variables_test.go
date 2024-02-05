package davinci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/samir-gandhi/davinci-client-go/tools"
)

var testDataVars = map[string]interface{}{
	"params": map[string]Params{
		"a":    {"0", "1", nil},
		"bNeg": {"100", "10", nil},
		"c":    {},
	},
}

func TestReadVariables(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataVars["params"].(map[string]Params)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			msg := ""
			resp, err := c.ReadVariables(c.CompanyID, &thisArg)
			if err != nil {
				msg = fmt.Sprintf("Failed Successfully: %v\n", err.Error())
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					msg = fmt.Sprintf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				var l []string
				for j, _ := range resp {
					l = append(l, j)
				}
				msg = fmt.Sprintf("Vars Returned Successfully\n vars returned is: %v\n", l)
			}
			fmt.Println(msg)
		})
	}
}

func testReadVariable_inputs() []string {
	a := []string{fmt.Sprint("flow##SK##flowInstance")}
	return a
}

func TestReadVariable(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args := testReadVariable_inputs()
	for _, thisArg := range args {
		t.Run(thisArg, func(t *testing.T) {
			msg := ""
			resp, err := c.ReadVariable(c.CompanyID, thisArg)
			if err != nil {
				msg = fmt.Sprintf("Failed Successfully: %v\n", err.Error())
				if !(strings.Contains(thisArg, "neg")) && !(strings.Contains(thisArg, "Neg")) {
					msg = fmt.Sprintf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				var l []string
				for j, _ := range resp {
					l = append(l, j)
				}
				msg = fmt.Sprintf("Vars Returned Successfully\n vars returned is: %v\n", l)
			}
			fmt.Println(msg)
		})
	}
}

func testCreateVariables_inputs(resourceName string) (varMap map[string]VariablePayload) {
	varMap = map[string]VariablePayload{
		"company": {
			Name:        fmt.Sprintf("company-%s", resourceName),
			Context:     "company",
			Value:       fmt.Sprintf(`value-%s`, resourceName),
			Description: fmt.Sprintf("description-%s", resourceName),
			Type:        "string",
			Mutable:     true,
		},
		"flow": {
			Name:        fmt.Sprintf("flow-%s", resourceName),
			Context:     "flowInstance",
			Value:       fmt.Sprintf(`value-%s`, resourceName),
			Description: fmt.Sprintf("description-%s", resourceName),
			Type:        "string",
			Mutable:     true,
		},
		"flowNeg": {
			Name:        fmt.Sprintf("flow-%s", resourceName),
			Context:     "badContext",
			Value:       fmt.Sprintf(`value-%s`, resourceName),
			Description: fmt.Sprintf("description-%s", resourceName),
			Type:        "string",
			Mutable:     true,
		},
	}
	return varMap
}

func TestCreateVariables(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	resourceName := tools.RandomString(10)
	args := testCreateVariables_inputs(resourceName)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			msg := ""
			resp, err := c.CreateVariable(c.CompanyID, &thisArg)
			if err != nil {
				msg = fmt.Sprintf("Failed Successfully: %v\n", err.Error())
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					msg = fmt.Sprintf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				msg = fmt.Sprintf("Var created successfully\n vars returned is: %v\n", resp)
			}
			fmt.Println(msg)
		})
	}
}

func testUpdateVariables_inputs(resourceName string) (varMap map[string]VariablePayload) {
	varMap = map[string]VariablePayload{
		"company": {
			Name:        fmt.Sprintf("company-%s", resourceName),
			Context:     "company",
			Value:       fmt.Sprintf(`value-%s`, resourceName),
			Description: fmt.Sprintf("description-%s", resourceName),
			Type:        "string",
			Mutable:     true,
		},
		"flow": {
			Name:        fmt.Sprintf("flow-%s", resourceName),
			Context:     "flowInstance",
			Value:       fmt.Sprintf(`value-%s`, resourceName),
			Description: fmt.Sprintf("description-%s", resourceName),
			Type:        "string",
			Mutable:     true,
		},
		"flowNeg": {
			Name:        fmt.Sprintf("flow-%s", resourceName),
			Context:     "badContext",
			Value:       fmt.Sprintf(`value-%s`, resourceName),
			Description: fmt.Sprintf("description-%s", resourceName),
			Type:        "string",
			Mutable:     true,
		},
	}
	return varMap
}

func TestUpdateVariables(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	resourceName := tools.RandomString(10)
	args := testUpdateVariables_inputs(resourceName)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			msg := ""
			resp, err := c.CreateVariable(c.CompanyID, &thisArg)
			if err != nil {
				msg = fmt.Sprintf("Failed Successfully: %v\n", err.Error())
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					msg = fmt.Sprintf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				// msg = fmt.Sprintf("Var created successfully\n vars returned is: %v\n", resp)
				thisArg.Description = thisArg.Description + tools.RandomString(10)
				update, err := c.UpdateVariable(c.CompanyID, &thisArg)
				if err != nil {
					msg = fmt.Sprintf("Failed Successfully: %v\n", err.Error())
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						msg = fmt.Sprintf("failed to get flows with params: %v \n Error is: %v", args, err)
						t.Fail()
					}
				}
				if update != nil {
					msg = fmt.Sprintf("Var updated successfully\n var returned is: %v\n", resp)
				}
			}
			fmt.Println(msg)
		})
	}
}

func testDeleteVariables_inputs(resourceName string) (varMap map[string]VariablePayload) {
	varMap = map[string]VariablePayload{
		"company": {
			Name:        fmt.Sprintf("company-%s", resourceName),
			Context:     "company",
			Value:       fmt.Sprintf(`value-%s`, resourceName),
			Description: fmt.Sprintf("description-%s", resourceName),
			Type:        "string",
			Mutable:     true,
		},
		"flow": {
			Name:        fmt.Sprintf("flow-%s", resourceName),
			Context:     "flowInstance",
			Value:       fmt.Sprintf(`value-%s`, resourceName),
			Description: fmt.Sprintf("description-%s", resourceName),
			Type:        "string",
			Mutable:     true,
		},
		"flowNeg": {
			Name:        fmt.Sprintf("flow-%s", resourceName),
			Context:     "badContext",
			Value:       fmt.Sprintf(`value-%s`, resourceName),
			Description: fmt.Sprintf("description-%s", resourceName),
			Type:        "string",
			Mutable:     true,
		},
	}
	return varMap
}

func TestDeleteVariables(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	resourceName := tools.RandomString(10)
	args := testDeleteVariables_inputs(resourceName)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			msg := ""
			resp, err := c.CreateVariable(c.CompanyID, &thisArg)
			if err != nil {
				msg = fmt.Sprintf("Failed Successfully: %v\n", err.Error())
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					msg = fmt.Sprintf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fail()
				}
			}
			if resp != nil {
				// msg = fmt.Sprintf("Var created successfully\n vars returned is: %v\n", resp)
				update, err := c.DeleteVariable(c.CompanyID, fmt.Sprintf(`%s##SK##%s`, thisArg.Name, thisArg.Context))
				if err != nil {
					msg = fmt.Sprintf("Failed Successfully: %v\n", err.Error())
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						msg = fmt.Sprintf("failed to get flows with params: %v \n Error is: %v", args, err)
						t.Fail()
					}
				}
				if update != nil {
					msg = fmt.Sprintf("Var updated successfully\n var returned is: %v\n", resp)
				}
			}
			fmt.Println(msg)
		})
	}
}
