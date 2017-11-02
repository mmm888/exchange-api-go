package main

import (
	"fmt"
	"log"

	ex "github.com/mmm888/exchange-api-go"
)

func main() {
	var err error
	var instruments []string
	var fields []string
	var layout string

	// Get an instrument list
	fmt.Println("* Get an instrument list")
	var cpr = &ex.PairList{}
	instruments = []string{
		"USD_JPY",
		"EUR_JPY",
	}
	fields = []string{
		"instrument",
		"displayName",
		"pip",
		"maxTradeUnits",
		"precision",
		"maxTrailingStop",
		"minTrailingStop",
		"marginRate",
		"halted",
	}

	pair := new(ex.OANDAPairList)
	pair.SetData(instruments, fields)
	pr, err := pair.GetData()
	if err != nil {
		log.Print(err)
	}

	for _, v := range pr.Instruments {
		if v == cpr.Instruments[0] {
			break
		}

		fmt.Printf("Code: %s, Pip: %s\n", v.Instrument, v.Pip)
	}
	fmt.Println()

	// Get Current prices
	fmt.Println("* Get Current prices")
	var cc = &ex.CurrentData{}

	layout = ""
	var since = ""

	current := new(ex.OANDACurrentData)
	current.SetData(instruments, layout, since)
	c, err := current.GetData()
	if err != nil {
		log.Print(err)
	}

	for _, v := range c.Prices {
		if v == cc.Prices[0] {
			break
		}

		fmt.Printf("Code: %s, Ask: %f Bid: %f\n", v.Instrument, v.Ask, v.Bid)
	}
	fmt.Println()

	// Retrieve instrument history
	fmt.Println("* Retrieve instrument history")
	var cps = &ex.PastData{}

	var pairCode = "USD_JPY"
	layout = "2006-01-02"
	var start = "2017-10-25"
	var end = "2017-10-26"
	var granularity = "D"

	past := new(ex.OANDAPastData)
	past.SetData(layout, pairCode, start, end, granularity)
	ps, err := past.GetData()
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("Code: %s\n", pairCode)
	for _, v := range ps.Candles {
		if v == cps.Candles[0] {
			break
		}

		fmt.Printf("Time: %s, Ask: %f, Bid: %f\n", v.Time.Format(layout), v.OpenAsk, v.OpenBid)
	}

	fmt.Println()

	// Rates Streaming
	fmt.Println("* Rates Streaming")

	s := new(ex.OANDAStreamData)
	s.SetData(instruments)
	s.GetData()

	for {
		st := <-s.Chan
		fmt.Printf("Time: %s, Code: %s, Ask: %f, Bid: %f\n",
			st.Tick.Time.Format("2006-01-02 03:04:05"), st.Tick.PairCode, st.Tick.Ask, st.Tick.Bid)
	}

	fmt.Println()
}
