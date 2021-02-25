package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

const (
	exitOK = 0
	exitFail = 1
)
type Config struct{
	server string
	user string
	password string
	databaseName string
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
	os.Exit(exitOK)
}

func run() error{
	params, output, err :=  parseFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		return err
	} else if err != nil {
		fmt.Println("encountered error:", err)
		fmt.Println("output:\n ", output)
	}
	fmt.Println(params)
	return nil
}

//parse flags and return map
func parseFlags(progName string, args []string) (*Config, string, error ){
	flags := flag.NewFlagSet(progName, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)
	fmt.Println(args, progName)
	var conf Config
	flags.StringVar(&conf.server, "server", "", "name of server where database is located")
	flags.StringVar( &conf.user,"u", "", "user id to sign in to database server")
	flags.StringVar(&conf.password,"pass", "", "password for sign in to db server")
	flags.StringVar(&conf.databaseName,"db", "", "name of database to be queried")

	err := flags.Parse(args)
	fmt.Println(conf)
	if err != nil {
		return nil, buf.String(), err
	}

	return &conf, buf.String(), nil
}
