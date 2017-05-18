package main

import (
	"fmt"
	"time"

	"github.com/shaunmza/coinmarketcap"
)

func main() {

	var r coinmarketcap.Ticker
	var err error

	// Which coins do you want to watch?
	t := make([]string, 0)
	t = append(t, "bitcoin")
	t = append(t, "litecoin")
	t = append(t, "steem")
	t = append(t, "steem-dollars")

	// Endpoints are updated every 5 minutes, se we use that here
	period := 60 * 5
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

}

func printData(r coinmarketcap.Ticker) {
	// Just print it out for now
	for _, coin := range r.Coins {
		fmt.Println("Symbol: '" + coin.Symbol + "', Name: '" +
			coin.Name + "', Price Bitcoin: '" +
			coin.PriceBtc + "', Price USD: '" + coin.PriceUsd + "'")
	}
}
