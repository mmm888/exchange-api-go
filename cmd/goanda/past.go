package main

import (
	"fmt"
	"log"

	"flag"

	ex "github.com/mmm888/exchange-api-go"
)

type Past struct{}

func (l *Past) Help() string {
	return `Usage: goanda past [option]

Options:
  -code string
        Pair Code (default "USD_JPY")
  -end string
        End date (default "2020-01-01")
  -granularity string
        Granularity to take value (default "D")
  -layout string
        Time format for Golang (default "2006-01-02")
  -start string
        Start date (default "2017-01-01")
`
}

func (l *Past) Run(args []string) int {
	var code, layout, start, end, granularity string
	f := flag.NewFlagSet("", flag.ContinueOnError)
	f.StringVar(&code, "code", "USD_JPY", "Pair Code")
	f.StringVar(&layout, "layout", "2006-01-02", "Time format for Golang")
	f.StringVar(&start, "start", "2017-01-01", "Start date")
	f.StringVar(&end, "end", "2020-01-01", "End date")
	f.StringVar(&granularity, "granularity", "D", "Granularity to take value")
	if err := f.Parse(args); err != nil {
		log.Print(err)
		return 1
	}

	d := new(ex.OANDAPastData)
	d.SetData(layout, code, start, end, granularity)
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

		fmt.Printf("Time: %s, Ask: %10f, Bid: %10f\n",
			v.Time.Format(layout), v.OpenAsk, v.OpenBid)
	}

	return 0
}

func (l *Past) Synopsis() string {
	return "Get ask, bid from start to end (default: USD_JPY)"
}
