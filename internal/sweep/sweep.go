package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"

	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func main() {

	if os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID") == "" {
		log.Fatalf("PINGONE_TARGET_ENVIRONMENT_ID not set")
	}

	sweepApps()
	sweepFlows()
	sweepConnectors()
	sweepVariables()
}

func sweepApps() {
	c, err := newTestClient()
	if err != nil {
		log.Fatalf("failed to make client: %v", err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	apps, err := c.ReadApplications(companyID, nil)
	if err != nil {
		log.Fatalf("failed to get apps: %v", err)
	}
	for _, app := range apps {
		log.Printf("deleting app: %v", app.Name)
		_, err := c.DeleteApplication(companyID, *app.AppID)
		if err != nil {
			log.Fatalf("failed to delete app: %v", err)
			continue
		}
	}
}

func sweepFlows() {
	c, err := newTestClient()
	if err != nil {
		log.Fatalf("failed to make client: %v", err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	flows, err := c.ReadFlows(companyID, nil)
	if err != nil {
		log.Fatalf("failed to get flows: %v", err)
	}
	for _, flow := range flows {
		log.Printf("deleting flow: %v", flow.Name)
		_, err := c.DeleteFlow(companyID, flow.FlowID)
		if err != nil {
			log.Fatalf("failed to delete flow: %v", err)
			continue
		}
	}
}

func sweepConnectors() {
	c, err := newTestClient()
	if err != nil {
		log.Fatalf("failed to make client: %v", err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	connections, err := c.ReadConnections(companyID, nil)
	if err != nil {
		log.Fatalf("failed to get connections: %v", err)
	}
	for _, connection := range connections {
		log.Printf("deleting connection: %v", *connection.Name)
		_, err := c.DeleteConnection(companyID, *connection.ConnectionID)
		if err != nil {
			log.Fatalf("failed to delete connection: %v", err)
			continue
		}
	}
}

func sweepVariables() {
	c, err := newTestClient()
	if err != nil {
		log.Fatalf("failed to make client: %v", err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	variables, err := c.ReadVariables(companyID, nil)
	if err != nil {
		log.Fatalf("failed to get connections: %v", err)
	}
	for key, variable := range variables {
		log.Printf("deleting variable: %v", variable.DisplayName)
		_, err := c.DeleteVariable(companyID, key)
		if err != nil {
			log.Fatalf("failed to delete variable: %v", err)
			continue
		}
	}
}

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
			log.Fatalf("failed to unmarshal json: %v", errJ)
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
		log.Fatalf("failed to make client: %v", err)
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
