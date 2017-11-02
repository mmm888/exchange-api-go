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

	c := cli.NewCLI("goanda", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &List{}, nil
		},
		"current": func() (cli.Command, error) {
			return &Current{}, nil
		},
		"past": func() (cli.Command, error) {
			return &Past{}, nil
		},
		"streaming": func() (cli.Command, error) {
			return &Streaming{}, nil
		},
		"influxdb init": func() (cli.Command, error) {
			return &InfluxInit{}, nil
		},
		/*
			"influxdb insert": func() (cli.Command, error) {
				return &InfluxInsert{}, nil
			},
			"influxdb streaming": func() (cli.Command, error) {
				return &InfluxStreaming{}, nil
			},
		*/
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
