package market_data

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/quantfamily/lemonmarkets/common"
	"github.com/quantfamily/lemonmarkets/common/helpers"
)

func TestGetTrades(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_trades.json")
	expectedResponse := new(common.Response)
	expectedTrades := new([]Trade)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedTrades); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		ch, err := GetTrades(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		trades := make([]Trade, 0)
		for trade := range ch {
			trades = append(trades, trade)
		}
		if !reflect.DeepEqual(&trades, expectedTrades) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error getting account"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnError = lemonErr
		_, err := GetTrades(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}

func TestGetTradesIntegration(t *testing.T) {
	client := NewClient(helpers.APIKey(t))
	from := time.Now().AddDate(0, -1, 0)
	to := time.Now()
	isins := []string{"DE000CBK1001"}
	query := TradesQuery{ISIN: isins, From: from, To: to}
	_, err := GetTrades(client, &query)
	if err != nil {
		t.Errorf(err.Error())
	}
}
