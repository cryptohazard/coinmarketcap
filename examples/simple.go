package main

import (
	"fmt"

	"github.com/shaunmza/coinmarketcap"
	"github.com/shaunmza/coinmarketcap/data"
)

func main() {

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

}
