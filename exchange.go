package main

import "time"

type Dummy struct {
	HeartBeart struct {
		Time time.Time `json:"time"`
	} `json:"heartbeat""`
}

type Data struct {
	Tick struct {
		PairCode string    `json:"instrument"`
		Time     time.Time `json:"time"`
		Bid      float64   `json:"bid"`
		Ask      float64   `json:"ask"`
	} `json:tick`
}

type Exchange interface {
	GetData()
	Get()
	Init(string)
}

func SetData(exchange Exchange, pair string) {
	exchange.Init(pair)
	exchange.Get()
}
