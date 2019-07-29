package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Server struct {
	cfg              *ServerFlags
	target_words     []string
	target_words_len int
	active_words     []string
	start_time       time.Time
	mut              *sync.Mutex
}

func NewServer(cfg *ServerFlags) (*Server, error) {
	self := &Server{
		cfg:          cfg,
		start_time:   time.Now(),
		target_words: strings.Split(cfg.Words, " "),
	}
	self.target_words_len = len(self.target_words)
	self.mut = &sync.Mutex{}
	return self, nil
}

func (self *Server) Main() {

	opts := MakeOptions(self.cfg.BrokerAddr, "goserver")
	opts.SetDefaultPublishHandler(self.handler)

	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	for _, i := range self.target_words {
		{
			i := i
			go func() {
				topic := fmt.Sprintf("topic_%s", i)
				log.Print("subscribing " + topic)
				if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
					log.Fatal(token.Error())
				}
			}()
		}
	}

	return
}

func (self *Server) handler(c mqtt.Client, m mqtt.Message) {

	self.mut.Lock()
	defer self.mut.Unlock()

	ps := string(m.Payload())

	self.active_words = append(self.active_words, ps)

	len_self_active_words := len(self.active_words)

	if self.active_words[len_self_active_words-1] != self.target_words[len_self_active_words-1] {
		self.active_words = make([]string, 0)
		return
	}

	if len_self_active_words == self.target_words_len {
		time_spent := time.Now().Sub(self.start_time)
		log.Printf("%s: time waited: %s", strings.Join(self.active_words, " "), time_spent.String())
		self.active_words = make([]string, 0)
		self.start_time = time.Now()
	}

}
