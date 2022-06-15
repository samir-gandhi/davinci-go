package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samir-gandhi/davinci/davinci"
)

func main() {
	fmt.Print("Hello World")
	var host *string
	username := os.Getenv("DAVINCI_USERNAME")
	password := os.Getenv("DAVINCI_PASSWORD")
	c, err := davinci.NewClient(host, &username, &password)
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	fmt.Printf("got client successfully: %v", *c)

	envs, err := c.GetEnvironments()
	if err != nil {
		log.Fatalf("Couldn't get envs %v: ", err)
	}
	fmt.Printf("got envs successfully: %v", envs)
}
