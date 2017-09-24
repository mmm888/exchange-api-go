package exchange

import (
	"bufio"
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

type OANDAStreamData struct {
	paircode string
	c        chan Data
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

// {"tick":{"instrument":"USD_JPY","time":"2017-09-08T20:59:58.315562Z","bid":107.832,"ask":107.858}}
// {"heartbeat":{"time":"2017-09-11T07:12:35.258498Z"}}
//
//	d := new(OANDAStreamData)
//	SetData(d, "USD_JPY")
//	for {
//	fmt.Println(<-d.c)
//	}
//
func (sd OANDAStreamData) GetData() {
	resp, err := OANDARequest(streamURL, sd.paircode)
	if err != nil {
		log.Printf("Cannot get body: %v", err)
	}
	defer resp.Body.Close()

	var data Data
	var d Dummy
	var prevtime time.Time
	reader := bufio.NewReader(resp.Body)
	for {
		line, _ := reader.ReadBytes('\n')
		err := json.Unmarshal(line, &d)
		if err != nil {
			log.Printf("Cannot decode json to Dummy: %v", err)
		}
		if prevtime == d.HeartBeart.Time {
			err = json.Unmarshal(line, &data)
			if err != nil {
				log.Printf("Cannot decode json to Change: %v", err)
			}
			sd.c <- data
		}
		prevtime = d.HeartBeart.Time
	}
}

func (sd OANDAStreamData) Get() {
	go func() {
		sd.GetData()
	}()
}

func (sd *OANDAStreamData) Init(code string) {
	sd.paircode = code
	sd.c = make(chan Data)
}

func OANDARequest(URL string, option ...string) (*http.Response, error) {
	paircode := option[0]
	//	layout := "2006-01-02 15:04:05"
	layout := "2006-01"
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
		// TODO: 変数に置く
		values.Add("granularity", "D")
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

func GetData(code, start, end string) PastData {
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
	return data
}
