package main

import (
	"log"

	client "github.com/influxdata/influxdb/client/v2"
	"github.com/mmm888/exchange-api-go"
)

var (
	code = []string{
		"USD_JPY",
	}
)

const (
	MyDB     = "exchange"
	host     = "http://localhost:8086"
	username = ""
	password = ""
)

//curl -i -XPOST "http://localhost:8086/write?db=science_is_cool" --data-binary 'weather,location=us-midwest temperature=82 1465839830100400200'

func main() {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     host,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Get streaming data
	d := new(exchange.OANDAStreamData)
	d.SetData(code)
	d.GetData()
	var data exchange.StreamingData
	for {
		data = <-d.Chan

		// Create a point and add to batch
		tags := map[string]string{"api": "oanda"}
		fields := map[string]interface{}{
			"ask": data.Tick.Ask,
			"bid": data.Tick.Bid,
		}

		pt, err := client.NewPoint(data.Tick.PairCode, tags, fields, data.Tick.Time)
		if err != nil {
			log.Printf("Cannot get streaming data: %v", err)
		}
		bp.AddPoint(pt)

		// Write the batch
		err = c.Write(bp)
		if err != nil {
			log.Printf("Cannot write streaming data: %v", err)
		}
	}

}
