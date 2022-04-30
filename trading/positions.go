package trading

import (
	"encoding/json"

	"github.com/quantfamily/lemonmarkets/client"
)

/*
PortfolioPosition is information about Positions inside the Portfolio
*/
type Position struct {
	ISIN                string  `json:"isin"`
	ISINTitle           string  `json:"isin_title"`
	Quantity            int     `json:"quantity"`
	BuyPriceAverage     float64 `json:"buy_price_avg"`
	EstimatedPriceTotal float64 `json:"estimated_price_total"`
	EstimatedPrice      float64 `json:"estimated_price"`
}

/*
GetPositions returns current positions in LemonMarkets
*/
func GetPositions(client *client.Client) <-chan Item[Position, error] {
	ch := make(chan Item[Position, error])
	go returnPositions(client, ch)
	return ch
}

func returnPositions(client *client.Client, ch chan<- Item[Position, error]) {
	defer close(ch)
	response, err := client.Do("GET", "positions", nil, nil)
	if err != nil {
		order := Item[Position, error]{}
		order.Error = err
		ch <- order
		return
	}
	for {
		var positions []Position
		order := Item[Position, error]{}
		order.Error = json.Unmarshal(response.Results, &positions)
		if order.Error != nil {
			ch <- order
			return
		}
		for _, position := range positions {
			ch <- Item[Position, error]{position, nil}
		}
		if response.Next == "" {
			return
		}
		response, order.Error = client.Do("GET", response.Next, nil, nil)
		if order.Error != nil {
			ch <- order
			return
		}
	}
}
