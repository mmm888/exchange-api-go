package main

import (
	"fmt"
	"log"
)

type InfluxInit struct{}

func (i *InfluxInit) Help() string {
	return `Usage: goanda influxdb init

ENVIRONMENTS:
    INFLUXDB_ADDRESS     IP Adrress for InfluxDB
    INFLUXDB_PORT        Port for InfluxDB
    INFLUXDB_DB          Database name for InfluxDB
`
}

func (i *InfluxInit) Run(args []string) int {
	var err error
	getEnv()
	err = getHTTPClient()
	if err != nil {
		log.Printf("Error: %s", err)
		return 1
	}

	_, err = queryDB(clnt, fmt.Sprintf("CREATE DATABASE %s", db))
	if err != nil {
		log.Printf("Error: %s", err)
		return 1
	}

	return 0
}

func (i *InfluxInit) Synopsis() string {
	return initSynopsis
}
