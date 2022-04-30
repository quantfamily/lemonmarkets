package trading

import "github.com/quantfamily/lemonmarkets/client"

type DataTypes interface {
	Order | Position | Account
}

type Item[data DataTypes, err error] struct {
	Data  data
	Error err
}

type Environment string

const (
	PAPER Environment = ""
	LIVE  Environment = ""
)

func NewClient(APIKey string, environment Environment) *client.Client {
	return &client.Client{BaseURL: string(environment), APIKey: APIKey}
}
