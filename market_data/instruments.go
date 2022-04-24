package market_data

import (
	"encoding/json"

	"github.com/quantfamily/lemonmarkets/common"
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
func GetInstruments(client common.Client, query *GetInstrumentsQuery) (<-chan Instrument, error) {
	response, err := client.Do("GET", "instruments", nil, nil)
	if err != nil {
		return nil, err
	}
	var instruments []Instrument
	err = json.Unmarshal(response.Results, &instruments)
	if err != nil {
		return nil, err
	}
	ch := make(chan Instrument)
	go returnInstruments(client, response, ch)
	return ch, nil
}

func returnInstruments(client common.Client, response *common.Response, outchan chan<- Instrument) {
	defer close(outchan)
	var instruments []Instrument
	for {
		err := json.Unmarshal(response.Results, &instruments)
		if err != nil {
			return
		}
		for _, instrument := range instruments {
			outchan <- instrument
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
