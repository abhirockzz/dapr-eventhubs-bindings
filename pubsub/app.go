package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var port string

func init() {
	port = os.Getenv("APP_PORT")
	if port == "" {
		log.Fatalf("missing environment variable %s", "APP_PORT")
	}
}

type subscription struct {
	Topic string `json:"topic"`
	Route string `json:"route"`
}

func newTopicHandler(topic string, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/%s", topic), func(writer http.ResponseWriter, request *http.Request) {
		topic := topic
		defer request.Body.Close()

		bb, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Error(err)
			return
		}
		log.WithFields(log.Fields{"topic": topic}).Info(string(bb))
		if _, err := writer.Write(bb); err != nil {
			log.Printf("error while sending message: %v", err)
		}
	})

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/dapr/subscribe", func(writer http.ResponseWriter, request *http.Request) {
		subscriptions := []subscription{
			// {
			// 	Topic: "A",
			// 	Route: "A",
			// },
			// {
			// 	Topic: "B",
			// 	Route: "B",
			// },
			// {
			// 	Topic: "C",
			// 	Route: "C",
			// },
			//{
			//	Topic: "D",
			//	Route: "D",
			//},
			{
				Topic: "E",
				Route: "E",
			},
			{
				Topic: "F",
				Route: "F",
			},
		}
		subs, err := json.Marshal(subscriptions)
		if err != nil {
			log.Error(err)
			return
		}
		log.Info(string(subs))
		if _, err := writer.Write(subs); err != nil {
			log.Printf("error while sending message: %v", err)
		}
	})

	// newTopicHandler("A", r)
	// newTopicHandler("B", r)
	// newTopicHandler("C", r)
	// newTopicHandler("D", r)
	newTopicHandler("E", r)
	newTopicHandler("F", r)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		panic(err)
	}
}
