package main

import (
	"fmt"

	"github.com/mmm888/exchange-api-go"
)

var (
	// oanda-api
	code        = "USD_JPY"
	start       = "2014-06"
	end         = "2017-06"
	granularity = "D"

	// influxdb
	db    = "exchange"
	table = "USDJPY"
	api   = "oanda"
)

func main() {
	//	d := GetData("USD_JPY", "2016-06-03 5:53:09", "2016-06-03 15:40:33")
	//	d := exchange.GetData(code, start, end)

	d := new(exchange.OANDAPastData)
	d.SetData(code, start, end, granularity)
	data := d.GetData()

	var text string
	text = fmt.Sprintf("# DDL\n")
	text += fmt.Sprintf("CREATE DATABASE %s\n\n", db)
	text += fmt.Sprintf("# DML\n")
	text += fmt.Sprintf("# CONTEXT-DATABASE: %s\n", db)
	for _, v := range data.Candles {
		text += fmt.Sprintf("%s,api=%s bid=%f,ask=%f %d\n", table, api, v.OpenBid, v.OpenAsk, v.Time.Unix())
	}

	fmt.Println(text)
}
