package market_data

import (
	"encoding/json"
	"time"

	"github.com/quantfamily/lemonmarkets/common"
)

/*
GetQuotesQuery is a query used to filter quotes that we can receive from LemonMarkets
Read more at: https://docs.lemon.markets/market-data/historical-data#quotes
*/
type GetQuotesQuery struct {
	ISIN    []string  `url:"isin,omitempty"`
	MIC     string    `url:"mic,omitempty"`
	From    time.Time `url:"from,omitempty"`
	To      time.Time `url:"to,omitempty"`
	Sorting string    `url:"sorting,omitempty"`
	Limit   int       `url:"limit,omitempty"`
	Page    int       `url:"page,omitempty"`
}

/*
Quote contains quote data for a specific asset known by its ISIN
*/
type Quote struct {
	ISIN      string    `json:"isin"`
	BidVolume int       `json:"b_v"`
	AskVolume int       `json:"a_v"`
	Bid       float64   `json:"bid"`
	Ask       float64   `json:"ask"`
	Time      time.Time `json:"t"`
	Mic       string    `json:"mic"`
}

/*
GetQuotes takes a possible query parameter and returns Response containing one or more quotes from LemonMarkets
*/
func GetQuotes(client common.Client, q *GetQuotesQuery) (<-chan Quote, error) {
	response, err := client.Do("GET", "quotes", q, nil)
	if err != nil {
		return nil, err
	}
	var quotes []Quote
	err = json.Unmarshal(response.Results, &quotes)
	if err != nil {
		return nil, err
	}
	ch := make(chan Quote)
	go returnQuotes(client, response, ch)
	return ch, nil
}

func returnQuotes(client common.Client, response *common.Response, outchan chan<- Quote) {
	defer close(outchan)
	var quotes []Quote
	for {
		err := json.Unmarshal(response.Results, &quotes)
		if err != nil {
			return
		}
		for _, quote := range quotes {
			outchan <- quote
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
