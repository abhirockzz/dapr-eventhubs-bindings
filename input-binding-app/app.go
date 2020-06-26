package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := getAddress()

	fmt.Printf("will listen on %s\n", addr)
	http.HandleFunc("/eventhubs-input", func(rw http.ResponseWriter, req *http.Request) {
		printRequest(req)
		if subscribe(rw, req) {
			return
		}
		writeResponse(rw)
	})
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

// getAddress retturns the address where the http server will bind to
func getAddress() string {
	// parse application port
	fmt.Println("init")
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatalf("missing environment variable %s", "APP_PORT")
	}
	addr := ":" + port
	return addr
}

// printRequest prints the HTTP request
func printRequest(req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Printf("received request, method: %s, payload: %s\n", req.Method, bodyString)
}

// signal interest in the input binding
func subscribe(rw http.ResponseWriter, req *http.Request) bool {
	if req.Method == "OPTIONS" {
		fmt.Printf("subscribing to binding")
		rw.WriteHeader(200)
		return true
	}
	return false
}

// be nice and write a empty response bck
func writeResponse(rw http.ResponseWriter) {
	rw.WriteHeader(200)
	if _, err := rw.Write([]byte(`{}`)); err != nil {
		fmt.Printf("error while sending response: %v\n", err)
	}
}
