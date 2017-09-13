package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	streamURL = "https://stream-fxpractice.oanda.com/v1/prices"
	pastURL   = "https://api-fxpractice.oanda.com/v1/candles"
)

func OANDARequest(URL string, option ...string) (*http.Response, error) {
	paircode := option[0]
	layout := "2006-01-02 15:04:05"
	var start, end time.Time
	if len(option) == 3 {
		start, _ = time.Parse(layout, option[1])
		end, _ = time.Parse(layout, option[2])
	}

	values := url.Values{}
	values.Set("accountId", userID)

	if URL == streamURL {
		values.Add("instruments", paircode)
	} else if URL == pastURL {
		values.Add("instrument", paircode)
	}

	if !start.IsZero() {
		values.Add("start", fmt.Sprint(start.Format(time.RFC3339)))
	}

	if !end.IsZero() {
		values.Add("end", fmt.Sprint(end.Format(time.RFC3339)))
	}

	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.URL.RawQuery = values.Encode()

	client := new(http.Client)
	resp, err := client.Do(req)

	return resp, err
}

type PastData struct {
	PairCode    string `json:"instrument"`
	Granularity string `json:"granularity"`
	Candles     []struct {
		Time     time.Time `json:time`
		OpenBid  float64   `json:openBid`
		OpenAsk  float64   `json:openAsk`
		HighBid  float64   `json:highBid`
		HighAsk  float64   `json:highAsk`
		LowBid   float64   `json:lowBid`
		LowAsk   float64   `json:lowAsk`
		CloseBid float64   `json:closeBid`
		CloseAsk float64   `json:closeAsk`
		Volume   int       `json:volume`
		Complete bool      `json:complete`
	} `json:"candles"`
}

func GetData(c chan PastData, code, start, end string) {
	resp, err := OANDARequest(pastURL, code, start, end)
	if err != nil {
		log.Printf("Cannot get body: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data PastData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Cannot decode json: %v", err)
	}
	fmt.Println(data)
}

func main() {
	c := make(chan PastData)
	GetData(c, "USD_JPY", "2016-06-03 15:20:33", "2016-06-03 15:23:33")
	/*
		d := new(OANDAStreamData)
		SetData(d, "USD_JPY")
		for {
			fmt.Println(<-d.c)
		}
	*/
}
