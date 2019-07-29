package main

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MakeOptions(broker string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("goslientserver")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	return opts
}
