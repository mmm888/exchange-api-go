package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/influxdata/influxdb/client/v2"
)

func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: db,
	}

	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}

	return res, nil
}

// Create a new HTTPClient
func getHTTPClient() error {
	var err error

	clnt, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: fmt.Sprintf("http://%s:%s", addr, port),
	})

	if err != nil {
		return err
	}

	return nil
}

func getEnv() {
	envAddr := os.Getenv("INFLUXDB_ADDRESS")
	if envAddr != "" {
		addr = envAddr
	}

	envPort := os.Getenv("INFLUXDB_PORT")
	if envPort != "" {
		port = envPort
	}

	envDB := os.Getenv("INFLUXDB_DB")
	if envDB != "" {
		db = envDB
	}

	envCode := os.Getenv("PAIRCODE")
	if envCode != "" {
		code = strings.Split(envCode, " ")
	}

	envLayout := os.Getenv("LAYOUT")
	if envLayout != "" {
		layout = envLayout
	}

	envStart := os.Getenv("START")
	if envStart != "" {
		start = envStart
	}

	envEnd := os.Getenv("END")
	if envEnd != "" {
		end = envEnd
	}

	envGranularity := os.Getenv("GRANULARITY")
	if envGranularity != "" {
		granularity = envGranularity
	}
}
