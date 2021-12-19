package broker

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConnectToMqtt(host string, port int, clientID string) *mqtt.ClientOptions {
	brokerHost := fmt.Sprintf("tcp://%s:%d", host, port)
	opts := mqtt.NewClientOptions().AddBroker(brokerHost).SetClientID(clientID)
	opts.SetKeepAlive(60 * time.Second)

	return opts
}
