package main

import (
	"backend/broker"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var wg sync.WaitGroup

type Temperature struct {
	ID        string  `json:"id"`
	Temp      float64 `json:"temp"`
	Timestamp string  `json:"timestamp"`
}

var temp Temperature

func GetDataFromBroker() {
	opts := broker.ConnectToMqtt("broker.hivemq.com", 1883, "emqx_test_client")
	// set the message callback handler
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("%s\n", msg.Payload())
		json.Unmarshal(msg.Payload(), &temp)
	})
	opts.SetPingTimeout(1 * time.Second)
	// init new client
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	// subscribe a topic from TEMPERATURE channel
	if token := c.Subscribe("TEMPERATURE", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func HttpServer() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		json.NewEncoder(rw).Encode(temp)
	})
	http.ListenAndServe(":4000", nil)
}

func main() {
	wg.Add(2)
	go GetDataFromBroker()
	go HttpServer()
	wg.Wait()
}
