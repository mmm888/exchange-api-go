package main

import (
	"fmt"
)

type Influx struct{}

var (
	initSynopsis      = "Create database of oanda for InfluxDB"
	pastSynopsis      = "Insert past data to InfluxDB (default: USD_JPY)"
	streamingSynopsis = "Insert streaming data to InfluxDB"

	influxdbUsage = fmt.Sprintf(`Usage: goanda influxdb [subcommand]

SubCommands:
    init          %s
    past          %s
    streaming     %s
`, initSynopsis, pastSynopsis, streamingSynopsis)
)

func (i *Influx) Help() string {
	return "Usage: goanda influxdb [subcommand]"
}

func (i *Influx) Run(args []string) int {
	fmt.Println(influxdbUsage)

	return 0
}

func (i *Influx) Synopsis() string {
	return "Operate InfluxDB"
}
