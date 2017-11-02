package main

import (
	"fmt"

	"time"

	ex "github.com/mmm888/exchange-api-go"
)

type Streaming struct{}

func (l *Streaming) Help() string {
	return "Usage: goanda streaming"
}

func (l *Streaming) Run(args []string) int {
	instruments := []string{
		"USD_JPY",
	}
	format := "2006-01-02 15:04:05"

	d := new(ex.OANDAStreamData)
	d.SetData(instruments)
	d.GetData()

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	for {
		data := <-d.Chan
		time := data.Tick.Time.In(jst)
		fmt.Printf("Time: %s, Code: %v, Ask: %10f, Bid: %10f\n",
			time.Format(format), data.Tick.PairCode, data.Tick.Ask, data.Tick.Bid)
	}

	return 0
}

func (l *Streaming) Synopsis() string {
	return "Get streaming data of ask, bid (default: USD_JPY)"
}
