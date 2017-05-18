package coinmarketcap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// The Ticker struct holds a map of Coin structs, the last time
// data was fetched successfully and in the case of an error, the error message
type Ticker struct {
	Coins      []*Coin
	LastUpdate time.Time
	Error      error
}

// The Coin struct holds the data pertaining to a specific coin
type Coin struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             int16  `json:"rank,string"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Usd24hVolume     string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h, string"`
	PercentChange7d  string `json:"percent_change_7d, string"`
	LastUpdated      string `json:"last_updated"`
}

// As per API documentation found at `https://coinmarketcap.com/api/`
const tickerURL = "https://api.coinmarketcap.com/v1/ticker/"

// GetData makes the actual http request, parses the JSON and returns
// the data in a struct
func GetData(coins []string) (Ticker, error) {

	resp, err := http.Get(tickerURL)
	if err != nil {
		return Ticker{}, err
	}

	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return Ticker{}, err
	}

	var j []*Coin
	var res []*Coin
	if err := json.Unmarshal(ret, &j); err != nil {
		log.Fatal(err)
		return Ticker{}, err
	}

	for _, c1 := range j {
		for _, c2 := range coins {
			if c1.ID == c2 {
				res = append(res, c1)
			}
		}

	}

	r := Ticker{Coins: res, LastUpdate: time.Now()}
	return r, err
}
