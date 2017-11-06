# goanda

goanda is command line tool of using [exchange-api-go](https://github.com/mmm888/exchange-api-go)


## Build

**Require setting UserID and API Token for OANDA API**

~~~
git clone https://github.com/mmm888/exchange-api-go
cat << '_EOF_' > secret.go
package main

var (
    userID = "REPLACE THIS WITH YOUR ACCOUNT ID, ie 1234567"
    token  = "REPLACE THIS WITH YOUR ACCESS TOKEN"
)
_EOF_
cd cmd/goanda
go build .
mv goanda ${GOPATH}/bin
~~~

## Usage

"goanda influxdb" is used [here](https://github.com/mmm888/exchange-api-docker)

~~~
$ goanda --help
Usage: goanda [--version] [--help] <command> [<args>]

Available commands are:
    current      Get ask, bid now (default: USD_JPY)
    influxdb     Operate InfluxDB
    list         Show a list of pair code
    past         Get ask, bid from start to end (default: USD_JPY)
    streaming    Get streaming data of ask, bid (default: USD_JPY)
~~~

## TODO

* Add test code
* Get more than 5000 past data for "oanda past"
