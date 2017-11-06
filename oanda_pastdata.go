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
	start       time.Time
	end         time.Time
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

func (pd *OANDAPastData) SetData(layout, pairCode, start, end, granularity string) error {
	var err error
	pd.url = pastURL
	pd.layout = layout
	pd.pairCode = pairCode
	pd.granularity = granularity

	pd.start, err = time.Parse(pd.layout, start)
	if err != nil {
		return &ParseTimeError{}
	}
	pd.end, err = time.Parse(pd.layout, end)
	if err != nil {
		return &ParseTimeError{}
	}

	return nil
}

func (pd *OANDAPastData) GetData() (*PastData, error) {
	// 5000 件以上返ってくる場合、nil を返すので。1 年ごとに値を取得する
	// 1 年ごとの場合、Granularity=H2 まで値を取得できる (Granularity=H1 の場合は nil)
	// 2000 年から今の時間までの間の値を取得する
	minStart, _ := time.Parse("2006", "2000")
	maxEnd := time.Now()

	s := pd.start
	e := pd.end
	if pd.start.IsZero() || pd.start.Before(minStart) {
		s = minStart
	}
	if pd.end.IsZero() || pd.end.After(maxEnd) {
		e = maxEnd
	}

	var data, tmpData PastData
	tmpStart := s
	tmpEnd := tmpStart.AddDate(1, 0, 0)
	for {
		if tmpEnd.After(e) {
			tmpEnd = e
		}

		resp, err := pd.GetResponse(tmpStart, tmpEnd)
		if err != nil {
			return nil, errors.Wrap(err, "Error1 at PastData")
		}

		err = GetUnmarshal(resp.Body, &tmpData)
		if err != nil {
			return nil, errors.Wrap(err, "Error2 at PastData")
		}
		resp.Body.Close()

		if data.Candles == nil {
			data = tmpData
			data.PairCode = tmpData.PairCode
			data.Granularity = tmpData.Granularity
		}
		data.Candles = append(data.Candles, tmpData.Candles...)

		if tmpEnd == e {
			break
		}
		tmpStart = tmpEnd
		tmpEnd = tmpStart.AddDate(1, 0, 0)
	}

	return &data, nil
}

func (pd *OANDAPastData) GetResponse(start, end time.Time) (*http.Response, error) {
	values := url.Values{}
	values.Set("accountId", userID)
	values.Add("instrument", pd.pairCode)
	values.Add("start", fmt.Sprint(start.Format(time.RFC3339)))
	values.Add("end", fmt.Sprint(end.Format(time.RFC3339)))
	values.Add("granularity", pd.granularity)

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
