package main

import (
	"fmt"
	"log"

	"flag"

	"strings"

	ex "github.com/mmm888/exchange-api-go"
)

type Current struct{}

func (l *Current) Help() string {
	return `Usage: goanda current [option]

Options:
  -code string
        Pair Code (default "USD_JPY"")
  -layout string
        Time format for Golang
  -since string
        Start date
`
}

func (l *Current) Run(args []string) int {
	var code, layout, since string
	f := flag.NewFlagSet("", flag.ContinueOnError)
	f.StringVar(&code, "code", "USD_JPY", "Pair Code")
	f.StringVar(&layout, "layout", "", "Time format for Golang")
	f.StringVar(&since, "since", "", "Start date")
	if err := f.Parse(args); err != nil {
		log.Print(err)
		return 1
	}

	instruments := strings.Split(code, ",")
	d := new(ex.OANDACurrentData)
	d.SetData(instruments, "", "")
	data, err := d.GetData()
	if err != nil {
		log.Printf("Error: %s", err)
		return 1
	}

	for _, v := range data.Prices {
		fmt.Printf("Code: %s, Ask: %10f, Bid: %10f\n", v.Instrument, v.Ask, v.Bid)
	}

	return 0
}

func (l *Current) Synopsis() string {
	return "Get ask, bid now (default: USD_JPY)"
}
