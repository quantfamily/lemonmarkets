package trading

import (
	"encoding/json"

	"github.com/quantfamily/lemonmarkets/common"
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
func GetPositions(client common.Client) (<-chan Position, error) {
	response, err := client.Do("GET", "portfolio", nil, nil)
	if err != nil {
		return nil, err
	}
	var positions []Position
	err = json.Unmarshal(response.Results, &positions)
	if err != nil {
		return nil, err
	}
	ch := make(chan Position)
	go returnPositions(client, response, ch)
	return ch, nil
}

func returnPositions(client common.Client, response *common.Response, outchan chan<- Position) {
	defer close(outchan)
	var positions []Position
	for {
		err := json.Unmarshal(response.Results, &positions)
		if err != nil {
			return
		}
		for _, position := range positions {
			outchan <- position
		}
		if response.Next == "" {
			return
		}
		response, err = client.Do("GET", response.Next, nil, nil)
		if err != nil {
			return
		}
	}
}
