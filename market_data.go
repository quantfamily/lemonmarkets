package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type GetInstrumentsResult struct {
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

func GetInstruments(client Client) (*GetInstrumentsResult, error) {
	responseData, err := client.Do("GET", "instruments", nil)
	if err != nil {
		return nil, err
	}
	getInstruments := new(GetInstrumentsResult)
	err = json.Unmarshal(responseData, getInstruments)
	return getInstruments, err
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

func GetQuotes(client Client, isin ...string) (*GetQuotesResponse, error) {
	endpoints := fmt.Sprintf("quotes?isin=%s", strings.Join(isin, ","))
	responseData, err := client.Do("GET", endpoints, nil)
	if err != nil {
		return nil, err
	}
	getQuotesResponse := new(GetQuotesResponse)
	err = json.Unmarshal(responseData, getQuotesResponse)
	return getQuotesResponse, err
}

type OHLC struct {
	ISIN  string      `json:"isin"`
	Open  float32     `json:"open"`
	High  float32     `json:"high"`
	Low   float32     `json:"low"`
	Close float32     `json:"close"`
	Time  time.Ticker `json:"t"`
	Mic   string      `json:"mic"`
}

type GetOHLCResponse struct {
	ListReply
	Results []OHLC `json:"results"`
}

func GetOHLCPerMinute(client Client, from int64, to int64, isin ...string) (*GetOHLCResponse, error) {
	endpoints := fmt.Sprintf("ohlc/m1?from=%d&to=%d&isin=%s", from, to, strings.Join(isin, ","))
	responseData, err := client.Do("GET", endpoints, nil)
	if err != nil {
		return nil, err
	}
	getOHLCresponse := new(GetOHLCResponse)
	err = json.Unmarshal(responseData, getOHLCresponse)
	return getOHLCresponse, err
}

func GetOHLCPerHour(client Client, from int64, to int64, isin ...string) (*GetOHLCResponse, error) {
	endpoints := fmt.Sprintf("ohlc/h1?from=%d&to=%d&isin=%s", from, to, strings.Join(isin, ","))
	responseData, err := client.Do("GET", endpoints, nil)
	if err != nil {
		return nil, err
	}
	getOHLCresponse := new(GetOHLCResponse)
	err = json.Unmarshal(responseData, getOHLCresponse)
	return getOHLCresponse, err
}

func GetOHLCPerDay(client Client, from int64, to int64, isin ...string) (*GetOHLCResponse, error) {
	endpoints := fmt.Sprintf("ohlc/d1?from=%d&to=%d&isin=%s", from, to, strings.Join(isin, ","))
	responseData, err := client.Do("GET", endpoints, nil)
	if err != nil {
		return nil, err
	}
	getOHLCresponse := new(GetOHLCResponse)
	err = json.Unmarshal(responseData, getOHLCresponse)
	return getOHLCresponse, err
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

func GetTrades(client Client, from int64, to int64, isin ...string) (*GetTradesResponse, error) {
	endpoints := fmt.Sprintf("quotes?from=%d&to=%d&isin=%s", from, to, strings.Join(isin, ","))
	responseData, err := client.Do("GET", endpoints, nil)
	if err != nil {
		return nil, err
	}
	getTradesResponse := new(GetTradesResponse)
	err = json.Unmarshal(responseData, getTradesResponse)
	return getTradesResponse, err
}
