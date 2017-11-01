package exchange

import (
	"log"
	"testing"
)

func TestOANDACurrentData_GetData(t *testing.T) {
	//var instruments = []string{
	//	"EUR_USD",
	//	"USD_JPY",
	//}
	//var layout = "2006-01-02"
	//var since = "2017-09-25"

	var checkCurrentData = &CurrentData{}

	var instruments = []string{
		"USD_JPY",
		"EUR_JPY",
		"EUR_USD",
	}
	var layout = ""
	var since = ""

	d := new(OANDACurrentData)
	d.SetData(instruments, layout, since)
	c, err := d.GetData()

	if *c == *checkCurrentData {
		t.Error("Current Data is nil")
	}

	if err != nil {
		t.Error(err, "Error1 at TestOANDACurrentData_GetData")
	}

	log.Print("CurrentData test finished")
}
