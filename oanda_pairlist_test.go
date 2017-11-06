package exchange

import (
	"testing"

	"log"

	"github.com/pkg/errors"
)

func TestOANDAPairList_GetData(t *testing.T) {
	var instruments = []string{
		"USD_JPY",
		"EUR_JPY",
		"EUR_USD",
	}
	var fields = []string{
		"instrument",
		"displayName",
		"pip",
		"maxTradeUnits",
		"precision",
		"maxTrailingStop",
		"minTrailingStop",
		"marginRate",
		"halted",
	}

	d := new(OANDAPairList)
	d.SetData(instruments, fields)
	p, err := d.GetData()

	if p.Instruments == nil {
		t.Error("PairList is nil")
	}

	if err != nil {
		t.Error(errors.Wrap(err, "Error1 at TestOANDAPairList_GetData"))
	}

	log.Print("PairList test finished")
}
