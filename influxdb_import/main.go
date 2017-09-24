package main

import (
	"fmt"

	"github.com/mmm888/exchange-api-go"
)

var (
	// oanda-api
	code  = "USD_JPY"
	start = "2014-06"
	end   = "2017-06"
	// influxdb
	db    = "exchange"
	table = "USDJPY"
)

func main() {
	//	d := GetData("USD_JPY", "2016-06-03 5:53:09", "2016-06-03 15:40:33")
	d := exchange.GetData(code, start, end)

	var exportdata string
	exportdata = fmt.Sprintf("# DDL\n")
	exportdata += fmt.Sprintf("CREATE DATABASE %s\n\n", db)
	exportdata += fmt.Sprintf("# DML\n")
	exportdata += fmt.Sprintf("# CONTEXT-DATABASE: %s\n", db)
	for _, v := range d.Candles {
		exportdata += fmt.Sprintf("%s,bid=%f ask=%f %d\n", table, v.OpenBid, v.OpenAsk, v.Time.Unix())
	}

	fmt.Println(exportdata)
}
