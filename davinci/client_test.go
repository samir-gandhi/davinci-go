package davinci

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

type envs struct {
	PINGONE_USERNAME              string `json:"PINGONE_USERNAME"`
	PINGONE_PASSWORD              string `json:"PINGONE_PASSWORD"`
	PINGONE_ENVIRONMENT_ID        string `json:"PINGONE_ENVIRONMENT_ID"`
	PINGONE_REGION                string `json:"PINGONE_REGION"`
	PINGONE_TARGET_ENVIRONMENT_ID string `json:"PINGONE_TARGET_ENVIRONMENT_ID"`
}

func TestNewClient_GA(t *testing.T) {
	var host, username, password string
	jsonFile, err := os.Open("../local/env-ga.json")
	// jsonFile, err := os.Open("../local/env-ga.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		username = envs.PINGONE_USERNAME
		password = envs.PINGONE_PASSWORD
	} else {
		fmt.Println("File: ./local/env.json not found, \n trying env vars for PINGONE_USERNAME/PINGONE_PASSWORD")
		username = os.Getenv("PINGONE_USERNAME")
		password = os.Getenv("PINGONE_PASSWORD")
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
			cInput := ClientInput{
				HostURL:  hostStruct.host,
				Username: username,
				Password: password,
			}
			_, err := NewClient(&cInput)
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

func TestNewClient_V2_HostAndRegion(t *testing.T) {
	var host, username, password string
	jsonFile, err := os.Open("../local/env-v2.json")
	// jsonFile, err := os.Open("../local/env-ga.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		username = envs.PINGONE_USERNAME
		password = envs.PINGONE_PASSWORD
	} else {
		fmt.Println("File: ./local/env.json not found, \n trying env vars for PINGONE_USERNAME/PINGONE_PASSWORD")
		username = os.Getenv("PINGONE_USERNAME")
		password = os.Getenv("PINGONE_PASSWORD")
	}
	emptyVars := username == "" || password == ""
	if emptyVars {
		log.Panicf("Missing Required Vars")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	var tests = map[string]struct {
		host   string
		region string
	}{
		"regionOnly":     {"", "NorthAmerica"},
		"badRegionNeg":   {"", "Europe"},
		"emptyStringNeg": {"", "NorthAmerica"},
		"testNeg":        {"https://badhost.io/v1", ""},
	}
	for name, hostStruct := range tests {
		testName := name
		t.Run(testName, func(t *testing.T) {
			cInput := ClientInput{
				HostURL:       hostStruct.host,
				Username:      username,
				Password:      password,
				PingOneRegion: hostStruct.region,
			}
			msg := fmt.Sprintf("\nGot client successfully, for test: %v\n", testName)
			_, err := NewClient(&cInput)
			// if client.Token == "" {
			// 	msg = fmt.Sprintf("\nNewClient Failed, no AccessToken for test: %v\n", testName)
			// }
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
func TestNewClient_V2_SSO(t *testing.T) {
	var region, username, password, p1SSOEnv, companyId string
	jsonFile, err := os.Open("../local/env-v2-sso.json")
	// jsonFile, err := os.Open("../local/env-v2-sso.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		username = envs.PINGONE_USERNAME
		password = envs.PINGONE_PASSWORD
		companyId = envs.PINGONE_TARGET_ENVIRONMENT_ID
		p1SSOEnv = envs.PINGONE_ENVIRONMENT_ID
	} else {
		fmt.Println("File: ./local/env-v2-sso.json not found, \n trying env vars for PINGONE_USERNAME/PINGONE_PASSWORD")
		username = os.Getenv("PINGONE_USERNAME")
		password = os.Getenv("PINGONE_PASSWORD")
		p1SSOEnv = os.Getenv("PINGONE_ENVIRONMENT_ID")
		region = os.Getenv("PINGONE_REGION")
		companyId = os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	}
	emptyVars := username == "" || password == "" || p1SSOEnv == ""
	if emptyVars {
		log.Panicf("Missing Required Vars")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	var tests = map[string]ClientInput{
		"correct": {
			HostURL:         "https://orchestrate-api.pingone.com/v1",
			Username:        username,
			Password:        password,
			PingOneSSOEnvId: p1SSOEnv,
		},
		"fromEnv": {
			PingOneRegion:   region,
			Username:        username,
			Password:        password,
			PingOneSSOEnvId: p1SSOEnv,
		},
		"emptyStringNeg": {
			HostURL:         "host",
			Username:        username,
			Password:        password,
			PingOneSSOEnvId: p1SSOEnv,
		},
		"badhostNeg": {
			HostURL:         "https://badhost.io/v1",
			Username:        username,
			Password:        password,
			PingOneSSOEnvId: p1SSOEnv,
		},
	}
	for name, inputStruct := range tests {
		testName := name
		t.Run(testName, func(t *testing.T) {
			client, err := NewClient(&inputStruct)
			if companyId != "" {
				client.CompanyID = companyId
			}
			msg := fmt.Sprintf("\nGot client successfully, for test: %v\n", testName)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(testName, "neg")) && !(strings.Contains(testName, "Neg")) {
					msg = fmt.Sprintf("failed to make client with host: %v \n Error is: %v", region, err)
					t.Fail()
				}
			}
			fmt.Println(msg)
		})
	}
}

func newTestClient() (*APIClient, error) {
	var region, username, password, p1SSOEnv, companyId string
	jsonFile, err := os.Open("../local/env-v2-sso.json")
	// jsonFile, err := os.Open("../local/env-v2-sso.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		username = envs.PINGONE_USERNAME
		password = envs.PINGONE_PASSWORD
		region = envs.PINGONE_REGION
		companyId = envs.PINGONE_TARGET_ENVIRONMENT_ID
		p1SSOEnv = envs.PINGONE_ENVIRONMENT_ID
	} else {
		fmt.Println("File: ./local/env-v2-sso.json not found, \n trying env vars for PINGONE_USERNAME/PINGONE_PASSWORD")
		username = os.Getenv("PINGONE_USERNAME")
		password = os.Getenv("PINGONE_PASSWORD")
		p1SSOEnv = os.Getenv("PINGONE_ENVIRONMENT_ID")
		region = os.Getenv("PINGONE_REGION")
		companyId = os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	}
	cInput := ClientInput{
		PingOneRegion:   region,
		Username:        username,
		Password:        password,
		PingOneSSOEnvId: p1SSOEnv,
	}
	client, err := NewClient(&cInput)
	fmt.Println("clientcompany: ", client.CompanyID)
	if companyId != "" {
		client.CompanyID = companyId
	}
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	return client, nil
}
