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
}

func NewServer(cfg *ServerFlags) (*Server, error) {
	self := &Server{
		cfg:          cfg,
		start_time:   time.Now(),
		target_words: strings.Split(cfg.Words, " "),
	}
	self.target_words_len = len(self.target_words)
	return self, nil
}

func (self *Server) Main() {

	opts := MakeOptions(self.cfg.BrokerAddr)
	opts.SetDefaultPublishHandler(self.handler)

	s := mqtt.NewClient(opts)

	for _, i := range self.target_words {
		go func() {
			topic := fmt.Sprintf("topic_%s", i)
			log.Print("subscribing " + topic)
			s.Subscribe(topic, 0, nil)
		}()
	}

	return
}

func (self *Server) handler(c mqtt.Client, m mqtt.Message) {
	mut := &sync.Mutex{}
	mut.Lock()
	defer mut.Unlock()

	ps := string(m.Payload())

	self.active_words = append(self.active_words, ps)

	len_self_active_words := len(self.active_words)

	if self.active_words[len_self_active_words-1] != self.target_words[len_self_active_words] {
		self.active_words = make([]string, 0)
		return
	}

	if len_self_active_words == self.target_words_len {
		time_spent := time.Now().Sub(self.start_time)
		log.Printf("%s: time waited: %s", strings.Join(self.active_words, " "), time_spent.String())
		self.start_time = time.Now()
	}

}
