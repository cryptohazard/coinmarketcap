package data

import "time"

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
