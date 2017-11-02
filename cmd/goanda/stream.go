package main

/*
import (
	"fmt"
	"log"

	"strings"

	"github.com/mmm888/exchange-api-go"
)

type Stream struct{}

func (s *Stream) Help() string {
	return "Usage: otoi stream <code>"
}

func (s *Stream) Run(args []string) int {
	if len(args) == 0 {
		log.Println("Not enough argument")
		return 1
	}
	code := strings.Split(args[0], ",")

	// Get streaming data
	d := new(exchange.OANDAStreamData)
	d.SetData(code)
	d.GetData()
	var data exchange.StreamingData
	for {
		data = <-d.Chan

		fmt.Println(data)
	}

	return 0
}

func (s *Stream) Synopsis() string {
	return "Insert streaming data for influxdb"
}
*/
