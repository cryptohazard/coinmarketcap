package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// The Ticker struct holds a map of Coin structs, the last time
// data was fetched successfully and in the case of an error, the error message
type Ticker struct {
	Coins      map[string]*Coin
	LastUpdate time.Time
	Error      error
}

// The Coin struct holds the data pertaining to a specific coin
type Coin struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Rank             int16   `json:"rank,string"`
	PriceUSD         float64 `json:"price_usd,string"`
	PriceBTC         float64 `json:"price_btc,string"`
	Volume24USD      float64 `json:"24h_volume_usd,string"`
	MarketCapUSD     float64 `json:"market_cap_usd,string"`
	AvailableSupply  float64 `json:"available_supply,string"`
	TotalSupply      float64 `json:"total_supply,string"`
	PercentChange1h  float64 `json:"percent_change_1h,string"`
	PercentChange24h float64 `json:"percent_change_24h,string"`
	PercentChange7d  float64 `json:"percent_change_7d,string"`
	LastUpdated      string  `json:"last_updated"`
	PriceEUR         float64 `json:"price_eur,string"`
	Volume24EUR      float64 `json:"24h_volume_eur,string"`
	MarketCapEUR     float64 `json:"market_cap_eur,string"`
}

// As per API documentation found at `https://coinmarketcap.com/api/`
const tickerURL = "https://api.coinmarketcap.com/v1/ticker/?convert=EUR&limit=0"

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

	if errr := json.Unmarshal(ret, &j); errr != nil {
		log.Fatal(errr)
		return Ticker{}, errr
	}

	var res = make(map[string]*Coin, len(j))
	for _, c1 := range j {
		for _, c2 := range coins {
			if c1.ID == c2 {
				res[c1.ID] = c1
			}
		}
	}

	r := Ticker{Coins: res, LastUpdate: time.Now()}
	return r, err
}

//Quick convenient function to print the price information
func PrintData(ticker Ticker) {
	// Just print it out for now
	for _, coin := range ticker.Coins {
		fmt.Println("Symbol: '"+coin.Symbol+"', Name: '"+
			coin.Name+"', Price Bitcoin: '",
			coin.PriceBTC, "', Price EUR: '", coin.PriceEUR, "'")
	}
}
