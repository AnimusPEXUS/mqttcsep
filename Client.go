package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	cfg              *ClientFlags
	target_words_len int
	target_words     []string
}

func NewClient(cfg *ClientFlags) (*Client, error) {
	self := &Client{
		cfg:          cfg,
		target_words: strings.Split(cfg.Words, " "),
	}
	self.target_words_len = len(self.target_words)
	return self, nil
}

func (self *Client) Main() {
	opts := MakeOptions(self.cfg.BrokerAddr, "goclient")

	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	for {

		in, err := rand.Int(rand.Reader, big.NewInt(int64(self.target_words_len)))
		if err != nil {
			log.Fatal("Int")
		}

		i := self.target_words[in.Int64()]

		go func(i string) {
			log.Print("sending " + i)

			if token := c.Publish(fmt.Sprintf("topic_%s", i), 0, false, i); token.Wait() && token.Error() != nil {
				log.Fatal(token.Error())
			}
		}(i)
	}

	return
}
