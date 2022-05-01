package trading

import (
	"encoding/json"
	"time"

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
		position := Item[Position, error]{}
		position.Error = err
		ch <- position
		return
	}
	for {
		var positions []Position
		position := Item[Position, error]{}
		position.Error = json.Unmarshal(response.Results, &positions)
		if position.Error != nil {
			ch <- position
			return
		}
		for _, position := range positions {
			ch <- Item[Position, error]{position, nil}
		}
		if response.Next == "" {
			return
		}
		response, position.Error = client.Do("GET", response.Next, nil, nil)
		if position.Error != nil {
			ch <- position
			return
		}
	}
}

type Statement struct {
	ID         string    `json:"id,omitempty"`
	OrderID    string    `json:"order_id,omitempty"`
	ExternalID string    `json:"external_id,omitempty"`
	Type       string    `json:"type,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
	ISIN       string    `json:"isin,omitempty"`
	ISINTitle  string    `json:"isin_title,omitempty"`
	Date       string    `json:"date,omitempty"` // TODO convert to time.Time (YYYY-MM-DD)
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func GetStatements(client *client.Client) <-chan Item[Statement, error] {
	ch := make(chan Item[Statement, error])
	go returnStatements(client, ch)
	return ch
}

func returnStatements(client *client.Client, ch chan<- Item[Statement, error]) {
	defer close(ch)
	response, err := client.Do("GET", "statements", nil, nil)
	if err != nil {
		statement := Item[Statement, error]{}
		statement.Error = err
		ch <- statement
		return
	}
	for {
		var statements []Statement
		statement := Item[Statement, error]{}
		statement.Error = json.Unmarshal(response.Results, &statements)
		if statement.Error != nil {
			ch <- statement
			return
		}
		for _, statement := range statements {
			ch <- Item[Statement, error]{statement, nil}
		}
		if response.Next == "" {
			return
		}
		response, statement.Error = client.Do("GET", response.Next, nil, nil)
		if statement.Error != nil {
			ch <- statement
			return
		}
	}
}
