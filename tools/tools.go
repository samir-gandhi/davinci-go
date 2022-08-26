package tools

import (
	// "encoding/json"
	"fmt"
	// "io"
	// "log"
	// "os"
	// "testing"
)

type envs struct {
	DAVINCI_USERNAME string `json:"DAVINCI_USERNAME"` 
	DAVINCI_PASSWORD string `json:"DAVINCI_PASSWORD"`
}

func PrintHeader(str string) {
	fmt.Printf("***START:%v***\n",str)
}

func PrintFooter(str string) {
	fmt.Printf("***END:%v***\n",str)
}



// 
// func NewTestClient() (*davinci.Client, error) {
// 	PrintHeader("Initialize Test Client")
// 	var host *string
// 	var username, password string
// 	jsonFile, err := os.Open("./local/env.json")
// 	if err == nil {
// 		defer jsonFile.Close()
// 		var envs envs
// 		byteValue, _ := io.ReadAll(jsonFile)
// 		json.Unmarshal(byteValue, &envs)
// 		username = envs.DAVINCI_USERNAME
// 		password = envs.DAVINCI_PASSWORD
// 		} else{
// 		fmt.Println("File: ./local/env.json not found, \n trying env vars for DAVINCI_USERNAME/DAVINCI_PASSWORD")
// 		username = os.Getenv("DAVINCI_USERNAME")
// 		password = os.Getenv("DAVINCI_PASSWORD")
// 	}
// 	client, err := davinci.NewClient(host, &username, &password)
// 	if err != nil {
// 		log.Fatalf("failed to make client %v: ", err)
// 	}
// 	PrintFooter("Initialize Test Client")
// 	return client, nil
// }
