package atccustomlist

import "fmt"
import "os"
import "gopkg.in/resty.v1"

func CustomListCreate(name string, desc string, debugflag bool){
	if os.Getenv("ATCKEY") == "" {
		fmt.Printf("Environment variable ATCKEY must be set to your ATC API key.\n")
		os.Exit(1)
	}
	apikey := "Token " + os.Getenv("ATCKEY")
	body := "{\"name\":\"" + name + "\", \"description\":\"" + desc + "\"}"
	client1 := resty.New()
	if debugflag{
		client1.SetDebug(true)
	}
	resp, err := client1.R().
		SetHeaders(map[string]string{"Content-Type": "application/json", "Authorization": apikey}).
		SetBody(body).
		Post("https://www-test.csp.infoblox.com/api/atcfw/v1/custom_list")
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v\n", resp)
}
