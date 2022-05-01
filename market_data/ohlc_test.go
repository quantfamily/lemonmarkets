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
		client := client.Client{BaseURL: server.URL}
		ohlcCh := GetOHLCPerMinute(&client, nil)
		ohlc := <-ohlcCh
		assert.NotNil(t, ohlc.Error)
		assert.Equal(t, &expectedErr, ohlc.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		client := client.Client{BaseURL: server.URL}
		ohlcCh := GetOHLCPerMinute(&client, nil)
		ohlc := <-ohlcCh
		assert.NotNil(t, ohlc.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, ohlc.Error)
	})
	t.Run("Successful test, m1", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		client := client.Client{BaseURL: server.URL}
		ohlcCh := GetOHLCPerMinute(&client, nil)
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
		client := client.Client{BaseURL: server.URL}
		ohlcCh := GetOHLCPerHour(&client, nil)
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
		client := client.Client{BaseURL: server.URL}
		ohlcCh := GetOHLCPerDay(&client, nil)
		ohlc := <-ohlcCh
		ohlc = <-ohlcCh
		assert.Nil(t, ohlc.Error)
		assert.Equal(t, 609.5, ohlc.Data.Low)
	})
}
