package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseFlags_Correct(t *testing.T) {
	var tests = []struct {
		args []string
		conf Config
	}{
		{
			[]string{"-server=server"},
			Config{"server", "", "", ""},
		},
		{
			 []string{"-server=server", "-u=user"},
			 Config{"server", "user", "", ""},
		},
		{
			[]string{"-server=server", "-u=user", "-pass=password"},
			Config{"server", "user", "password", ""},
		},
		{
			[]string{"-server=server", "-u=user", "-pass=password", "-db=db"},
			Config{"server", "user", "password", "db"},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			conf, output, err := parseFlags("prog", tt.args)
			if err != nil {
				t.Errorf("err got %v, want nil", err)
			}
			if output != "" {
				t.Errorf("output got %v, want empty", output)
			}
			if !reflect.DeepEqual(*conf, tt.conf) {
				t.Errorf("conf got %v, want %v", *conf, tt.conf)
			}
		})
	}
}

func TestValidateFlag_MissingFlags(t *testing.T) {
	var tests = []struct {
		name string
		config Config
		expected string
	}{
		{"no server", Config{"", "user", "pass", "db"}, "A database server name must be included. Use -server= to pass the parameter"},
		{"no user",Config{"server", "", "pass", "db"}, "A user must be passed for connection to the database. Use -user="},
		{"no password", Config{"server", "user", "", "db"}, "No password was passed. To continue make sure to pass a password for the database connection. Use -pass="},
		{"no db name", Config{"server", "user", "pass", ""},  "Name of database to be used must be included. Use -db=."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := validateFlags(tt.config)
			if actual == nil {
				t.Errorf("expected %v got nil", tt.expected)
			}
			if actual.Error() != tt.expected {
				t.Errorf("expecting error %v, got %v", tt.expected, actual.Error())
			}
		})
	}

}
