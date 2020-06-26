package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const delayBetweenMessages = 10 * time.Second
const format = `{"data": {"time":"%s"}, "operation": "create"}`
const contentType = "application/json"

func main() {
	daprHttpPortString := os.Getenv("DAPR_HTTP_PORT")
	if daprHttpPortString == "" {
		log.Fatalf("missing environment variable %s", "DAPR_HTTP_PORT")
	}

	daprURL := fmt.Sprintf("http://localhost:%s/v1.0/bindings/eventhubs-output", daprHttpPortString)
	fmt.Printf("dapr output binding URL: %s", daprURL)
	fmt.Println("sending messages to Event Hubs")

	for {
		resp := sendMessage(daprURL)
		if err := printResponse(resp); err != nil {
			fmt.Printf("error while printing response: %s\n", err)
		}
		time.Sleep(delayBetweenMessages)
	}
}

// sendMessage sends a message to the output binding
func sendMessage(daprURL string) *http.Response {
	body := fmt.Sprintf(format, time.Now().String())
	resp, err := http.Post(daprURL, contentType, strings.NewReader(body))
	if err != nil {
		fmt.Println("failed to send request to dapr -", err)
		return nil
	}
	if resp == nil {
		fmt.Println("nil response")
		return nil
	}
	fmt.Println("sent message", body)
	return resp
}

// printResponse prints the HTTP response
func printResponse(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyString := string(bodyBytes)
	fmt.Printf("response status:%s, body: %s\n", resp.Status, bodyString)
}
