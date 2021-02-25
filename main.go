package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
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
	err = validateFlags(*params)
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

func validateFlags(flags Config) error{
	const serverError = "A database server name must be included. Use -server= to pass the parameter"
	const userError = "A user must be passed for connection to the database. Use -user="
	const passwordError = "No password was passed. To continue make sure to pass a password for the database connection. Use -pass="
	const dbNameError = "Name of database to be used must be included. Use -db=."
	if flags.server == "" {
		log.Print(serverError)
		return errors.New(serverError)
	}
	if flags.user == "" {
		log.Print(userError)
		return errors.New(userError)
	}
	if flags.password == "" {
		log.Print(passwordError)
		return errors.New(passwordError)
	}
	if flags.databaseName == ""{
		log.Print(dbNameError)
		return errors.New(dbNameError)
	}
	return nil
}
