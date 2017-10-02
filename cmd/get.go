package main

import (
	"strings"

	"fmt"

	"log"

	"github.com/mmm888/exchange-api-go"
)

type Get struct{}

func (g *Get) Help() string {
	return "Usage: otoi get <code> <option>"
}

func (g *Get) Run(args []string) int {
	if len(args) == 0 {
		log.Println("Not enough argument")
		return 1
	}
	instruments := strings.Split(args[0], ",")

	var layout, since string
	if len(args) == 1 {
		layout = ""
		since = ""
	}

	d := new(exchange.OANDACurrentData)
	d.SetData(instruments, layout, since)
	data := d.GetData()
	for _, v := range data.Prices {
		fmt.Printf("%s: %f (Bid), %f (Ask)\n", v.Instrument, v.Bid, v.Ask)
	}
	return 0
}

func (g *Get) Synopsis() string {
	return "Show current value"
}
