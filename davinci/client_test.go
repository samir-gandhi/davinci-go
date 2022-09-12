package davinci

import (
	"encoding/json"
	"fmt"
	"github.com/samir-gandhi/davinci-client-go/tools"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

type envs struct {
	DAVINCI_USERNAME  string `json:"DAVINCI_USERNAME"`
	DAVINCI_PASSWORD  string `json:"DAVINCI_PASSWORD"`
	DAVINCI_COMPANYID string `json:"DAVINCI_COMPANYID"`
	DAVINCI_HOST      string `json:"DAVINCI_HOST"`
}

func TestNewClient(t *testing.T) {
	var host, username, password string
	jsonFile, err := os.Open("../local/env.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		host = envs.DAVINCI_HOST
		username = envs.DAVINCI_USERNAME
		password = envs.DAVINCI_PASSWORD
	} else {
		fmt.Println("File: ./local/env.json not found, \n trying env vars for DAVINCI_USERNAME/DAVINCI_PASSWORD")
		host = os.Getenv("DAVINCI_HOST")
		username = os.Getenv("DAVINCI_USERNAME")
		password = os.Getenv("DAVINCI_PASSWORD")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	var tests = map[string]struct {
		host string
	}{
		"default":     {"https://api.singularkey.com/v1"},
		"nil":         {},
		"emptystring": {""},
		"testNeg":     {"https://badhost.io/v1"},
	}

	for name, hostStruct := range tests {
		testName := name
		t.Run(testName, func(t *testing.T) {
			_, err := NewClient(&hostStruct.host, &username, &password)
			msg := fmt.Sprintf("\nGot client successfully, for test: %v\n", testName)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(testName, "neg")) && !(strings.Contains(testName, "Neg")) {
					msg = fmt.Sprintf("failed to make client with host: %v \n Error is: %v", host, err)
					t.Fail()
				}
			}
			fmt.Printf(msg)
		})
	}
}

func newTestClient() (*Client, error) {
	tools.PrintHeader("Initializing Test Client")
	defer tools.PrintFooter("Initializing Test Client")
	// var host string
	var username, password string
	jsonFile, err := os.Open("../local/env.json")
	var envs envs
	var companyid string
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		username = envs.DAVINCI_USERNAME
		password = envs.DAVINCI_PASSWORD
		companyid = envs.DAVINCI_COMPANYID
	} else {
		fmt.Println("File: ./local/env.json not found, \n trying env vars for DAVINCI_USERNAME/DAVINCI_PASSWORD")
		username = os.Getenv("DAVINCI_USERNAME")
		password = os.Getenv("DAVINCI_PASSWORD")
		companyid = os.Getenv("DAVINCI_COMPANYID")
	}
	client, err := NewClient(nil, &username, &password)
	if companyid != "" {
		client.CompanyID = companyid
	}
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	return client, nil
}
