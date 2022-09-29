package davinci

import (
	"fmt"
	"testing"
)

// testData for Roles functions
var testDataRoles = map[string]interface{}{
	"params": map[string]Params{
		"a": {"1", "10", nil},
		"b": {"1000", "10", nil},
		"c": {},
	},
	"rolesCreate": map[string]RoleCreate{
		"a": {"roleA"},
		"b": {"roleB"},
		"c": {},
	},
}

/* TEST PLAN:
* - Read all roles
* - Create new roles
* - Read new roles
* - Update new roles
* - Delete new roles
 */

// Gets an array of all roles for a company
func TestReadRoles(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataRoles["params"].(map[string]Params)
	for i := range args {
		thisArgs := args[i]
		fmt.Printf("role[i] is %q\n", thisArgs)
		roles, err := c.ReadRoles(&c.CompanyID, &thisArgs)
		if err != nil {

			fmt.Println(err.Error())
			t.Fail()
		}
		fmt.Printf("roles is: %v\n", roles)
	}
}

func TestCreateRole(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataRoles["rolesCreate"].(map[string]RoleCreate)
	for i := range args {
		thisRole := args[i]
		fmt.Printf("role[i] is %q\n", thisRole)
		roles, err := c.CreateRole(&c.CompanyID, &thisRole)
		if err != nil {
			fmt.Println(err.Error())
			t.Fail()
		}
		fmt.Printf("roles is: %q\n", roles)
	}
}

func TestReadRole(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataRoles["rolesCreate"].(map[string]RoleCreate)
	for i := range args {
		thisArgs := args[i]
		fmt.Printf("role[i] is %q\n", thisArgs)
		roles, err := c.ReadRole(&c.CompanyID, thisArgs.Name)
		if err != nil {
			fmt.Println(err.Error())
			t.Fail()
		}
		fmt.Printf("roles is: %v\n", roles)
	}
}
func TestUpdateRole(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataRoles["rolesCreate"].(map[string]RoleCreate)
	for i := range args {
		thisArgs := args[i]
		fmt.Printf("role[i] is %q\n", thisArgs)
		if thisArgs.Name != "" {
			role, err := c.ReadRole(&c.CompanyID, thisArgs.Name)
			if err != nil {
				t.FailNow()
			}
			role.Description = "tempdesc"
			thisRoleUpdate := RoleUpdate{
				"tempdesc",
				role.Policy,
			}
			roleResp, err := c.UpdateRole(&c.CompanyID, thisArgs.Name, &thisRoleUpdate)
			if err != nil {

				fmt.Println(err.Error())
				t.Fail()
			}
			fmt.Printf("roleResp is: %v\n", roleResp)
		}
	}
}

func TestDeleteRole(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	args, _ := testDataRoles["rolesCreate"].(map[string]RoleCreate)
	for i := range args {
		thisRole := args[i]
		fmt.Printf("role[i] is %q\n", thisRole)
		role, err := c.DeleteRole(&c.CompanyID, thisRole.Name)
		if err != nil {
			fmt.Println(err.Error())
			t.Fail()
		}
		fmt.Printf("roleResp is: %q\n", role)
	}
}
