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
