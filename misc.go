package main

import (
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MakeOptions(broker string, clientid string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientid)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	return opts
}

func GetEnvValue(name string) (string, bool) {
	for _, i := range os.Environ() {
		if strings.HasPrefix(i, name+"=") {
			spl := strings.SplitN(i, "=", 2)
			return spl[1], true
		}
	}
	return "", false
}
