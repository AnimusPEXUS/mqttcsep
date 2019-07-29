package main

import (
	"flag"
	"os"
)

type ClientFlags struct {
	*CommonFlags
}

func NewClientFlags() (*flag.FlagSet, *ClientFlags, error) {

	fs, cof, err := NewCommonFlags("client")
	if err != nil {
		return nil, nil, err
	}

	clf := &ClientFlags{cof}

	err = fs.Parse(os.Args[2:])
	if err != nil {
		return nil, nil, err
	}

	return fs, clf, nil
}
