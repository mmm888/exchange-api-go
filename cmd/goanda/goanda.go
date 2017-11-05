package main

import (
	"log"
	"os"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/mitchellh/cli"
)

var (
	addr        = "db"
	port        = "8086"
	db          = "goanda"
	code        = []string{"USD_JPY"}
	layout      = "2006-01-02"
	start       = "2016-01-01"
	end         = "2020-01-31"
	granularity = "H3"

	clnt client.Client
)

func main() {
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
		"influxdb": func() (cli.Command, error) {
			return &Influx{}, nil
		},
		"influxdb init": func() (cli.Command, error) {
			return &InfluxInit{}, nil
		},
		"influxdb past": func() (cli.Command, error) {
			return &InfluxPast{}, nil
		},
		"influxdb streaming": func() (cli.Command, error) {
			return &InfluxStreaming{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
