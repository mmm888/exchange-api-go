package exchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type OANDAPastData struct {
	url         string
	pairCode    string
	start       string
	end         string
	granularity string
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

func (pd *OANDAPastData) SetData(pairCode, start, end, granularity string) {
	pd.url = pastURL
	pd.pairCode = pairCode
	pd.start = start
	pd.end = end
	pd.granularity = granularity
}

func (pd *OANDAPastData) GetResponse() (*http.Response, error) {
	//	layout := "2006-01-02 15:04:05"
	layout := "2006-01"
	s, err := time.Parse(layout, pd.start)
	if err != nil {
		return nil, err
	}
	e, err := time.Parse(layout, pd.end)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Set("accountId", userID)
	values.Add("instrument", pd.pairCode)
	values.Add("granularity", pd.granularity)

	if !s.IsZero() {
		values.Add("start", fmt.Sprint(s.Format(time.RFC3339)))
	}
	if !e.IsZero() {
		values.Add("end", fmt.Sprint(e.Format(time.RFC3339)))
	}

	req, err := http.NewRequest("GET", pd.url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.URL.RawQuery = values.Encode()

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (pd *OANDAPastData) GetData() PastData {
	resp, err := pd.GetResponse()
	if err != nil {
		log.Printf("Cannot get body: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Cannot read body: %v", err)
	}

	var data PastData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Cannot decode json: %v", err)
	}

	return data
}
