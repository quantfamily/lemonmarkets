package market_data

import (
	"encoding/json"
	"time"

	"github.com/quantfamily/lemonmarkets/client"
)

/*
GetTradesQuery query used to filter the result of trades
Read more at: https://docs.lemon.markets/market-data/historical-data#get-trades
*/
type TradesQuery struct {
	ISIN    []string  `url:"isin,omitempty"`
	MIC     string    `url:"mic,omitempty"`
	From    time.Time `url:"from,omitempty"`
	To      time.Time `url:"to,omitempty"`
	Sorting string    `json:"sorting,omitempty"`
	Limit   int       `url:"limit,omitempty"`
	Page    int       `url:"page,omitempty"`
}

/*
Trade containing information about a specific trade
*/
type Trade struct {
	ISIN   string    `json:"isin"`
	Price  float32   `json:"p"`
	Volume int       `json:"v"`
	Time   time.Time `json:"t"`
	Mic    string    `json:"mic"`
}

/*
GetTrades take a possible query parameter and returns a object contaning one or mote trades
*/
func GetTrades(client *client.Client, query *TradesQuery) <-chan Item[Trade, error] {
	ch := make(chan Item[Trade, error])
	go returnTrades(client, query, ch)
	return ch
}

func returnTrades(client *client.Client, query *TradesQuery, ch chan<- Item[Trade, error]) {
	defer close(ch)
	response, err := client.Do("GET", "trades", query, nil)
	if err != nil {
		trade := Item[Trade, error]{}
		trade.Error = err
		ch <- trade
		return
	}
	for {
		var trades []Trade
		trade := Item[Trade, error]{}
		trade.Error = json.Unmarshal(response.Results, &trades)
		if trade.Error != nil {
			ch <- trade
			return
		}
		for _, trade := range trades {
			ch <- Item[Trade, error]{trade, nil}
		}
		if response.Next == "" {
			return
		}
		response, trade.Error = client.Do("GET", response.Next, nil, nil)
		if trade.Error != nil {
			ch <- trade
			return
		}
	}
}
