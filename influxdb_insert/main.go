package main

import (
	"log"

	client "github.com/influxdata/influxdb/client/v2"
	exchange "github.com/mmm888/exchange-api-go"
)

const (
	MyDB     = "exchange"
	Mytag    = "oanda"
	host     = "http://localhost:8086"
	username = ""
	password = ""
)

var (
	// oanda-api
	code        = "USD_JPY"
	start       = "2015-09-25"
	end         = "2017-09-25"
	granularity = "H3"
	layout      = "2006-01-02"
)

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

	d := new(exchange.OANDAPastData)
	d.SetData(layout, code, start, end, granularity)
	data := d.GetData()
	for _, v := range data.Candles {
		// Create a point and add to batch
		tags := map[string]string{"api": Mytag}
		fields := map[string]interface{}{
			"ask": v.OpenAsk,
			"bid": v.OpenBid,
		}

		pt, err := client.NewPoint(data.PairCode, tags, fields, v.Time)
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
