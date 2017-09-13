package main

import (
	"bufio"
	"encoding/json"
	"log"
	"time"
)

type OANDAStreamData struct {
	paircode string
	c        chan Data
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
