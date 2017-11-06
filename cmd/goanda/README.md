# goanda

goanda is command line tool of using [exchange-api-go](https://github.com/mmm888/exchange-api-go)

## Usage

**Require setting UserID and API Token for OANDA API**

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
* Use the concurrency in "goanda past"
