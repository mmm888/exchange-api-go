package main

import (
	"fmt"

	"github.com/mmm888/exchange-api-go"
)

var (
	code = []string{
		"USD_JPY",
	}
)

func main() {
	d := new(exchange.OANDAStreamData)
	d.SetData(code)
	d.GetData()
	for {
		fmt.Println(<-d.Chan)

	}
}
