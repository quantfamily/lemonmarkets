package trading

import "github.com/quantfamily/lemonmarkets/client"

type DataTypes interface {
	Order | Position | Account | Withdrawal | BankStatement | Document | Statement
}

type Item[data DataTypes, err error] struct {
	Data  data
	Error err
}

type Environment string

const (
	PAPER Environment = "https://paper-trading.lemon.markets/v1"
	LIVE  Environment = "https://trading.lemon.markets/v1"
)

type TradingClient struct {
	backend *client.Backend
}

func NewClient(APIKey string, environment Environment) *TradingClient {
	backend := client.Backend{APIKey: APIKey, BaseURL: string(environment)}
	return &TradingClient{backend: &backend}
}
