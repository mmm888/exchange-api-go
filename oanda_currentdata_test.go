package exchange

import "testing"

func TestOANDACurrentData_GetData(t *testing.T) {
	var i []string
	var l, s string
	i = []string{
		"EUR_USD",
		"USD_JPY",
	}
	l = "2006-01-02"
	s = "2017-09-25"

	d := new(OANDACurrentData)
	d.SetData(i, l, s)
	data := d.GetData()
	t.Log(data)
}
