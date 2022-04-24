package market_data

import (
	"encoding/json"
	"time"

	"github.com/quantfamily/lemonmarkets/common"
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
func GetOHLCPerMinute(client common.Client, query *GetOHLCQuery) (<-chan OHLC, error) {
	response, err := client.Do("GET", "ohlc/m1", query, nil)
	if err != nil {
		return nil, err
	}
	var ohlc []OHLC
	err = json.Unmarshal(response.Results, ohlc)
	if err != nil {
		return nil, err
	}
	ch := make(chan OHLC)
	go returnOHLC(client, response, ch)
	return ch, err
}

/*
GetOHLCPerHour returns a response containing a list of OHLC per hour
*/
func GetOHLCPerHour(client common.Client, query *GetOHLCQuery) (<-chan OHLC, error) {
	response, err := client.Do("GET", "ohlc/h1", query, nil)
	if err != nil {
		return nil, err
	}
	var ohlc []OHLC
	err = json.Unmarshal(response.Results, ohlc)
	if err != nil {
		return nil, err
	}
	ch := make(chan OHLC)
	go returnOHLC(client, response, ch)
	return ch, err
}

/*
GetOHLCPerDay returns a response containing a list of OHLC per day
*/
func GetOHLCPerDay(client common.Client, query *GetOHLCQuery) (<-chan OHLC, error) {
	response, err := client.Do("GET", "ohlc/d1", query, nil)
	if err != nil {
		return nil, err
	}
	var ohlc []OHLC
	err = json.Unmarshal(response.Results, ohlc)
	if err != nil {
		return nil, err
	}
	ch := make(chan OHLC)
	go returnOHLC(client, response, ch)
	return ch, err
}

func returnOHLC(client common.Client, response *common.Response, outchan chan<- OHLC) {
	defer close(outchan)
	var ohlcs []OHLC
	for {
		err := json.Unmarshal(response.Results, &ohlcs)
		if err != nil {
			return
		}
		for _, ohlc := range ohlcs {
			outchan <- ohlc
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
