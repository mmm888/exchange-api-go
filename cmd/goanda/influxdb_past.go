package main

import (
	"log"

	client "github.com/influxdata/influxdb/client/v2"
	ex "github.com/mmm888/exchange-api-go"
)

type InfluxPast struct{}

func (i *InfluxPast) Help() string {
	return `Usage: goanda influxdb past

ENVIRONMENTS:
    INFLUXDB_ADDRESS     IP Adrress for InfluxDB
    INFLUXDB_PORT        Port for InfluxDB
    INFLUXDB_DB          Database name for InfluxDB
    PAIRCODE             Pair Code (ex: USD_JPY)
    LAYOUT               Time format for Golang (ex: 2006-01-02)
    START                Start date
    END                  End date
    GRANULARITY          Granularity to take value (ex: 3H)
`
}

func (i *InfluxPast) Run(args []string) int {
	getEnv()
	getHTTPClient()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})
	if err != nil {
		log.Printf("Error: %s", err)
		return 1
	}

	for _, pairCode := range code {

		d := new(ex.OANDAPastData)
		d.SetData(layout, pairCode, start, end, granularity)
		data, err := d.GetData()
		if err != nil {
			log.Printf("Error: %s", err)
			return 1
		}

		checkNil := &ex.PastData{}

		for _, v := range data.Candles {
			if v == checkNil.Candles[0] {
				break
			}

			// Create a point and add to batch
			//tags:=map[string]string{"api":"oanda"}
			fields := map[string]interface{}{
				"Ask": v.OpenAsk,
				"Bid": v.OpenBid,
			}

			pt, err := client.NewPoint(pairCode, nil, fields, v.Time)
			if err != nil {
				log.Printf("Error: %s", err)
				return 1
			}

			bp.AddPoint(pt)
		}
	}

	// Write the batch
	if err := clnt.Write(bp); err != nil {
		log.Printf("Error: %s", err)
		return 1
	}

	return 0
}

func (i *InfluxPast) Synopsis() string {
	return pastSynopsis
}
