# CoinMarketCap API client library

Simple API client library for the CoinMarketCap API, written in Go

## Installation

This is not a standalone app, use it in your project as a package. To get it simply run;
`go get github.com/shaunmza/coinmarketcap`

## Usage

There is an example file in the `examples` directory, it should be pretty much
self explanatory, snippet from the file below;

```go

var r data.Ticker

// Initialse, so we can get the channel to receive updates from
tChan := coinmarketcap.Init()

// Which coins do you want to watch?
t := make([]string, 0)
t = append(t, "bitcoin")
t = append(t, "litecoin")
t = append(t, "steem")
t = append(t, "steem-dollars")

// Start watching, second parameter is period between updates
// Endpoints are updated every 5 minutes, se we use that here
coinmarketcap.WatchCoins(t, 60*5)

// Infinite loop so we keep getting prices
for {
  // Get off of the channel
  r = <-tChan

  // If this is not nil then we encountered a problem, use this to determine
  // what to do next.
  // LastUpdate can be used to determine how stale the data is
  if r.Error != nil {
    fmt.Printf("Error! %s, Last Updated: %s\n", r.Error, r.LastUpdate)
  }

  // Just print it out for now
  for _, coin := range r.Coins {
    fmt.Println("Symbol: '" + coin.Symbol + "', Name: '" +
      coin.Name + "', Price Bitcoin: '" +
      coin.PriceBtc + "', Price USD: '" + coin.PriceUsd + "'")
  }
}

```

The values for the coins to watch are the CoinMarketCap id's, which can be seen in the url when viewing details on the site.
An example for Bitcoin; 'coinmarketcap.com/currencies/__bitcoin__/'

## Authors

* **Shaun Morrow** - *Initial work* - [shaunmza](https://github.com/shaunmza)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
