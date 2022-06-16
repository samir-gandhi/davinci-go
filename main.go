package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samir-gandhi/davinci-go/davinci"
)

func main() {
	fmt.Println("Hello World")
	var host *string
	username := os.Getenv("DAVINCI_USERNAME")
	password := os.Getenv("DAVINCI_PASSWORD")
	c, err := davinci.NewClient(host, &username, &password)
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	fmt.Printf("got client successfully: %s", c.HostURL)
	fmt.Println("Got Client Successfully")

	envs, err := c.GetEnvironments()
	if err != nil {
		log.Fatalf("Couldn't get envs %v: ", err)
	}
	fmt.Printf("got envs successfully: %s", envs.CustomerID)
	fmt.Println("Got All Envs Successfully")
	
	var cId *string
	env, err := c.GetEnvironment(cId)
	if err != nil {
		log.Fatalf("Couldn't get env %v: ", err)
	}
	fmt.Printf("Single Env: %s", env.CreatedByCompanyID)
	fmt.Println("Got Single Env with nil Successfully")

	env, err = c.GetEnvironment(&c.CompanyID)
	if err != nil {
		log.Fatalf("Couldn't get env %v: ", err)
	}
	fmt.Printf("Single Env: %s", env.CreatedByCompanyID)
	fmt.Println("Got Single Env with client companyID Successfully")

	envStats, err := c.GetEnvironmentStats(&c.CompanyID)
	if err != nil {
		log.Fatalf("Couldn't get env %v: ", err)
	}
	fmt.Printf("Single Env Popular Flows 0 Key: %s", envStats.PopularFlows[0].Key)
	fmt.Println("Got Env Stats with client companyID Successfully")

	msg, err := c.SetEnvironment(&c.CompanyID)
	if err != nil {
		log.Fatalf("Couldn't get env %v: ", err)
	}
	fmt.Println("Got Env Stats with client companyID Successfully")
	fmt.Printf("Single Env: %s", msg.Message)
}
