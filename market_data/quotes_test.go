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

func TestGetQuotes(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_quotes.json")

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
		quoteCh := client.GetQuotes(nil)
		quote := <-quoteCh
		assert.NotNil(t, quote.Error)
		assert.Equal(t, &expectedErr, quote.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		quoteCh := client.GetQuotes(nil)
		quote := <-quoteCh
		assert.NotNil(t, quote.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, quote.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := MarketDataClient{backend: &backend}
		quoteCh := client.GetQuotes(nil)
		quote := <-quoteCh
		assert.Nil(t, quote.Error)
		assert.Equal(t, 921.1, quote.Data.Ask)
	})
}

func TestGetQuotesIntegration(t *testing.T) {
	client := IntegrationClient(t)
	quotesq := GetQuotesQuery{ISIN: []string{"SE0000115446"}}
	ch := client.GetQuotes(&quotesq)

	quote := <-ch
	assert.Nil(t, quote.Error)
}
