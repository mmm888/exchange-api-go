package exchange

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

type PairList struct {
	Instruments []struct {
		Instrument      string  `json:"instrument"`
		DisplayName     string  `json:"displayName"`
		Pip             string  `json:"pip"`
		MaxTradeUnits   int     `json:"maxTradeUnits"`
		Precision       string  `json:"precision"`
		MaxTrailingStop float64 `json:"maxTrailingStop"`
		MinTrailingStop float64 `json:"minTrailingStop"`
		MarginRate      float64 `json:"marginRate"`
		Halted          bool    `json:"halted"`
	} `json:"instruments"`
}s

type OANDAPairList struct {
	url         string
	instruments []string
	fields      []string
}

func (pl *OANDAPairList) SetData(instruments, fields []string) {
	pl.url = listURL
	pl.instruments = instruments
	pl.fields = fields
}

func (pl *OANDAPairList) GetResponse() (*http.Response, error) {
	values := url.Values{}
	values.Set("accountId", userID)

	if pl.instruments != nil {
		values.Add("instruments", strings.Join(pl.instruments, ","))
	}
	if pl.fields != nil {
		values.Add("fields", strings.Join(pl.fields, ","))
	}

	req, err := http.NewRequest("GET", pl.url, nil)
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

func (pl *OANDAPairList) GetData() PairList {
	resp, err := pl.GetResponse()
	if err != nil {
		log.Printf("Cannot get body: %v", err)
	}
	defer resp.Body.Close()

	var data PairList
	err = GetUnmarshal(resp.Body, &data)
	if err != nil {
		log.Printf("Cannot get unmarshal data: %v", err)
	}

	return data
}
