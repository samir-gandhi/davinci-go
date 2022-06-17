package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samir-gandhi/davinci-go/davinci"
	"time"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fmt.Println("Hello World")
	var host *string
	username := os.Getenv("DAVINCI_USERNAME")
	password := os.Getenv("DAVINCI_PASSWORD")
	c, err := davinci.NewClient(host, &username, &password)
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	fmt.Printf("got client successfully: %s\n", c.HostURL)
	fmt.Println("Got Client Successfully")

	envs, err := c.GetEnvironments()
	if err != nil {
		log.Fatalf("Couldn't get envs %v: ", err)
	}
	fmt.Printf("got envs successfully: %s\n", envs.Companies[0])
	fmt.Println("Got All Envs Successfully")
	var comp string

	for i, v := range envs.Companies {
    if v.Name == "tempdvflows" {
        // fmt.Printf("company id is: %s", envs.Companies[i].CompanyID)
				c := &comp
				*c = envs.Companies[i].CompanyID
    }
	}

	fmt.Printf("using company Id: %s", comp)

	var cId *string
	env, err := c.GetEnvironment(cId)
	if err != nil {
		log.Fatalf("Couldn't get env %v: ", err)
	}
	fmt.Printf("Single Env: %s\n", env.CreatedByCompanyID)
	fmt.Println("Got Single Env with nil Successfully")


	env, err = c.GetEnvironment(&comp)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	fmt.Printf("Single Env: %s\n", env.CreatedByCompanyID)
	fmt.Println("Got Single Env with client companyID Successfully")

	envStats, err := c.GetEnvironmentStats(&comp)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	fmt.Printf("Single Env Popular Flows 0 Key: %s\n", envStats.PopularFlows[0].Key)
	fmt.Println("Got Env Stats with client companyID Successfully")

	msg, err := c.SetEnvironment(&comp)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	fmt.Printf("Single Env: %s\n", msg.Message)
	fmt.Println("Got Env Stats with client companyID Successfully")

	rand.Seed(time.Now().UnixNano())
	ccEmail := "samirgandhi+"+randSeq(10)+"tf@pingidentity.com"
	ccPayload := &davinci.CustomerCreate{
		Email: ccEmail,
		FirstName:   "samir",
		LastName:    "Gandhi",
		Roles:       []string{"default:admin", "default:read"},
		PhoneNumber: "1234",
	}
	ccMsg, err := c.CreateCustomer(&comp, ccPayload)
	if err != nil {
		fmt.Printf("Failed to create customer response: %v\n", err)
	}
	if ccMsg != nil {
		fmt.Printf("Created customer response: %s\n", ccMsg.CustomerID)
		fmt.Println("Created One Customer Successfully")
	}

	team, err := c.GetCustomers(&comp)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	if team != nil {
		fmt.Printf("First Customer in Env: %s\n", team.Customers[0].FirstName)
		fmt.Println("Got Customers Successfully")
	}

	// TODO: cannot use ccMsg.CustomerID because it is returning a 400
	var cust string
	// for i, v := range envs.Companies {
  //   if v.Name == "tempdvflows" {
  //       // fmt.Printf("company id is: %s", envs.Companies[i].CompanyID)
	// 			c := &comp
	// 			*c = envs.Companies[i].CompanyID
  //   }
	// }
	for i, v := range team.Customers {
		fmt.Printf("Checking if: %v\n",v.Email)
		fmt.Printf("is equal to: %v\n",ccPayload.Email)
    if v.Email == ccPayload.Email {
        // fmt.Printf("company id is: %s", envs.Companies[i].CompanyID)
				cu := &cust
				*cu = team.Customers[i].CustomerID
			}
		}
	if &cust == nil {
		log.Fatalf("Customer not found %v: ", err)
	}

	fmt.Printf("using customer Id: %s", cust)

	// customer, err := c.GetCustomer(&comp, &cust)
	// if err != nil {
	// 	log.Fatalf("Couldn't get %v: ", err)
	// }
	// if customer != nil {
	// 	// fmt.Printf("This Customer First Name: %s\n", customer.FirstName)
	// 	fmt.Print(customer)
	// 	fmt.Println("Got One Customer Successfully")
	// }

	// // cPayload := davinci.CustomerUpdate{"samir","gandhi",[]string{"hello"},"1234"}
	// cuPayload := &davinci.CustomerUpdate{
	// 	FirstName:   "samir",
	// 	LastName:    "Gandhi",
	// 	Roles:       []string{"default:admin", "default:read"},
	// 	PhoneNumber: "1234",
	// }
	// cuMsg, err := c.UpdateCustomer(&comp, &team.Customers[0].CustomerID, cuPayload)
	// if err != nil {
	// 	log.Fatalf("Couldn't update %v: ", err)
	// }
	// if cuMsg != nil {
	// 	fmt.Printf("Updated customer response: %s\n", cuMsg.Message)
	// 	fmt.Println("Got One Customer Successfully")
	// }

	// cdMsg, err := c.DeleteCustomer(&comp, &cust)
	// if err != nil {
	// 	// fmt.Printf("Failed to create customer response: %v\n", err)
	// 	log.Fatalf("Couldn't get %v: ", err)
	// }
	// if cdMsg != nil {
	// 	// fmt.Printf("This Customer First Name: %s\n", customer.FirstName)
	// 	fmt.Print(cdMsg)
	// 	fmt.Println("Deleted Customer Successfully")
	// }

}
