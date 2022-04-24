package market_data

import "github.com/quantfamily/lemonmarkets/common"

type Environment string

const DATA string = "https://data.lemon.markets/v1"

func NewClient(APIKey string) common.Client {
	return &common.LemonClient{BaseURL: string(DATA), APIKey: APIKey}
}
