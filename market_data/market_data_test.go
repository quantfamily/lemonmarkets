package market_data

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/quantfamily/lemonmarkets/common"
	"github.com/quantfamily/lemonmarkets/common/helpers"
)

func TestGetInstruments(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_instruments.json")
	expectedResponse := new(common.Response)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		clientResponse, err := GetInstruments(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(clientResponse, expectedResponse) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error getting account"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnResponse = nil
		client.ReturnError = lemonErr
		_, err := GetInstruments(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}

func TestGetQuotes(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_quotes.json")
	expectedResponse := new(common.Response)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		clientResponse, err := GetQuotes(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(clientResponse, expectedResponse) {
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

func TestGetOHLCPerMinute(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_ohlc.json")
	expectedResponse := new(common.Response)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		clientResponse, err := GetOHLCPerMinute(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(clientResponse, expectedResponse) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error getting account"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnError = lemonErr
		_, err := GetOHLCPerMinute(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}

func TestGetOHLCPerHour(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_ohlc.json")
	expectedResponse := new(common.Response)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		clientResponse, err := GetOHLCPerHour(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(clientResponse, expectedResponse) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error getting account"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnError = lemonErr
		_, err := GetOHLCPerHour(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}

func TestGetOHLCPerDay(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_ohlc.json")
	expectedResponse := new(common.Response)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		clientResponse, err := GetOHLCPerDay(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(clientResponse, expectedResponse) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error getting account"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnError = lemonErr
		_, err := GetOHLCPerDay(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}

func TestGetTrades(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_trades.json")
	expectedResponse := new(common.Response)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		clientResponse, err := GetTrades(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(clientResponse, expectedResponse) {
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

/*
/ Integration tests below
*/

func TestGetInstrumentsIntegration(t *testing.T) {
	client := NewClient(helpers.APIKey(t))

	_, err := GetInstruments(client, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
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

func TestGetOHLCIntegration(t *testing.T) {
	client := NewClient(helpers.APIKey(t))
	isins := []string{"DE000CBK1001"}

	t.Run("Minute", func(t *testing.T) {
		to := time.Now().AddDate(0, 0, 0)
		from := time.Now().Add(time.Duration(-time.Hour))
		query := GetOHLCQuery{ISIN: isins, From: from, To: to}
		_, err := GetOHLCPerMinute(client, &query)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("Hour", func(t *testing.T) {
		to := time.Now().AddDate(0, 0, 0)
		from := time.Now().Add(time.Duration(-time.Hour))
		query := GetOHLCQuery{ISIN: isins, From: from, To: to}
		_, err := GetOHLCPerHour(client, &query)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("Day", func(t *testing.T) {
		from := time.Now().AddDate(0, -1, 0)
		to := time.Now()
		query := GetOHLCQuery{ISIN: isins, From: from, To: to}
		_, err := GetOHLCPerDay(client, &query)
		if err != nil {
			t.Errorf(err.Error())
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
