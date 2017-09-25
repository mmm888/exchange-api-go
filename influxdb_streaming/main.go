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

//curl -i -XPOST "http://localhost:8086/write?db=science_is_cool" --data-binary 'weather,location=us-midwest temperature=82 1465839830100400200'

func main() {
	d := new(exchange.OANDAStreamData)
	d.SetData(code)
	d.GetData()
	for {
		fmt.Println(<-d.Chan)
	}
}
