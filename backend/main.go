package main

import (
	"backend/broker"
	"backend/http"

	"sync"
)

var wg sync.WaitGroup

func GetDataFromBroker() {
	// connect to broker
	opts := broker.ConfigMqtt("broker.hivemq.com", 1883, "emqx_test_client")
	// init new client
	broker.InitNewClient(opts)
	// subscribe a topic from TEMPERATURE channel

}

func main() {
	wg.Add(2)
	go GetDataFromBroker()
	go http.HttpServer()
	wg.Wait()
}
