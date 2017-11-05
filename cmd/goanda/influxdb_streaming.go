package main

import (
	"log"

	client "github.com/influxdata/influxdb/client/v2"
	ex "github.com/mmm888/exchange-api-go"
)

type InfluxStreaming struct{}

func (i *InfluxStreaming) Help() string {
	return `Usage: goanda influxdb streaming

ENVIRONMENTS:
    INFLUXDB_ADDRESS     IP Adrress for InfluxDB
    INFLUXDB_PORT        Port for InfluxDB
    INFLUXDB_DB          Database name for InfluxDB
    PAIRCODE             Pair Code (ex: USD_JPY)
`
}

func (i *InfluxStreaming) Run(args []string) int {
	var err error
	getEnv()
	err = getHTTPClient()
	if err != nil {
		log.Printf("Error: %s", err)
		return 1
	}

	d := new(ex.OANDAStreamData)
	d.SetData(code)
	d.GetData()
	var data ex.StreamingData
	for {
		data = <-d.Chan

		// Create a new point batch
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  db,
			Precision: "s",
		})
		if err != nil {
			log.Printf("Error: %s", err)
			return 1
		}

		// Create a point and add to batch
		//tags:=map[string]string{"api":"oanda"}
		fields := map[string]interface{}{
			"Ask": data.Tick.Ask,
			"Bid": data.Tick.Bid,
		}

		pt, err := client.NewPoint(data.Tick.PairCode, nil, fields, data.Tick.Time)
		if err != nil {
			log.Printf("Error: %s", err)
			return 1
		}
		bp.AddPoint(pt)

		// Write the batch
		if err := clnt.Write(bp); err != nil {
			log.Printf("Error: %s", err)
			return 1
		}
	}

	return 0
}

func (i *InfluxStreaming) Synopsis() string {
	return streamingSynopsis
}
