package exchange

import (
	"log"
	"testing"
)

func TestOANDAPastData_GetDatatData_GetData(t *testing.T) {
	var checkPastData = &PastData{}

	var pairCode = "USD_JPY"
	var layout = "2006-01-02"
	var start = "2017-09-25"
	var end = "2017-10-25"
	var granularity = "H1"

	d := new(OANDAPastData)
	d.SetData(layout, pairCode, start, end, granularity)
	p, err := d.GetData()

	if *p == *checkPastData {
		t.Error("Past Data is nil")
	}

	if err != nil {
		t.Error(err, "Error1 at TestOANDAPastData_GetData")
	}

	log.Print("PastData test finished")
}
