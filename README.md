# CoinMarketCap API client library

Simple API client library for the CoinMarketCap API, written in Go

## Installation

This is not a standalone app, use it in your project as a package. To get it simply run;
`go get github.com/shaunmza/coinmarketcap`

## Usage

There is an example file in the `examples` directory, it should be pretty much
self explanatory, snippet from the file below;

```go

var r coinmarketcap.Ticker
var err error

// Which coins do you want to watch?
t := make([]string, 0)
t = append(t, "bitcoin")
t = append(t, "litecoin")
t = append(t, "steem")
t = append(t, "steem-dollars")

// Endpoints are updated every 5 minutes, se we use that here
period := 10 // 60 * 5
ticker := time.NewTicker(time.Second * time.Duration(period))

// Because we are impatient, call it now
r, err = coinmarketcap.GetData(t)

// If this is not nil then we encountered a problem, use this to determine
// what to do next.
// LastUpdate can be used to determine how stale the data is
if err != nil {
  fmt.Printf("Error! %s, Last Updated: %s\n", err, r.LastUpdate)
}

printData(r)

// Infinite loop so we keep getting prices
for _ = range ticker.C {
  // Get off of the channel
  r, err = coinmarketcap.GetData(t)

  // If this is not nil then we encountered a problem, use this to determine
  // what to do next.
  // LastUpdate can be used to determine how stale the data is
  if err != nil {
    fmt.Printf("Error! %s, Last Updated: %s\n", err, r.LastUpdate)
  }
  printData(r)

}

```

The values for the coins to watch are the CoinMarketCap id's, which can be seen in the url when viewing details on the site.
An example for Bitcoin; 'coinmarketcap.com/currencies/__bitcoin__/'

## Authors

* **Shaun Morrow** - *Initial work* - [shaunmza](https://github.com/shaunmza)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
