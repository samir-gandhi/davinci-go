package main

// import (
// 	"log"
// 	"os"
// )

// // function that runs on it's own, looks in an environment for any testDataApps and deletes them.
// func sweepApps() {
// 	c, err := newTestClient()
// 	if err != nil {
// 		log.Fatalf("failed to make client %v: ", err)
// 	}
// 	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
// 	apps, err := c.ReadApplications(companyID, nil)
// 	if err != nil {
// 		log.Fatalf("failed to get apps %v: ", err)
// 	}
// 	for _, app := range apps {
// 		log.Printf("deleting app %v: ", app.Name)
// 		_, err := c.DeleteApplication(companyID, *app.AppID)
// 		if err != nil {
// 			log.Fatalf("failed to delete app %v: ", err)
// 		}
// 		break
// 	}
// }
