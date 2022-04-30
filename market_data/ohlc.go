package market_data

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/quantfamily/lemonmarkets/client"
)

/*
GetOHLCQuery query used to filter OHLC(Open, High, Low, Close) values from LemonMarkets
Read more at: https://docs.lemon.markets/market-data/historical-data#get-ohlcx1
*/
type GetOHLCQuery struct {
	ISIN    []string  `url:"isin,omitempty"`
	MIC     string    `url:"mic,omitempty"`
	From    time.Time `url:"from,omitempty" layout:"2006-01-02T15:04:05Z07:00"`
	To      time.Time `url:"to,omitempty" layout:"2006-01-02T15:04:05Z07:00"`
	Sorting string    `url:"sorting,omitempty"`
	Limit   int       `url:"limit,omitempty"`
	Page    int       `url:"page,omitempty"`
}

/*
OHLC (Open, High, Low, Closed) containing information regarding how a instrument preformed during a period of time
*/
type OHLC struct {
	ISIN   string    `json:"isin"`
	Open   float64   `json:"o"`
	High   float64   `json:"h"`
	Low    float64   `json:"l"`
	Close  float64   `json:"c"`
	Volume int       `json:"v"`
	Time   time.Time `json:"t"`
	Mic    string    `json:"mic"`
}

/*
GetOHLCPerMinute returns a response containing a list of OHLC per minute
*/
func GetOHLCPerMinute(client *client.Client, query *GetOHLCQuery) <-chan Item[OHLC, error] {
	ch := make(chan Item[OHLC, error])
	go returnOHLC(client, "m1", query, ch)
	return ch
}

/*
GetOHLCPerHour returns a response containing a list of OHLC per hour
*/
func GetOHLCPerHour(client *client.Client, query *GetOHLCQuery) <-chan Item[OHLC, error] {
	ch := make(chan Item[OHLC, error])
	go returnOHLC(client, "h1", query, ch)
	return ch
}

/*
GetOHLCPerDay returns a response containing a list of OHLC per day
*/
func GetOHLCPerDay(client *client.Client, query *GetOHLCQuery) <-chan Item[OHLC, error] {
	ch := make(chan Item[OHLC, error])
	go returnOHLC(client, "d1", query, ch)
	return ch
}

func returnOHLC(client *client.Client, interval string, query *GetOHLCQuery, ch chan<- Item[OHLC, error]) {
	defer close(ch)
	response, err := client.Do("GET", fmt.Sprintf("ohlc/%s", interval), query, nil)
	if err != nil {
		ohlc := Item[OHLC, error]{}
		ohlc.Error = err
		ch <- ohlc
		return
	}
	for {
		var ohlcs []OHLC
		ohlc := Item[OHLC, error]{}
		ohlc.Error = json.Unmarshal(response.Results, &ohlcs)
		if ohlc.Error != nil {
			ch <- ohlc
			return
		}
		for _, ohlc := range ohlcs {
			ch <- Item[OHLC, error]{ohlc, nil}
		}
		if response.Next == "" {
			return
		}
		response, ohlc.Error = client.Do("GET", response.Next, nil, nil)
		if ohlc.Error != nil {
			ch <- ohlc
			return
		}
	}
}
