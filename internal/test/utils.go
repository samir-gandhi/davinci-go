package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

type envs struct {
	PINGONE_USERNAME              string `json:"PINGONE_USERNAME"`
	PINGONE_PASSWORD              string `json:"PINGONE_PASSWORD"`
	PINGONE_ENVIRONMENT_ID        string `json:"PINGONE_ENVIRONMENT_ID"`
	PINGONE_REGION                string `json:"PINGONE_REGION"`
	PINGONE_TARGET_ENVIRONMENT_ID string `json:"PINGONE_TARGET_ENVIRONMENT_ID"`
}

func newTestClient() (*dv.APIClient, error) {
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
	cInput := dv.ClientInput{
		PingOneRegion:   region,
		Username:        username,
		Password:        password,
		PingOneSSOEnvId: p1SSOEnv,
	}
	client, err := dv.NewClient(&cInput)
	if companyId != "" {
		client.CompanyID = companyId
	}
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	return client, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
