package main

import (
	"flag"
)

type CommonFlags struct {
	BrokerAddr string
	Words      string
	Help       bool
}

func NewCommonFlags(name string) (*flag.FlagSet, *CommonFlags, error) {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)

	options := new(CommonFlags)

	fs.StringVar(&options.BrokerAddr, "broker", "tcp://127.0.0.1:1883", "broker address")
	fs.StringVar(&options.Words, "words", "Hello world", "words to work with")
	fs.BoolVar(&options.Help, "help", false, "show help")
	fs.BoolVar(&options.Help, "h", false, "show help")

	return fs, options, nil
}