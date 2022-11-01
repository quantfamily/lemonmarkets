package market_data

import "github.com/quantfamily/lemonmarkets/client"

// DataTypes that we use in this package
type DataTypes interface {
	Instrument | OHLC | Quote | Trade | Venue
}

// Item returned over channels
type Item[data DataTypes, err error] struct {
	Data  data
	Error err
}

// Environment, pointing to a url that we use as backend base- url
type Environment string

// BASE_URL points to backend for market_data
const BASE_URL string = "https://data.lemon.markets/v1"

// MarketDataClient with methods that we use to fetch data
type MarketDataClient struct {
	backend *client.Backend
}

// NewClient takes APIKey and returns a MarketDataClient
func NewClient(APIKey string) *MarketDataClient {
	backend := client.Backend{APIKey: APIKey, BaseURL: BASE_URL}
	return &MarketDataClient{backend: &backend}
}
