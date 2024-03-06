package davinci_test

// import (
// 	"fmt"

// 	"strings"
// 	"testing"
// )

// func TestReadConnectors(t *testing.T) {
// 	c, err := newTestClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	args, _ := testDataConnections["params"].(map[string]Params)
// 	for i := range args {
// 		thisArgs := args[i]
// 		fmt.Printf("args[i] is %q\n", thisArgs)
// 		resp, err := c.ReadConnectors(&c.CompanyID, &thisArgs)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
// 				t.Fail()
// 			}
// 		}
// 		fmt.Printf("resp is: %v\n", resp)
// 	}
// }
