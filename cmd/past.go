package main

import (
	"fmt"
	"log"

	ex "github.com/mmm888/exchange-api-go"
)

type Past struct{}

func (l *Past) Help() string {
	return "Usage: goanda past"
}

func (l *Past) Run(args []string) int {
	layout := "2006-01-02"
	pairCode := "USD_JPY"
	start := "2016-01-01"
	end := "2016-01-31"
	granularity := "D"

	d := new(ex.OANDAPastData)
	d.SetData(layout, pairCode, start, end, granularity)
	data, err := d.GetData()
	if err != nil {
		log.Print(err)
		return 1
	}

	checkNil := &ex.PastData{}
	for _, v := range data.Candles {
		if v == checkNil.Candles[0] {
			break
		}

		fmt.Printf("Time: %s, Ask: %10f, Bid: %10f\n",
			v.Time.Format(layout), v.OpenAsk, v.OpenBid)
	}

	return 0
}

func (l *Past) Synopsis() string {
	return "Get ask, bid from start to end (default: USD_JPY)"
}
