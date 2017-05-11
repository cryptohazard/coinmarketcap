package coinmarketcap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/shaunmza/coinmarketcap/data"
)

// As per API documentation found at `https://coinmarketcap.com/api/`
const tickerURL = "https://api.coinmarketcap.com/v1/ticker/"

// Channel into which fetched data is sent
var coinCh chan data.Ticker

// Init is called so we can return the channel.
// This allows messages to be read off of it,
// and we can close it from outside too
func Init() chan data.Ticker {
	coinCh = make(chan data.Ticker)

	return coinCh
}

// WatchCoins runs a ticker to fetch the data periodically
func WatchCoins(coins []string, period int) {
	// Don't call too often
	if period < 10 {
		period = 10
	}

	// Call now so we don't have to wait for the first batch of data
	go func() {
		s := getData(coins)
		coinCh <- s
	}()

	ticker := time.NewTicker(time.Second * time.Duration(period))
	go func() {
		for _ = range ticker.C {
			s := getData(coins)
			coinCh <- s
		}
	}()
}

// getData makes the actuall http request, parses the JSON and returns
// the data in a struct
func getData(coins []string) data.Ticker {

	resp, err := http.Get(tickerURL)
	if err != nil {
		return data.Ticker{Error: err}
	}

	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return data.Ticker{Error: err}
	}

	var j []*data.Coin
	var res []*data.Coin
	if err := json.Unmarshal(ret, &j); err != nil {
		log.Fatal(err)
		return data.Ticker{Error: err}
	}

	for _, c1 := range j {
		for _, c2 := range coins {
			if c1.ID == c2 {
				res = append(res, c1)
			}
		}

	}

	r := data.Ticker{Coins: res, LastUpdate: time.Now()}
	return r
}
