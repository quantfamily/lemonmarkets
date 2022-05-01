package market_data

import (
	"encoding/json"
	"time"
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
	Bid       float64   `json:"b"`
	Ask       float64   `json:"a"`
	Time      time.Time `json:"t"`
	Mic       string    `json:"mic"`
}

/*
GetQuotes takes a possible query parameter and returns Response containing one or more quotes from LemonMarkets
*/
func (cl *MarketDataClient) GetQuotes(query *GetQuotesQuery) <-chan Item[Quote, error] {
	ch := make(chan Item[Quote, error])
	go cl.returnQuotes(query, ch)
	return ch
}

func (cl *MarketDataClient) returnQuotes(query *GetQuotesQuery, ch chan<- Item[Quote, error]) {
	defer close(ch)
	response, err := cl.backend.Do("GET", "quotes", query, nil)
	if err != nil {
		quote := Item[Quote, error]{}
		quote.Error = err
		ch <- quote
		return
	}
	for {
		var quotes []Quote
		quote := Item[Quote, error]{}
		quote.Error = json.Unmarshal(response.Results, &quotes)
		if quote.Error != nil {
			ch <- quote
			return
		}
		for _, quote := range quotes {
			ch <- Item[Quote, error]{quote, nil}
		}
		if response.Next == "" {
			return
		}
		response, quote.Error = cl.backend.Do("GET", response.Next, nil, nil)
		if quote.Error != nil {
			ch <- quote
			return
		}
	}
}
