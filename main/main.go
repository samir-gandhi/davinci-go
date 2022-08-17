package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"math/rand"
	"time"

	"github.com/samir-gandhi/davinci-client-go"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func printheader(str string) {
	fmt.Println("***START:"+str+"***")
}
func printfooter(str string) {
	fmt.Println("***END:"+str+"***")
}

func main() {
	printheader("Initialize")
	var host *string
	username := os.Getenv("DAVINCI_USERNAME")
	password := os.Getenv("DAVINCI_PASSWORD")
	c, err := davinci.NewClient(host, &username, &password)
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	fmt.Printf("got client successfully: %s\n", c.HostURL)
	printfooter("Got Client Successfully")
	
	printheader("Get Environments")
	envs, err := c.GetEnvironments()
	if err != nil {
		log.Fatalf("Couldn't get envs %v: ", err)
	}
	fmt.Printf("got envs successfully: %s\n", envs.Companies[0])
	printfooter("Got All Envs Successfully")


	var comp string
	// Set Company Id to preferred env
	for i, v := range envs.Companies {
		// if v.Name == "tempdvflows" {
		if v.Name == "tempdvflows" {
			// fmt.Printf("company id is: %s", envs.Companies[i].CompanyID)
			c := &comp
			*c = envs.Companies[i].CompanyID
		}
	}
	fmt.Printf("Setting Company Id for this run: %s", comp)

	// Sample, easy gets
	var cId *string
	env, err := c.GetEnvironment(cId)
	if err != nil {
		log.Fatalf("Couldn't get env %v: ", err)
	}
	fmt.Printf("Single Env: %s\n", env.CreatedByCompanyID)
	printfooter("Got Single Env with nil Successfully")

	env, err = c.GetEnvironment(&comp)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	fmt.Printf("Single Env: %s\n", env.CreatedByCompanyID)
	printfooter("Got Single Env with client companyID Successfully")

	envStats, err := c.GetEnvironmentStats(&comp)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	fmt.Printf("Single Env Number of Flows: %d\n", envStats.TableStats[0].Flows)
	printfooter("Got Env Stats with defined companyID Successfully")

	msg, err := c.SetEnvironment(&comp)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	fmt.Printf("Single Env: %s\n", msg.Message)
	printfooter("Set Env Successfully")

	// Create Customer
	printheader("Create Customer")
	rand.Seed(time.Now().UnixNano())
	ccEmail := "samirgandhi+" + randSeq(10) + "tf@pingidentity.com"
	ccPayload := &davinci.CustomerCreate{
		Email:       ccEmail,
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

	//Get All Customers
	team, err := c.GetCustomers(&comp, nil)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	if team != nil {
		fmt.Printf("First Customer in Env: %s\n", team.Customers[0].FirstName)
		printfooter("Got Customers Successfully")
	}

	// Find the new customer
	var cust string
	total := float64(team.CustomerCount)
	pagesize := float64(10)
	pages := int(math.Ceil(total / pagesize)-1)
	for i := 0; i <= pages; i++ {
		args := &davinci.Params{
			Limit: strconv.Itoa(int(math.RoundToEven(pagesize))),
			Page: strconv.Itoa(i),
		}
		team, err := c.GetCustomers(&comp, args)
		if err != nil {
			log.Fatalf("Couldn't get %v: ", err)
		}
		if team != nil {
			fmt.Println("Got Customers round: ", i)
		}
		for i, v := range team.Customers {
			fmt.Printf("Checking if: %v\n", v.Email)
			// fmt.Printf("is equal to: %v\n", "samirgandhi+ckdxsnbvuatf@pingidentity.com")
			// if v.Email == "samirgandhi+tmphzpgcssbutf@pingidentity.com" {
			fmt.Printf("is equal to: %v\n", ccPayload.Email)
			if v.Email == ccPayload.Email {
				fmt.Printf("company id is: %s\n", team.Customers[i].CustomerID)
				cu := &cust
				*cu = team.Customers[i].CustomerID
				break
			}
		}
		if cust != "" {
			break
		}
	}
	if cust == "" {
		log.Fatal("customer not found!")
	}
	fmt.Printf("using customer Id: %s", cust)

	// TODO: cannot use ccMsg.CustomerID because it is returning a 400
	// for i, v := range envs.Companies {
	//   if v.Name == "tempdvflows" {
	//       // fmt.Printf("company id is: %s", envs.Companies[i].CompanyID)
	// 			c := &comp
	// 			*c = envs.Companies[i].CompanyID
	//   }
	// }

	customer, err := c.GetCustomer(&comp, &cust)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	if customer != nil {
		// fmt.Printf("This Customer First Name: %s\n", customer.FirstName)
		fmt.Print(customer)
		printfooter("Got New Customer Successfully")
	}

	//Update Customer
	cuPayload := &davinci.CustomerUpdate{
		FirstName:   "samir",
		LastName:    "Gandhi",
		Roles:       []string{"default:read"},
		PhoneNumber: "1234",
	}
	cuMsg, err := c.UpdateCustomer(&comp, &cust, cuPayload)
	if err != nil {
		log.Fatalf("Couldn't update %v: ", err)
	}
	if cuMsg != nil {
		fmt.Printf("Updated customer response: %s\n", cuMsg.Message)
		printfooter("Updated one customer Successfully")
	}

	cdMsg, err := c.DeleteCustomer(&comp, &cust)
	if err != nil {
		// fmt.Printf("Failed to create customer response: %v\n", err)
		log.Fatalf("Couldn't get %v: ", err)
	}
	if cdMsg != nil {
		// fmt.Printf("This Customer First Name: %s\n", customer.FirstName)
		fmt.Print(cdMsg.Message)
		fmt.Println("Deleted Customer Successfully")
	}
	printfooter("Customers")

	//Get All roles
	printheader("Roles")
	roles, err := c.GetRoles(&comp, nil)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	if roles != nil {
		fmt.Printf("Roles is: %+v\n", roles)
		printfooter("Got Roles Successfully")
	}

	seq := "tftest"+randSeq(4)
	crPayload := davinci.RoleCreate{Name: seq}
	// rand.Seed(time.Now().UnixNano())
	// ccEmail := "samirgandhi+" + randSeq(10) + "tf@pingidentity.com"
	// ccPayload := &davinci.CustomerCreate{
	// 	Email:       ccEmail,
	// 	FirstName:   "samir",
	// 	LastName:    "Gandhi",
	// 	Roles:       []string{"default:admin", "default:read"},
	// 	PhoneNumber: "1234",
	// }
	crMsg, err := c.CreateRole(&comp, &crPayload)
	if err != nil {
		fmt.Printf("Failed to create role: %v\n", err)
	}
	if crMsg != nil {
		fmt.Printf("Created Role response: %s\n", crMsg.ID.Name)
		fmt.Println("Created Role Successfully")
	}

	//Update Role 
	jsonString := "{\"description\":\"roledesc\",\"policy\":[{\"resource\":\"company\",\"actions\":[{\"action\":\"read\",\"allow\":true},{\"action\":\"update\",\"allow\":true},{\"action\":\"create\",\"allow\":true},{\"action\":\"delete\",\"allow\":true}]}]}"
	bytes := []byte(jsonString)

	var role davinci.RoleUpdate
	err = json.Unmarshal(bytes, &role)
	ruMsg, err := c.UpdateRole(&comp, &crMsg.ID.Name, &role)
	if err != nil {
		log.Fatalf("Couldn't update %v: ", err)
	}
	if cuMsg != nil {
		fmt.Printf("Updated role: %s\n", ruMsg.Description)
		printfooter("Updated role successfully")
	}

	drMsg, err := c.DeleteRole(&comp, &crMsg.ID.Name)
	if err != nil {
		// fmt.Printf("Failed to create customer response: %v\n", err)
		log.Fatalf("Couldn't get %v: ", err)
	}
	if cdMsg != nil {
		// fmt.Printf("This Customer First Name: %s\n", customer.FirstName)
		fmt.Print(drMsg.Message)
		fmt.Println("Deleted Role Successfully")
	}
	
	printfooter("Roles")

	printheader("Connections")
	connections, err := c.GetConnections(&comp, nil)
	if err != nil {
		log.Fatalf("Couldn't get %v: ", err)
	}
	if roles != nil {
		fmt.Printf("Connections is: %+v\n", connections)
		printfooter("Got Connections Successfully")
	}


}

