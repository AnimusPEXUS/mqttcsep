package main

import (
	"log"
	"os"
)

const SOCM = "required parameter 'server' or 'client'"

func main() {
	mode := "none"
	if len(os.Args) < 2 {
		log.Fatal(SOCM)
	}
	mode = os.Args[1]

	if mode != "server" && mode != "client" {
		log.Fatal(SOCM)
	}

	if mode == "server" {
		log.Print("server")
		fs, sf, err := NewServerFlags()
		if err != nil {
			log.Fatal(err)
		}
		if sf.Help {
			fs.PrintDefaults()
			log.Fatal("showing help")
		}
		log.Println("words", sf.Words)
		c, err := NewServer(sf)
		if err != nil {
			log.Fatal(err)
		}
		c.Main()
	} else {
		log.Print("client")
		fs, cf, err := NewClientFlags()
		if err != nil {
			log.Fatal(err)
		}
		if cf.Help {
			fs.PrintDefaults()
			log.Fatal("showing help")
		}
		log.Println("words", cf.Words)
		s, err := NewClient(cf)
		if err != nil {
			log.Fatal(err)
		}
		s.Main()
	}
}
