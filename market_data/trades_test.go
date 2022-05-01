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

func TestGetTrades(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_trades.json")

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
		client := client.Client{BaseURL: server.URL}
		tradeCh := GetTrades(&client, nil)
		trade := <-tradeCh
		assert.NotNil(t, trade.Error)
		assert.Equal(t, &expectedErr, trade.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		client := client.Client{BaseURL: server.URL}
		tradeCh := GetTrades(&client, nil)
		trade := <-tradeCh
		assert.NotNil(t, trade.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, trade.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		client := client.Client{BaseURL: server.URL}
		tradeCh := GetTrades(&client, nil)
		trade := <-tradeCh
		assert.Nil(t, trade.Error)
		assert.Equal(t, 2, trade.Data.Volume)
	})
}
