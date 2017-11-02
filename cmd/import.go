package main

/*
import (
	"fmt"

	"log"

	"github.com/mmm888/exchange-api-go"
)

type Import struct{}

func (i *Import) Help() string {
	return "Usage: otoi import <code> <start> <end>"
}

func (i *Import) Run(args []string) int {
	if len(args) < 3 {
		log.Println("Not enough argument")
		return 1
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

		var text string
		text = fmt.Sprintf("# DDL\n")
		text += fmt.Sprintf("CREATE DATABASE %s\n\n", db)
		text += fmt.Sprintf("# DML\n")
		text += fmt.Sprintf("# CONTEXT-DATABASE: %s\n", db)
		for _, v := range data.Candles {
			text += fmt.Sprintf("%s,api=%s bid=%f,ask=%f %d\n", table, api, v.OpenBid, v.OpenAsk, v.Time.Unix())
		}

		fmt.Println(text)
	return 0
}

func (i *Import) Synopsis() string {
	return "Show import data for Influxdb"
}
*/
