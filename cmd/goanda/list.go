package main

import (
	"fmt"
	"log"

	ex "github.com/mmm888/exchange-api-go"
)

type List struct{}

func (l *List) Help() string {
	return "Usage: goanda list"
}

func (l *List) Run(args []string) int {
	fields := []string{
		"instrument",
	}

	d := new(ex.OANDAPairList)
	d.SetData(nil, fields)
	data, err := d.GetData()
	if err != nil {
		log.Printf("Error: %s", err)
		return 1
	}

	checkNil := &ex.PairList{}
	for _, v := range data.Instruments {
		if v == checkNil.Instruments[0] {
			break
		}

		fmt.Println(v.Instrument)
	}

	return 0
}

func (l *List) Synopsis() string {
	return "Show a list of pair code"
}
