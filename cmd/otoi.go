package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/cli"
)

type Config struct {
	DB          string
	Tag         string
	Host        string
	Layout      string
	Granularity string
}

var (
	config Config
)

func main() {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Printf("Cannot get config: %v", err)
	}

	c := cli.NewCLI("otoi", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"stream": func() (cli.Command, error) {
			return &Stream{}, nil
		},
		"list": func() (cli.Command, error) {
			return &List{}, nil
		},
		"get": func() (cli.Command, error) {
			return &Get{}, nil
		},
		"import": func() (cli.Command, error) {
			return &Import{}, nil
		},
		"insert": func() (cli.Command, error) {
			return &Insert{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
