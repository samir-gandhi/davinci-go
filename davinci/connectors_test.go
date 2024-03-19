package davinci_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

func TestConnectors_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	args, _ := testDataConnections["params"].(map[string]davinci.Params)
	for i := range args {
		thisArgs := args[i]
		fmt.Printf("args[i] is %q\n", thisArgs)
		resp, err := c.ReadConnectors(&companyID, &thisArgs)
		if err != nil {
			fmt.Println(err.Error())
			if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
				t.Fatal()
			}
		}
		fmt.Printf("resp is: %v\n", resp)
	}
}
