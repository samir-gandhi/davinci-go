package main

import (
	"fmt"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

func main() {
	c, err := davinci.NewClient(nil)
	if err == nil {
		fmt.Println("Using region:", c.PingOneRegion)
	}
}
