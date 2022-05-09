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

func TestGetOHLCs(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_ohlc.json")

	t.Run("fail to get response", func(t *testing.T) {
		expectedErr := client.LemonError{
			Time:    time.Time{},
			Mode:    "paper",
			Status:  "error",
			Code:    "order_total_price_limit_exceeded",
			Message: "cannot place/activate buy order if estimated total price is greater than 25k Euro",
		}
		errRsp, _ := json.Marshal(&expectedErr)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, string(errRsp), 400)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		ohlcCh := client.GetOHLCPerMinute(nil)
		ohlc := <-ohlcCh
		assert.NotNil(t, ohlc.Error)
		assert.Equal(t, &expectedErr, ohlc.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		ohlcCh := client.GetOHLCPerMinute(nil)
		ohlc := <-ohlcCh
		assert.NotNil(t, ohlc.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, ohlc.Error)
	})
	t.Run("Successful test, m1", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		ohlcCh := client.GetOHLCPerMinute(nil)
		ohlc := <-ohlcCh
		ohlc = <-ohlcCh
		assert.Nil(t, ohlc.Error)
		assert.Equal(t, 609.5, ohlc.Data.Low)
	})
	t.Run("Successful test, h1", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		ohlcCh := client.GetOHLCPerHour(nil)
		ohlc := <-ohlcCh
		ohlc = <-ohlcCh
		assert.Nil(t, ohlc.Error)
		assert.Equal(t, 609.5, ohlc.Data.Low)
	})
	t.Run("Successful test, d1", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		ohlcCh := client.GetOHLCPerDay(nil)
		ohlc := <-ohlcCh
		ohlc = <-ohlcCh
		assert.Nil(t, ohlc.Error)
		assert.Equal(t, 609.5, ohlc.Data.Low)
	})
}

func TestGetOHLCPerMinuteIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ohlcQ := GetOHLCQuery{ISIN: []string{"SE0000115446"}}
	ch := client.GetOHLCPerMinute(&ohlcQ)

	ohlc := <-ch
	assert.Nil(t, ohlc.Error)
}

func TestGetOHLCPerHourIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ohlcQ := GetOHLCQuery{ISIN: []string{"SE0000115446"}}
	ch := client.GetOHLCPerHour(&ohlcQ)

	ohlc := <-ch
	assert.Nil(t, ohlc.Error)
}

func TestGetOHLCPerDayIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ohlcQ := GetOHLCQuery{ISIN: []string{"SE0000115446"}}
	ch := client.GetOHLCPerDay(&ohlcQ)

	ohlc := <-ch
	assert.Nil(t, ohlc.Error)
}

func TestMyWay(t *testing.T) {
	from, _ := time.Parse(time.RFC3339, "2022-04-01T00:00:00Z")
	to, _ := time.Parse(time.RFC3339, "2022-05-08T00:00:00Z")

	client := IntegrationClient(t)
	ohlcQ := GetOHLCQuery{ISIN: []string{"DE0005190003", "DE0005439004", "DE0006062144", "DE0008404005", "DE000A1DAHH0", "DE000A1EWWW0", "DE000BASF111", "DE000BAY0017", "DE000DTR0CK8", "NL0000235190"}, From: from, To: to}
	ch := client.GetOHLCPerDay(&ohlcQ)

	for ohlc := range ch {
		assert.Nil(t, ohlc.Error)
	}
}
