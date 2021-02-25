package main

import (
	"fmt"
	"os"
)

const (
	exitOK = 0
	exitFail = 1
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
	os.Exit(exitOK)
}

func run() error{
	return nil
}