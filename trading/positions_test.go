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

func TestGetPositions(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_positions.json")

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
		positionCh := client.GetPositions()
		position := <-positionCh
		assert.NotNil(t, position.Error)
		assert.Equal(t, &expectedErr, position.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		positionCh := client.GetPositions()
		position := <-positionCh
		assert.NotNil(t, position.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, position.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		positionCh := client.GetPositions()
		position := <-positionCh
		assert.Nil(t, position.Error)
		assert.Equal(t, 5800000, position.Data.EstimatedPriceTotal)
	})
}

func TestGetStatements(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_statements.json")

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
		statementCh := client.GetStatements()
		statement := <-statementCh
		assert.NotNil(t, statement.Error)
		assert.Equal(t, &expectedErr, statement.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		statementCh := client.GetStatements()
		statement := <-statementCh
		assert.NotNil(t, statement.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, statement.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		statementCh := client.GetStatements()
		statement := <-statementCh
		assert.Nil(t, statement.Error)
		assert.Equal(t, "US19260Q1076", statement.Data.ISIN)
	})
}

func TestGetPositionsIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ch := client.GetPositions()

	position := <-ch
	assert.Nil(t, position.Error)
}

func TestGetStatementsIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ch := client.GetStatements()

	statement := <-ch
	assert.Nil(t, statement.Error)
}
