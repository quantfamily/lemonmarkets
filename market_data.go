package main

import (
	"encoding/json"
	"time"
)

type GetInstrumentsQuery struct {
	ISIN     []string `url:",isin,omitempty"`
	MIC      string   `url:"mic,omitempty"`
	Search   string   `url:"search,omitempty"`
	Type     string   `url:"type,omitempty"`
	Currency string   `url:"currency,omitempty"`
	Limit    int      `url:"limit,omitempty"`
	Page     int      `url:"page,omitempty"`
}

type GetInstrumentsResponse struct {
	ListReply
	Results []Instrument `json:"results"`
}

type Instrument struct {
	ISIN   string  `json:"isin"`
	WKN    string  `json:"wkn"`
	Name   string  `json:"name"`
	Title  string  `json:"title"`
	Symbol string  `json:"symbol"`
	Type   string  `json:"type"`
	Venues []Venue `json:"venues"`
}

type Venue struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Mic      string `json:"mic"`
	IsOpen   bool   `json:"is_open"`
	Tradable bool   `json:"tradable"`
	Currency string `json:"currency"`
}

func GetInstruments(client Client, query *GetInstrumentsQuery) (*GetInstrumentsResponse, error) {
	responseData, err := client.Do("GET", "instruments", nil, nil)
	if err != nil {
		return nil, err
	}
	getInstruments := new(GetInstrumentsResponse)
	err = json.Unmarshal(responseData, getInstruments)
	return getInstruments, err
}

type GetQuotesQuery struct {
	ISIN    []string  `url:",isin,omitempty"`
	MIC     string    `url:"mic,omitempty"`
	From    time.Time `url:"from,omitempty"`
	To      time.Time `url:"to,omitempty"`
	Sorting string    `url:"sorting,omitempty"`
	Limit   int       `url:"limit,omitempty"`
	Page    int       `url:"page,omitempty"`
}

type Quote struct {
	ISIN      string    `json:"isin"`
	BidVolume int       `json:"b_v"`
	AskVolume int       `json:"a_v"`
	Bid       float64   `json:"bid"`
	Ask       float64   `json:"ask"`
	Time      time.Time `json:"t"`
	Mic       string    `json:"mic"`
}

type GetQuotesResponse struct {
	ListReply
	Results []Quote `json:"results"`
}

func GetQuotes(client Client, q *GetQuotesQuery) (*GetQuotesResponse, error) {
	responseData, err := client.Do("GET", "quotes", q, nil)
	if err != nil {
		return nil, err
	}
	getQuotesResponse := new(GetQuotesResponse)
	err = json.Unmarshal(responseData, getQuotesResponse)
	return getQuotesResponse, err
}

type GetOHLCQuery struct {
	ISIN    []string  `url:",isin,omitempty"`
	MIC     string    `url:"mic,omitempty"`
	From    time.Time `url:"from,omitempty"`
	To      time.Time `url:"to,omitempty"`
	Sorting string    `url:"sorting,omitempty"`
	Limit   int       `url:"limit,omitempty"`
	Page    int       `url:"page,omitempty"`
}

type OHLC struct {
	ISIN  string    `json:"isin"`
	Open  float32   `json:"open"`
	High  float32   `json:"high"`
	Low   float32   `json:"low"`
	Close float32   `json:"close"`
	Time  time.Time `json:"t"`
	Mic   string    `json:"mic"`
}

type GetOHLCResponse struct {
	ListReply
	Results []OHLC `json:"results"`
}

func GetOHLCPerMinute(client Client, query *GetOHLCQuery) (*GetOHLCResponse, error) {
	responseData, err := client.Do("GET", "ohlc/m1", query, nil)
	if err != nil {
		return nil, err
	}
	getOHLCresponse := new(GetOHLCResponse)
	err = json.Unmarshal(responseData, getOHLCresponse)
	return getOHLCresponse, err
}

func GetOHLCPerHour(client Client, query *GetOHLCQuery) (*GetOHLCResponse, error) {
	responseData, err := client.Do("GET", "ohlc/h1", query, nil)
	if err != nil {
		return nil, err
	}
	getOHLCresponse := new(GetOHLCResponse)
	err = json.Unmarshal(responseData, getOHLCresponse)
	return getOHLCresponse, err
}

func GetOHLCPerDay(client Client, query *GetOHLCQuery) (*GetOHLCResponse, error) {
	responseData, err := client.Do("GET", "ohlc/d1", query, nil)
	if err != nil {
		return nil, err
	}
	getOHLCresponse := new(GetOHLCResponse)
	err = json.Unmarshal(responseData, getOHLCresponse)
	return getOHLCresponse, err
}

type GetTradesQuery struct {
	ISIN    []string  `url:",isin,omitempty"`
	MIC     string    `url:"mic,omitempty"`
	From    time.Time `url:"from,omitempty"`
	To      time.Time `url:"to,omitempty"`
	Sorting string    `json:"sorting,omitempty"`
	Limit   int       `url:"limit,omitempty"`
	Page    int       `url:"page,omitempty"`
}

type Trade struct {
	ISIN   string    `json:"isin"`
	Price  float32   `json:"p"`
	Volume int       `json:"v"`
	Time   time.Time `json:"t"`
	Mic    string    `json:"mic"`
}

type GetTradesResponse struct {
	ListReply
	Results []Trade `json:"results"`
}

func GetTrades(client Client, query *GetTradesQuery) (*GetTradesResponse, error) {
	responseData, err := client.Do("GET", "trades", query, nil)
	if err != nil {
		return nil, err
	}
	getTradesResponse := new(GetTradesResponse)
	err = json.Unmarshal(responseData, getTradesResponse)
	return getTradesResponse, err
}
