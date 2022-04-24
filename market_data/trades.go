package market_data

import (
	"encoding/json"
	"time"

	"github.com/quantfamily/lemonmarkets/common"
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
func GetTrades(client common.Client, query *TradesQuery) (<-chan Trade, error) {
	response, err := client.Do("GET", "trades", query, nil)
	if err != nil {
		return nil, err
	}
	var trades []Trade
	err = json.Unmarshal(response.Results, &trades)
	if err != nil {
		return nil, err
	}
	ch := make(chan Trade)
	go returnTrades(client, response, ch)
	return ch, err
}

func returnTrades(client common.Client, response *common.Response, outchan chan<- Trade) {
	defer close(outchan)
	var trades []Trade
	for {
		err := json.Unmarshal(response.Results, &trades)
		if err != nil {
			return
		}
		for _, trade := range trades {
			outchan <- trade
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
