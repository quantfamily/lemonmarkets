package trading

import "github.com/quantfamily/lemonmarkets/common"

type Environment string

const (
	PAPER Environment = ""
	LIVE  Environment = ""
)

func NewClient(APIKey string, environment Environment) common.Client {
	return &common.LemonClient{BaseURL: string(environment), APIKey: APIKey}
}
