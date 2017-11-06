# exchange-api-go

exchange-api-go is Go wrapper for the OANDA API

* OANDA API Document
  * http://developer.oanda.com/rest-live/development-guide/
  * http://developer.oanda.com/docs/jp/v1/guide/

## Precondition

Set UserID and API Token for OANDA API

* vi secret.go

~~~
package main

var (
    userID = "REPLACE THIS WITH YOUR ACCOUNT ID, ie 1234567"
    token  = "REPLACE THIS WITH YOUR ACCESS TOKEN"
)
~~~

## Usage

### Get an instrument list

* Code

~~~
package main

import (
    "fmt"
    "log"

    ex "github.com/mmm888/exchange-api-go"
)

func main() {
    var cpr = &ex.PairList{}
    var instruments = []string{
        "USD_JPY",
        "EUR_JPY",
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

    pair := new(ex.OANDAPairList)
    pair.SetData(instruments, fields)
    pr, err := pair.GetData()
    if err != nil {
        log.Print(err)
    }

    for _, v := range pr.Instruments {
        if v == cpr.Instruments[0] {
            break
        }

        fmt.Printf("Code: %s, Pip: %s\n", v.Instrument, v.Pip)
    }
}
~~~

* Result

~~~
Code: USD_JPY, Pip: 0.01
Code: EUR_JPY, Pip: 0.01
~~~

### Get Current prices

* Code

~~~
package main

import (
    "fmt"
    "log"

    ex "github.com/mmm888/exchange-api-go"
)

func main() {
    var cc = &ex.CurrentData{}

    var instruments = []string{
        "USD_JPY",
        "EUR_JPY",
    }
    var layout = ""
    var since = ""

    current := new(ex.OANDACurrentData)
    current.SetData(instruments, layout, since)
    c, err := current.GetData()
    if err != nil {
        log.Print(err)
    }

    for _, v := range c.Prices {
        if v == cc.Prices[0] {
            break
        }

        fmt.Printf("Code: %s, Ask: %f Bid: %f\n", v.Instrument, v.Ask, v.Bid)
    }
}
~~~

* Result

~~~
Code: USD_JPY, Ask: 114.034000 Bid: 114.030000
Code: EUR_JPY, Ask: 132.725000 Bid: 132.712000
~~~

### Retrieve instrument history

**Cannot get over 5000 historical information (then, the return value is nil)**

* Code

~~~
package main
(default: USD_JPY)
import (
    "fmt"
    "log"

    ex "github.com/mmm888/exchange-api-go"
)

func main() {
    var cps = &ex.PastData{}

    var pairCode = "USD_JPY"
    var layout = "2006-01-02"
    var start = "2017-10-25"
    var end = "2017-10-26"
    var granularity = "D"

    past := new(ex.OANDAPastData)
    past.SetData(layout, pairCode, start, end, granularity)
    ps, err := past.GetData()
    if err != nil {
        log.Print(err)
    }

    fmt.Printf("Code: %s\n", pairCode)
    for _, v := range ps.Candles {
        if v == cps.Candles[0] {
            break
        }

        fmt.Printf("Time: %s, Ask: %f, Bid: %f\n", v.Time.Format(layout), v.OpenAsk, v.OpenBid)
    }
}
~~~

* Result

~~~
Code: USD_JPY
Time: 2017-10-24, Ask: 113.927000, Bid: 113.896000
Time: 2017-10-25, Ask: 113.783000, Bid: 113.718000
~~~

### Rates Streaming

* Code

~~~
package main

import (
	"fmt"

	ex "github.com/mmm888/exchange-api-go"
)

func main() {
	var instruments = []string{
		"USD_JPY",
	}
	var format = "2006-01-02 03:04:05"

	s := new(ex.OANDAStreamData)
	s.SetData(instruments)
	s.GetData()

	for {
		d := <-s.Chan
		fmt.Printf("Time: %s, Code: %s, Ask: %f, Bid: %f\n",
			d.Tick.Time.Format(format), d.Tick.PairCode, d.Tick.Ask, d.Tick.Bid)
	}
}
~~~

* Result

~~~
Time: 2017-11-02 10:00:26, Code: USD_JPY, Ask: 114.060000, Bid: 114.056000
Time: 2017-11-02 10:00:29, Code: USD_JPY, Ask: 114.057000, Bid: 114.053000
Time: 2017-11-02 10:00:31, Code: USD_JPY, Ask: 114.059000, Bid: 114.055000
...
~~~

## TODO

* Get more than 5000 historical information for "Retrieve instrument history"
* Add Forex Labs
  * http://developer.oanda.com/rest-live/forex-labs/
