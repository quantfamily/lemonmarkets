package market_data

import (
	"encoding/json"

	"github.com/quantfamily/lemonmarkets/client"
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
GetInstruments calls backend with a optional query to filter data
Response will be list of one or more instruments that we received from LemonMarkets
*/
func GetInstruments(client *client.Client, query *GetInstrumentsQuery) <-chan Item[Instrument, error] {
	ch := make(chan Item[Instrument, error])
	go returnInstruments(client, query, ch)
	return ch
}

func returnInstruments(client *client.Client, query *GetInstrumentsQuery, ch chan<- Item[Instrument, error]) {
	defer close(ch)
	response, err := client.Do("GET", "instruments", query, nil)
	if err != nil {
		instrument := Item[Instrument, error]{}
		instrument.Error = err
		ch <- instrument
		return
	}
	for {
		var instruments []Instrument
		instrument := Item[Instrument, error]{}
		instrument.Error = json.Unmarshal(response.Results, &instruments)
		if instrument.Error != nil {
			ch <- instrument
			return
		}
		for _, instrument := range instruments {
			ch <- Item[Instrument, error]{instrument, nil}
		}
		if response.Next == "" {
			return
		}
		response, instrument.Error = client.Do("GET", response.Next, nil, nil)
		if instrument.Error != nil {
			ch <- instrument
			return
		}
	}
}
