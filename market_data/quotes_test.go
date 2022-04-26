package market_data

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/quantfamily/lemonmarkets/common"
	"github.com/quantfamily/lemonmarkets/common/helpers"
)

func TestGetQuotes(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_quotes.json")
	expectedResponse := new(common.Response)
	expectedQuotes := new([]Quote)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedQuotes); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		ch, err := GetQuotes(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		quotes := make([]Quote, 0)
		for quote := range ch {
			quotes = append(quotes, quote)
		}

		if !reflect.DeepEqual(expectedQuotes, expectedQuotes) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error getting account"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnError = lemonErr
		_, err := GetQuotes(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}

func TestGetQuotesIntegration(t *testing.T) {
	client := NewClient(helpers.APIKey(t))

	isins := []string{"DE000CBK1001"}
	query := GetQuotesQuery{ISIN: isins}

	_, err := GetQuotes(client, &query)
	if err != nil {
		t.Errorf(err.Error())
	}
}
