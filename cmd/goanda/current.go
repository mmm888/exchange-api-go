package main

import (
	"fmt"
	"log"

	ex "github.com/mmm888/exchange-api-go"
)

type Current struct{}

func (l *Current) Help() string {
	return "Usage: goanda current"
}

func (l *Current) Run(args []string) int {
	instruments := []string{
		"USD_JPY",
		"EUR_JPY",
		"EUR_USD",
	}

	d := new(ex.OANDACurrentData)
	d.SetData(instruments, "", "")
	data, err := d.GetData()
	if err != nil {
		log.Print(err)
		return 1
	}

	checkNil := &ex.CurrentData{}
	for _, v := range data.Prices {
		if v == checkNil.Prices[0] {
			break
		}

		fmt.Printf("Code: %s, Ask: %10f, Bid: %10f\n", v.Instrument, v.Ask, v.Bid)
	}

	return 0
}

func (l *Current) Synopsis() string {
	return "Get ask, bid now"
}
