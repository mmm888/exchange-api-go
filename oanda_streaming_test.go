package exchange

import (
	"testing"

	"log"

	"github.com/pkg/errors"
)

func TestOANDAStreamData_GetDataTest(t *testing.T) {
	//
	// Streaming Data
	// {"tick":{"instrument":"USD_JPY","time":"2017-09-08T20:59:58.315562Z","bid":107.832,"ask":107.858}}
	// Dummy Data
	// {"heartbeat":{"time":"2017-09-11T07:12:35.258498Z"}}
	//

	var err error
	var pairCode = []string{
		"USD_JPY",
	}

	// Check Streaming data
	streamData := "{\"tick\":{\"instrument\":\"USD_JPY\",\"time\":\"2017-09-08T20:59:58.315562Z\",\"bid\":107.832,\"ask\":107.858}}"

	s := new(OANDAStreamData)
	s.SetData(pairCode)
	StreamStruct, _, err := s.GetDataTest(streamData)

	var checkStream = &StreamingData{}
	if *StreamStruct == *checkStream {
		t.Error("Check Streaming data: Streaming data is nil")
	}

	if err != nil {
		t.Error(errors.Wrap(err, "Error1 at TestOANDAStreamData_GetDataTest"))
	}

	// Check Dummy Data
	dummyData := "{\"heartbeat\":{\"time\":\"2017-09-11T07:12:35.258498Z\"}}"

	d := new(OANDAStreamData)
	d.SetData(pairCode)
	_, dummyStruct, err := d.GetDataTest(dummyData)

	var checkDummy = &DummyData{}
	if *dummyStruct == *checkDummy {
		t.Error("Check Dummy Data: Dummy data is nil")
	}

	if err != nil {
		t.Error(errors.Wrap(err, "Error2 at TestOANDAStreamData_GetDataTest"))
	}

	log.Print("StreamData test finished")

}
