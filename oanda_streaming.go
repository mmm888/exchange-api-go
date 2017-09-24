package exchange

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Dummy struct {
	HeartBeat struct {
		Time time.Time `json:"time"`
	} `json:"heartbeat""`
}

type StreamingData struct {
	Tick struct {
		PairCode string    `json:"instrument"`
		Time     time.Time `json:"time"`
		Bid      float64   `json:"bid"`
		Ask      float64   `json:"ask"`
	} `json:tick`
}

type OANDAStreamData struct {
	url      string
	pairCode []string
	Chan     chan StreamingData
}

//
// {"tick":{"instrument":"USD_JPY","time":"2017-09-08T20:59:58.315562Z","bid":107.832,"ask":107.858}}
// {"heartbeat":{"time":"2017-09-11T07:12:35.258498Z"}}
//
func (sd *OANDAStreamData) SetData(code []string) {
	sd.url = streamURL
	sd.pairCode = code
	sd.Chan = make(chan StreamingData)
}

func (sd *OANDAStreamData) GetData() {
	go func() {
		resp, err := sd.GetResponse()
		if err != nil {
			log.Printf("Cannot get body: %v", err)
		}
		defer resp.Body.Close()

		var prevTime time.Time
		var data StreamingData
		var d Dummy
		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				log.Printf("Cannot read line: %v", err)
			}

			err = json.Unmarshal(line, &d)
			if err != nil {
				log.Printf("Cannot decode json to Dummy: %v", err)
			}

			if prevTime == d.HeartBeat.Time {
				err = json.Unmarshal(line, &data)
				if err != nil {
					log.Printf("Cannot decode json to Change: %v", err)
				}

				sd.Chan <- data
			}
			prevTime = d.HeartBeat.Time
		}
	}()
}

func (sd *OANDAStreamData) GetResponse() (*http.Response, error) {
	values := url.Values{}
	values.Set("accountId", userID)
	values.Add("instruments", strings.Join(sd.pairCode, ","))

	req, err := http.NewRequest("GET", sd.url, nil)
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
