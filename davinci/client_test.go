package davinci

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"github.com/samir-gandhi/davinci-client-go/tools"
)

type envs struct {
	DAVINCI_USERNAME string `json:"DAVINCI_USERNAME"` 
	DAVINCI_PASSWORD string `json:"DAVINCI_PASSWORD"`
	DAVINCI_COMPANYID string `json:"DAVINCI_COMPANYID"`
}


func TestNewClient(t *testing.T) {
	var host *string
	var username, password string
	jsonFile, err := os.Open("../local/env.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		username = envs.DAVINCI_USERNAME
		password = envs.DAVINCI_PASSWORD
		} else{
		fmt.Println("File: ./local/env.json not found, \n trying env vars for DAVINCI_USERNAME/DAVINCI_PASSWORD")
		username = os.Getenv("DAVINCI_USERNAME")
		password = os.Getenv("DAVINCI_PASSWORD")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	client, err := NewClient(host, &username, &password)
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	fmt.Printf("\ngot client successfully, with companyId: %v\n", client.CompanyID)
}


func newTestClient() (*Client, error) {
	tools.PrintHeader("Initializing Test Client")
	defer tools.PrintFooter("Initializing Test Client")
	var host *string
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
		} else{
		fmt.Println("File: ./local/env.json not found, \n trying env vars for DAVINCI_USERNAME/DAVINCI_PASSWORD")
		username = os.Getenv("DAVINCI_USERNAME")
		password = os.Getenv("DAVINCI_PASSWORD")
		companyid = os.Getenv("DAVINCI_COMPANYID")
	}
	client, err := NewClient(host, &username, &password)
	if companyid != "" {
		client.CompanyID = companyid
	}
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	return client, nil
}