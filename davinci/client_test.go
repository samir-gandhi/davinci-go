package davinci_test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

type envs struct {
	PINGONE_USERNAME              string `json:"PINGONE_USERNAME"`
	PINGONE_PASSWORD              string `json:"PINGONE_PASSWORD"`
	PINGONE_ENVIRONMENT_ID        string `json:"PINGONE_ENVIRONMENT_ID"`
	PINGONE_REGION                string `json:"PINGONE_REGION"`
	PINGONE_TARGET_ENVIRONMENT_ID string `json:"PINGONE_TARGET_ENVIRONMENT_ID"`
	PINGONE_DAVINCI_ACCESS_TOKEN  string `json:"PINGONE_DAVINCI_ACCESS_TOKEN"`
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
		errJ := json.Unmarshal(byteValue, &envs)
		if errJ != nil {
			log.Fatalf("failed to unmarshal json %v: ", errJ)
		}
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
			cInput := davinci.ClientInput{
				HostURL:  hostStruct.host,
				Username: username,
				Password: password,
			}
			_, err := davinci.NewClient(&cInput)
			msg := fmt.Sprintf("\nGot client successfully, for test: %v\n", testName)
			if err != nil {
				fmt.Println(err.Error())
				msg = "Failed Successfully\n"
				log.Printf("%s", msg)
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(testName, "neg")) && !(strings.Contains(testName, "Neg")) {
					log.Printf("failed to make client with host: %v \n Error is: %v", host, err)
					t.Fail()
				}
			}
			log.Printf("%s", msg)
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
		errJ := json.Unmarshal(byteValue, &envs)
		if errJ != nil {
			log.Fatalf("failed to unmarshal json %v: ", errJ)
		}
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
			cInput := davinci.ClientInput{
				HostURL:       hostStruct.host,
				Username:      username,
				Password:      password,
				PingOneRegion: hostStruct.region,
			}
			msg := fmt.Sprintf("\nGot client successfully, for test: %v\n", testName)
			_, err := davinci.NewClient(&cInput)
			// if client.Token == "" {
			// 	msg = fmt.Sprintf("\nNewClient Failed, no AccessToken for test: %v\n", testName)
			// }
			if err != nil {
				fmt.Println(err.Error())
				msg = "Failed Successfully\n"
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(testName, "neg")) && !(strings.Contains(testName, "Neg")) {
					msg = fmt.Sprintf("failed to make client with host: %v \n Error is: %v", host, err)
					t.Fail()
				}
			}
			log.Printf("%s", msg)
		})
	}
}
func TestNewClient_V2_SSO(t *testing.T) {
	var region, username, password, p1SSOEnv, companyId, accessToken, hostUrl string
	jsonFile, err := os.Open("../local/env-v2-sso.json")
	// jsonFile, err := os.Open("../local/env-v2-sso.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		errJ := json.Unmarshal(byteValue, &envs)
		if errJ != nil {
			log.Fatalf("failed to unmarshal json %v: ", errJ)
		}
		username = envs.PINGONE_USERNAME
		password = envs.PINGONE_PASSWORD
		companyId = envs.PINGONE_TARGET_ENVIRONMENT_ID
		p1SSOEnv = envs.PINGONE_ENVIRONMENT_ID
		accessToken = envs.PINGONE_DAVINCI_ACCESS_TOKEN
		// hostUrl = envs.PINGONE_DAVINCI_HOST_URL
	} else {
		fmt.Println("File: ./local/env-v2-sso.json not found, \n trying env vars for PINGONE_USERNAME/PINGONE_PASSWORD")
		username = os.Getenv("PINGONE_USERNAME")
		password = os.Getenv("PINGONE_PASSWORD")
		p1SSOEnv = os.Getenv("PINGONE_ENVIRONMENT_ID")
		region = os.Getenv("PINGONE_REGION")
		companyId = os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
		accessToken = os.Getenv("PINGONE_DAVINCI_ACCESS_TOKEN")
		hostUrl = os.Getenv("PINGONE_DAVINCI_HOST_URL")
	}
	userpw := username == "" || password == ""

	if (userpw && accessToken == "") || p1SSOEnv == "" {
		log.Panicf("Missing Required Vars")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	var tests = map[string]davinci.ClientInput{
		// "correct": {
		// 	HostURL:         "https://orchestrate-api.pingone.com/v1",
		// 	Username:        username,
		// 	Password:        password,
		// 	PingOneSSOEnvId: p1SSOEnv,
		// 	AccessToken:     accessToken,
		// },
		"fromEnv": {
			HostURL:         hostUrl,
			PingOneRegion:   region,
			Username:        username,
			Password:        password,
			PingOneSSOEnvId: p1SSOEnv,
			AccessToken:     accessToken,
		},
		"emptyStringNeg": {
			HostURL:         "host",
			Username:        username,
			Password:        password,
			PingOneSSOEnvId: p1SSOEnv,
			AccessToken:     accessToken,
		},
		"badhostNeg": {
			HostURL:         "https://badhost.io/v1",
			Username:        username,
			Password:        password,
			PingOneSSOEnvId: p1SSOEnv,
			AccessToken:     accessToken,
		},
	}
	for name, inputStruct := range tests {
		testName := name
		t.Run(testName, func(t *testing.T) {
			client, err := davinci.NewClient(&inputStruct)
			if companyId != "" {
				client.CompanyID = companyId
			}
			msg := fmt.Sprintf("\nGot client successfully, for test: %v\n", testName)
			if err != nil {
				fmt.Println(err.Error())
				msg = "Failed Successfully\n"
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

func newTestClient() (*davinci.APIClient, error) {
	var region, username, password, p1SSOEnv, companyId, accessToken, hostUrl string
	jsonFile, err := os.Open("../local/env-v2-sso.json")
	// jsonFile, err := os.Open("../local/env-v2-sso.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		errJ := json.Unmarshal(byteValue, &envs)
		if errJ != nil {
			log.Fatalf("failed to unmarshal json %v: ", errJ)
		}
		username = envs.PINGONE_USERNAME
		password = envs.PINGONE_PASSWORD
		companyId = envs.PINGONE_TARGET_ENVIRONMENT_ID
		p1SSOEnv = envs.PINGONE_ENVIRONMENT_ID
		accessToken = envs.PINGONE_DAVINCI_ACCESS_TOKEN
	} else {
		fmt.Println("File: ./local/env-v2-sso.json not found, \n trying env vars for PINGONE_USERNAME/PINGONE_PASSWORD")
		username = os.Getenv("PINGONE_USERNAME")
		password = os.Getenv("PINGONE_PASSWORD")
		p1SSOEnv = os.Getenv("PINGONE_ENVIRONMENT_ID")
		region = os.Getenv("PINGONE_REGION")
		companyId = os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
		accessToken = os.Getenv("PINGONE_DAVINCI_ACCESS_TOKEN")
		hostUrl = os.Getenv("PINGONE_DAVINCI_HOST_URL")
	}
	cInput := davinci.ClientInput{
		PingOneRegion:   region,
		Username:        username,
		Password:        password,
		PingOneSSOEnvId: p1SSOEnv,
		AccessToken:     accessToken,
		HostURL:         hostUrl,
	}
	client, err := davinci.NewClient(&cInput)
	fmt.Println("clientcompany: ", client.CompanyID)
	if companyId != "" {
		client.CompanyID = companyId
	}
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	return client, nil
}
