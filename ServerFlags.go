package main

import (
	"flag"
	"os"
)

type ServerFlags struct {
	*CommonFlags
}

func NewServerFlags() (*flag.FlagSet, *ServerFlags, error) {

	fs, cof, err := NewCommonFlags("server")
	if err != nil {
		return nil, nil, err
	}

	clf := &ServerFlags{cof}

	err = fs.Parse(os.Args[2:])
	if err != nil {
		return nil, nil, err
	}

	return fs, clf, nil
}
