package trading

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

func TestCreateOrder(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "create_order.json")

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
		client := TradingClient{backend: &backend}
		order := client.CreateOrder(&Order{Quantity: 10})
		assert.NotNil(t, order.Error)
		assert.Equal(t, &expectedErr, order.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		order := client.CreateOrder(&Order{Quantity: 10})
		assert.NotNil(t, order.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, order.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		order := client.CreateOrder(&Order{Quantity: 10})
		assert.Nil(t, order.Error)
	})
}

func TestActivateOrder(t *testing.T) {
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
		client := TradingClient{backend: &backend}
		err := client.ActivateOrder("22")
		assert.NotNil(t, err)
		assert.Equal(t, &expectedErr, err)
	})
	t.Run("successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{"status": "ok"}`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		err := client.ActivateOrder("22")
		assert.Nil(t, err)
	})
}

func TestGetOrders(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_orders.json")

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
		client := TradingClient{backend: &backend}
		orderCh := client.GetOrders(nil)
		order := <-orderCh
		assert.NotNil(t, order.Error)
		assert.Equal(t, &expectedErr, order.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		orderCh := client.GetOrders(nil)
		order := <-orderCh
		assert.NotNil(t, order.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, order.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		orderCh := client.GetOrders(nil)
		order := <-orderCh
		assert.Nil(t, order.Error)
		assert.Equal(t, 2965000, order.Data.ExecutedPrice)
	})
}

func TestGetOrder(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_order.json")

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
		client := TradingClient{backend: &backend}
		order := client.GetOrder("22")
		assert.NotNil(t, order.Error)
		assert.Equal(t, &expectedErr, order.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		order := client.GetOrder("22")
		assert.NotNil(t, order.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, order.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		order := client.GetOrder("22")
		assert.Nil(t, order.Error)
		assert.Equal(t, 2965000, order.Data.ExecutedPrice)
	})
}

func TestDeleteOrder(t *testing.T) {
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
		client := TradingClient{backend: &backend}
		err := client.DeleteOrder("22")
		assert.NotNil(t, err)
		assert.Equal(t, &expectedErr, err)
	})
	t.Run("successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{"status": "ok"}`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		err := client.DeleteOrder("22")
		assert.Nil(t, err)
	})
}
