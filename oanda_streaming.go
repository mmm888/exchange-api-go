package exchange

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Stream interface{}

type DummyData struct {
	HeartBeat struct {
		Time time.Time `json:"time"`
	} `json:"heartbeat"`
}

type StreamingData struct {
	Tick struct {
		PairCode string    `json:"instrument"`
		Time     time.Time `json:"time"`
		Bid      float64   `json:"bid"`
		Ask      float64   `json:"ask"`
	} `json:"tick"`
}

type OANDAStreamData struct {
	url      string
	pairCode []string
	Chan     chan StreamingData
}

func (sd *OANDAStreamData) SetData(code []string) {
	sd.url = streamURL
	sd.pairCode = code
	sd.Chan = make(chan StreamingData)
}

func (sd *OANDAStreamData) GetData() {
	go func() {
		resp, err := sd.GetResponse()
		if err != nil {
			log.Print(errors.Wrap(err, "Error1 at StreamData"))
		}
		defer resp.Body.Close()

		var prevTime time.Time
		var data StreamingData
		var d DummyData
		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				log.Print(errors.Wrap(&ReadBytesError{}, "Error2 at StreamData"))
			}

			err = json.Unmarshal(line, &d)
			if err != nil {
				log.Print(errors.Wrap(&UnmarshalError{}, "Error3 at StreamData"))
			}

			if prevTime == d.HeartBeat.Time {
				err = json.Unmarshal(line, &data)
				if err != nil {
					log.Print(errors.Wrap(&UnmarshalError{}, "Error4 at StreamData"))
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

func (sd *OANDAStreamData) GetDataTest(body string) (*StreamingData, *DummyData, error) {
	var s StreamingData
	var d DummyData

	/*
		reader := bufio.NewReader(resp.Body)
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Printf("Cannot read line: %v", err)
		}
	*/

	err := json.Unmarshal([]byte(body), &d)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Error1 at StreamDataTest")
	}

	err = json.Unmarshal([]byte(body), &s)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Error2 at StreamDataTest")
	}

	return &s, &d, nil
}
