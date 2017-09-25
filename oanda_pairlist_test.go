package exchange

import (
	"testing"
)

func TestOANDAPairList_GetData(t *testing.T) {
	var f, i []string
	f = []string{
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
	i = []string{
		"AUD_CAD",
		"AUD_CHF",
		"USD_JPY",
	}

	d := new(OANDAPairList)
	d.SetData(i, f)
	data := d.GetData()
	t.Log(data)
}
