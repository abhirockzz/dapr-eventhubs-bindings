package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const daprURL = "http://localhost:3500/v1.0/bindings/eventhubs-output"

func main() {
	format := `{"data": {"time":"%s"}}`
	contentType := "application/json"
	done := make(chan bool)
	go func() {
		fmt.Println("sending messages to Event Hubs")

		for i := 1; i <= 5; i++ {
			body := fmt.Sprintf(format, time.Now().String())
			resp, err := http.Post(daprURL, contentType, strings.NewReader(body))
			if err != nil {
				fmt.Println("failed to send request to dapr -", err)
				continue
			}
			if resp == nil {
				fmt.Println("nil response")
				continue
			}
			fmt.Println("sent message", body)
			fmt.Println("response", resp.Status)
			time.Sleep(2 * time.Second)
		}
		done <- true
	}()
	<-done
	fmt.Println("finished sending messages... use ctrl+c to exit")
}
