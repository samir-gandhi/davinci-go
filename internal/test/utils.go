package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"

	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

type envs struct {
	PINGONE_USERNAME       string `json:"PINGONE_USERNAME"`
	PINGONE_PASSWORD       string `json:"PINGONE_PASSWORD"`
	PINGONE_ENVIRONMENT_ID string `json:"PINGONE_ENVIRONMENT_ID"`
	PINGONE_REGION         string `json:"PINGONE_REGION"`
}

func newTestClient() (*dv.APIClient, error) {
	var region, username, password, p1SSOEnv string
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
		region = envs.PINGONE_REGION
		p1SSOEnv = envs.PINGONE_ENVIRONMENT_ID
	} else {
		fmt.Println("File: ./local/env-v2-sso.json not found, \n trying env vars for PINGONE_USERNAME/PINGONE_PASSWORD")
		username = os.Getenv("PINGONE_USERNAME")
		password = os.Getenv("PINGONE_PASSWORD")
		p1SSOEnv = os.Getenv("PINGONE_ENVIRONMENT_ID")
		region = os.Getenv("PINGONE_REGION")
	}
	cInput := dv.ClientInput{
		PingOneRegion:   region,
		Username:        username,
		Password:        password,
		PingOneSSOEnvId: p1SSOEnv,
	}
	client, err := dv.NewClient(&cInput)
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	return client, nil
}

func init() {
	seed := int64(1)
	rand.New(rand.NewSource(seed))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
