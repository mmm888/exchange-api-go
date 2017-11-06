package exchange

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type OANDAPastData struct {
	url         string
	layout      string
	pairCode    string
	start       string
	end         string
	granularity string
}

type PastData struct {
	PairCode    string   `json:"instrument"`
	Granularity string   `json:"granularity"`
	Candles     []Candle `json:"candles"`
}

type Candle struct {
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
}

func (pd *OANDAPastData) SetData(layout, pairCode, start, end, granularity string) {
	pd.url = pastURL
	pd.layout = layout
	pd.pairCode = pairCode
	pd.start = start
	pd.end = end
	pd.granularity = granularity
}

func (pd *OANDAPastData) GetData() (*PastData, error) {
	resp, err := pd.GetResponse()
	if err != nil {
		return nil, errors.Wrap(err, "Error1 at PastData")
	}
	defer resp.Body.Close()

	var data PastData
	err = GetUnmarshal(resp.Body, &data)
	if err != nil {
		return nil, errors.Wrap(err, "Error2 at PastData")
	}

	return &data, nil
}

func (pd *OANDAPastData) GetResponse() (*http.Response, error) {
	s, err := time.Parse(pd.layout, pd.start)
	if err != nil {
		return nil, &ParseTimeError{}
	}
	e, err := time.Parse(pd.layout, pd.end)
	if err != nil {
		return nil, &ParseTimeError{}
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
		return nil, &CreateReqError{}
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.URL.RawQuery = values.Encode()

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, &GetRespError{}
	}

	return resp, nil
}
