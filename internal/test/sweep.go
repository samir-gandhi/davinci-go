package test

import (
	"log"
	"os"
	"strings"
)

//Run this file with go run sweep.go to delete all test data from an environment.

// function that runs on it's own, looks in an environment for any testDataApps and deletes them.
func sweepApps() {
	c, err := newTestClient()
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	apps, err := c.ReadApplications(companyID, nil)
	if err != nil {
		log.Fatalf("failed to get apps %v: ", err)
	}
	for _, app := range apps {
		for _, tdApp := range testDataApps {
			if strings.Contains(app.Name, tdApp.Name) {
				log.Printf("deleting app %v: ", app.Name)
				_, err := c.DeleteApplication(companyID, *app.AppID)
				if err != nil {
					log.Fatalf("failed to delete app %v: ", err)
				}
				break
			}
		}
	}
}
