package exchange

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

var (
	streamURL  = "https://stream-fxpractice.oanda.com/v1/prices"
	currentURL = "https://api-fxpractice.oanda.com/v1/prices"
	pastURL    = "https://api-fxpractice.oanda.com/v1/candles"
	listURL    = "https://api-fxpractice.oanda.com/v1/instruments"
)

func GetUnmarshal(respBody io.ReadCloser, data interface{}) error {
	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}
