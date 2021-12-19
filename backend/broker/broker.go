package broker

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConfigMqtt(host string, port int, clientID string) *mqtt.ClientOptions {
	brokerHost := fmt.Sprintf("tcp://%s:%d", host, port)
	opts := mqtt.NewClientOptions().AddBroker(brokerHost).SetClientID(clientID)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("%s\n", msg.Payload())
		json.Unmarshal(msg.Payload(), &Temp)
		Temps = append(Temps, Temp)
	})

	opts.SetPingTimeout(1 * time.Second)

	return opts
}

func InitNewClient(opts *mqtt.ClientOptions) mqtt.Client {
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := c.Subscribe("TEMPERATURE", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	return c
}
