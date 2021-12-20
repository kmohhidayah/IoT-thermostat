package main

import (
	"backend/broker"
	"backend/http"
	"fmt"
	"os"

	"sync"

	_ "github.com/joho/godotenv/autoload"
)

var (
	HOST      string = os.Getenv("HOST")
	PORT      string = os.Getenv("PORT")
	CLIENT_ID string = os.Getenv("CLIENT_ID")
)

func GetDataFromBroker() {
	config := fmt.Sprintf("tcp://%s:%s", HOST, PORT)
	// connect to broker
	opts := broker.ConfigMqtt(config, CLIENT_ID)
	// init new client
	broker.InitNewClient(opts)
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go GetDataFromBroker()
	go http.HttpServer()
	wg.Wait()
}
