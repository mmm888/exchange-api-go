package main

import (
	"fmt"

	"time"

	"flag"

	"strings"

	"log"

	ex "github.com/mmm888/exchange-api-go"
)

type Streaming struct{}

func (s *Streaming) Help() string {
	return `Usage: goanda streaming [option]

Options:
    -code     Pair Code (default: "USD_JPY"")
`
}

func (s *Streaming) Run(args []string) int {
	var code string
	f := flag.NewFlagSet("", flag.ContinueOnError)
	f.StringVar(&code, "code", "USD_JPY", "Pair Code")

	if err := f.Parse(args); err != nil {
		log.Print(err)
		return 1
	}

	instruments := strings.Split(code, ",")
	d := new(ex.OANDAStreamData)
	d.SetData(instruments)
	d.GetData()

	format := "2006-01-02 15:04:05"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	for {
		data := <-d.Chan
		time := data.Tick.Time.In(jst)
		fmt.Printf("Time: %s, Code: %v, Ask: %10f, Bid: %10f\n",
			time.Format(format), data.Tick.PairCode, data.Tick.Ask, data.Tick.Bid)
	}

	return 0
}

func (s *Streaming) Synopsis() string {
	return "Get streaming data of ask, bid (default: USD_JPY)"
}
