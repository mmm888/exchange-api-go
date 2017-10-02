package main

import (
	"fmt"
	"log"

	"github.com/mmm888/exchange-api-go"
)

type Insert struct{}

func (i *Insert) Help() string {
	return "Usage: otoi insert <code> <start> <end>"
}

func (i *Insert) Run(args []string) int {
	if len(args) < 3 {
		log.Println("Not enough argument")
	}

	var (
		code        = args[0]
		start       = args[1]
		end         = args[2]
		granularity = config.Granularity
		layout      = config.Layout
	)

	if len(args) >= 5 {
		granularity = args[3]
		layout = args[4]
	}

	d := new(exchange.OANDAPastData)
	d.SetData(layout, code, start, end, granularity)
	data := d.GetData()

	fmt.Println(data)

	return 0
}

func (i *Insert) Synopsis() string {
	return "Insert data for influxdb"
}
