package market_data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/quantfamily/lemonmarkets/client"
	"github.com/quantfamily/lemonmarkets/client/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGetInstruments(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_instruments.json")

	t.Run("fail to get response", func(t *testing.T) {
		expectedErr := client.LemonError{
			Time:    time.Time{},
			Mode:    "paper",
			Status:  "error",
			Code:    "instrument_total_price_limit_exceeded",
			Message: "cannot place/activate buy instrument if estimated total price is greater than 25k Euro",
		}
		errRsp, _ := json.Marshal(&expectedErr)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, string(errRsp), 400)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		instrumentCh := client.GetInstruments(nil)
		instrument := <-instrumentCh
		assert.NotNil(t, instrument.Error)
		assert.Equal(t, &expectedErr, instrument.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		instrumentCh := client.GetInstruments(nil)
		instrument := <-instrumentCh
		assert.NotNil(t, instrument.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, instrument.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		instrumentCh := client.GetInstruments(nil)
		instrument := <-instrumentCh
		assert.Nil(t, instrument.Error)
		assert.Equal(t, "1QZ", instrument.Data.Symbol)
	})
}

func TestGetVenues(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_venues.json")

	t.Run("fail to get response", func(t *testing.T) {
		expectedErr := client.LemonError{
			Time:    time.Time{},
			Mode:    "paper",
			Status:  "error",
			Code:    "instrument_total_price_limit_exceeded",
			Message: "cannot place/activate buy instrument if estimated total price is greater than 25k Euro",
		}
		errRsp, _ := json.Marshal(&expectedErr)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, string(errRsp), 400)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		venueCh := client.GetVenues()
		venue := <-venueCh
		assert.NotNil(t, venue.Error)
		assert.Equal(t, &expectedErr, venue.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		venueCh := client.GetVenues()
		venue := <-venueCh
		assert.NotNil(t, venue.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, venue.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		venueCh := client.GetVenues()
		venue := <-venueCh
		assert.Nil(t, venue.Error)
		assert.Equal(t, true, venue.Data.IsOpen)
	})
}

func TestGetInstrumentsIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ch := client.GetInstruments(nil)

	instrument := <-ch
	assert.Nil(t, instrument.Error)
}

func TestGetVenueIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ch := client.GetVenues()

	venue := <-ch
	assert.Nil(t, venue.Error)
}
