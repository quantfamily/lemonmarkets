package market_data

import "github.com/quantfamily/lemonmarkets/client"

type DataTypes interface {
	Instrument | OHLC | Quote | Trade | Venue
}

type Item[data DataTypes, err error] struct {
	Data  data
	Error err
}

type Environment string

const DATA string = "https://data.lemon.markets/v1"

type MarketDataClient struct {
	backend *client.Backend
}

func NewClient(APIKey string) *MarketDataClient {
	backend := client.Backend{APIKey: APIKey, BaseURL: string(DATA)}
	return &MarketDataClient{backend: &backend}
}
