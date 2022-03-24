package lemonmarkets

import (
	"encoding/json"
	"time"
)

/*
GetInstrumentsQuery is query to backend for /instruments
Documentation: https://docs.lemon.markets/market-data/instruments-tradingvenues#get-instruments
*/
type GetInstrumentsQuery struct {
	ISIN     []string `url:"isin,omitempty"`
	MIC      string   `url:"mic,omitempty"`
	Search   string   `url:"search,omitempty"`
	Type     string   `url:"type,omitempty"`
	Currency string   `url:"currency,omitempty"`
	Limit    int      `url:"limit,omitempty"`
	Page     int      `url:"page,omitempty"`
}

/*
GetInstrumentsResponse is response data from call to /instruments
Documentation: https://docs.lemon.markets/market-data/instruments-tradingvenues#get-instruments
*/
type GetInstrumentsResponse struct {
	ListReply
	Results []Instrument `json:"results"`
}

/*
Instrument is possibly tradable asset that can be ordered
*/
type Instrument struct {
	ISIN   string  `json:"isin"`
	WKN    string  `json:"wkn"`
	Name   string  `json:"name"`
	Title  string  `json:"title"`
	Symbol string  `json:"symbol"`
	Type   string  `json:"type"`
	Venues []Venue `json:"venues"`
}

/*
Venue of where the tradeable instrument is located
*/
type Venue struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Mic      string `json:"mic"`
	IsOpen   bool   `json:"is_open"`
	Tradable bool   `json:"tradable"`
	Currency string `json:"currency"`
}

/*
GetInstruments calls backend with a optional query to filter data
Response will be list of one or more instruments that we received from LemonMarkets
*/
func GetInstruments(client Client, query *GetInstrumentsQuery) (*GetInstrumentsResponse, error) {
	responseData, err := client.Do("GET", "instruments", nil, nil)
	if err != nil {
		return nil, err
	}
	getInstruments := new(GetInstrumentsResponse)
	err = json.Unmarshal(responseData, getInstruments)
	return getInstruments, err
}

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
GetQuotesResponse is response from Lemonmarkets containing quotes
Read more at: https://docs.lemon.markets/market-data/historical-data#quotes
*/
type GetQuotesResponse struct {
	ListReply
	Results []Quote `json:"results"`
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
func GetQuotes(client Client, q *GetQuotesQuery) (*GetQuotesResponse, error) {
	responseData, err := client.Do("GET", "quotes", q, nil)
	if err != nil {
		return nil, err
	}
	getQuotesResponse := new(GetQuotesResponse)
	err = json.Unmarshal(responseData, getQuotesResponse)
	return getQuotesResponse, err
}

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
GetOHLCResponse response coming from LemonMarkets containing one or more results as OHLC list
Read more at: https://docs.lemon.markets/market-data/historical-data#get-ohlcx1
*/
type GetOHLCResponse struct {
	ListReply
	Results []OHLC `json:"results"`
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
func GetOHLCPerMinute(client Client, query *GetOHLCQuery) (*GetOHLCResponse, error) {
	responseData, err := client.Do("GET", "ohlc/m1", query, nil)
	if err != nil {
		return nil, err
	}
	getOHLCresponse := new(GetOHLCResponse)
	err = json.Unmarshal(responseData, getOHLCresponse)
	return getOHLCresponse, err
}

/*
GetOHLCPerHour returns a response containing a list of OHLC per hour
*/
func GetOHLCPerHour(client Client, query *GetOHLCQuery) (*GetOHLCResponse, error) {
	responseData, err := client.Do("GET", "ohlc/h1", query, nil)
	if err != nil {
		return nil, err
	}
	getOHLCresponse := new(GetOHLCResponse)
	err = json.Unmarshal(responseData, getOHLCresponse)
	return getOHLCresponse, err
}

/*
GetOHLCPerDay returns a response containing a list of OHLC per day
*/
func GetOHLCPerDay(client Client, query *GetOHLCQuery) (*GetOHLCResponse, error) {
	responseData, err := client.Do("GET", "ohlc/d1", query, nil)
	if err != nil {
		return nil, err
	}
	getOHLCresponse := new(GetOHLCResponse)
	err = json.Unmarshal(responseData, getOHLCresponse)
	return getOHLCresponse, err
}

/*
GetTradesQuery query used to filter the result of trades
Read more at: https://docs.lemon.markets/market-data/historical-data#get-trades
*/
type GetTradesQuery struct {
	ISIN    []string  `url:"isin,omitempty"`
	MIC     string    `url:"mic,omitempty"`
	From    time.Time `url:"from,omitempty"`
	To      time.Time `url:"to,omitempty"`
	Sorting string    `json:"sorting,omitempty"`
	Limit   int       `url:"limit,omitempty"`
	Page    int       `url:"page,omitempty"`
}

/*
GetTradesResponse is a response containing a list of one or more trades from LemonMarkets
Read more at: https://docs.lemon.markets/market-data/historical-data#get-trades
*/
type GetTradesResponse struct {
	ListReply
	Results []Trade `json:"results"`
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
func GetTrades(client Client, query *GetTradesQuery) (*GetTradesResponse, error) {
	responseData, err := client.Do("GET", "trades", query, nil)
	if err != nil {
		return nil, err
	}
	getTradesResponse := new(GetTradesResponse)
	err = json.Unmarshal(responseData, getTradesResponse)
	return getTradesResponse, err
}
