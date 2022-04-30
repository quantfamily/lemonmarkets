package market_data

import "github.com/quantfamily/lemonmarkets/client"

type DataTypes interface {
	Instrument | OHLC | Quote | Trade
}

type Item[data DataTypes, err error] struct {
	Data  data
	Error err
}

type Environment string

const DATA string = "https://data.lemon.markets/v1"

func NewClient(APIKey string) *client.Client {
	return &client.Client{BaseURL: string(DATA), APIKey: APIKey}
}
