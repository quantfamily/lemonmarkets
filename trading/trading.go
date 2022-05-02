package trading

import "github.com/quantfamily/lemonmarkets/client"

// DataTypes
type DataTypes interface {
	Order | Position | Account | Withdrawal | BankStatement | Document | Statement
}

// Item
type Item[data DataTypes, err error] struct {
	Data  data
	Error err
}

// Environment
type Environment string

const (
	PAPER Environment = "https://paper-trading.lemon.markets/v1"
	LIVE  Environment = "https://trading.lemon.markets/v1"
)

// TradingClient
type TradingClient struct {
	backend *client.Backend
}

// NewClient
func NewClient(APIKey string, environment Environment) *TradingClient {
	backend := client.Backend{APIKey: APIKey, BaseURL: string(environment)}
	return &TradingClient{backend: &backend}
}
