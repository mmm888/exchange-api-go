package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

//{"tick":{"instrument":"USD_JPY","time":"2017-09-08T20:59:58.315562Z","bid":107.832,"ask":107.858}}
//{"heartbeat":{"time":"2017-09-11T07:12:35.258498Z"}}

type Dummy struct {
	HeartBeart struct {
		Time time.Time `json:"time"`
	} `json:"heartbeat""`
}

type Change struct {
	Tick struct {
		PairCode string    `json:"instrument"`
		Time     time.Time `json:"time"`
		Bid      float64   `json:"bid"`
		Ask      float64   `json:"ask"`
	} `json:tick`
}

var (
	streamURL = "https://stream-fxpractice.oanda.com/v1/prices"
)

func main() {
	values := url.Values{}
	values.Set("accountId", userID)
	values.Add("instruments", "USD_JPY")

	req, _ := http.NewRequest("GET", streamURL, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.URL.RawQuery = values.Encode()

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Cannot get resp: %v", err)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	var c Change
	var d Dummy
	var prevtime time.Time
	var count int
	for {
		line, _ := reader.ReadBytes('\n')
		err := json.Unmarshal(line, &d)
		if err != nil {
			log.Printf("Cannot decode json to Dummy: %v", err)
		}
		if prevtime == d.HeartBeart.Time {
			err = json.Unmarshal(line, &c)
			if err != nil {
				log.Printf("Cannot decode json to Change: %v", err)
			}
			fmt.Println(c)
		}
		prevtime = d.HeartBeart.Time
	}
}
