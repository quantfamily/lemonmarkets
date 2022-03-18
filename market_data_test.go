package lemonmarkets

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestGetInstruments(t *testing.T) {
	rawFileBytes := ParseFile(t, "get_instruments.json")
	expectedResponse := new(GetInstrumentsResponse)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
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
		lemonErr := LemonError{Message: errMessage}

		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		_, err := GetInstruments(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetInstruments(client, nil)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestGetQuotes(t *testing.T) {
	rawFileBytes := ParseFile(t, "get_quotes.json")
	expectedResponse := new(GetQuotesResponse)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
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
		lemonErr := LemonError{Message: errMessage}

		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		_, err := GetQuotes(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetQuotes(client, nil)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestGetOHLCPerMinute(t *testing.T) {
	rawFileBytes := ParseFile(t, "get_ohlc.json")
	expectedResponse := new(GetOHLCResponse)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
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
		lemonErr := LemonError{Message: errMessage}

		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		_, err := GetOHLCPerMinute(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetOHLCPerMinute(client, nil)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestGetOHLCPerHour(t *testing.T) {
	rawFileBytes := ParseFile(t, "get_ohlc.json")
	expectedResponse := new(GetOHLCResponse)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
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
		lemonErr := LemonError{Message: errMessage}

		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		_, err := GetOHLCPerHour(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetOHLCPerHour(client, nil)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestGetOHLCPerDay(t *testing.T) {
	rawFileBytes := ParseFile(t, "get_ohlc.json")
	expectedResponse := new(GetOHLCResponse)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
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
		lemonErr := LemonError{Message: errMessage}

		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		_, err := GetOHLCPerDay(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetOHLCPerDay(client, nil)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestGetTrades(t *testing.T) {
	rawFileBytes := ParseFile(t, "get_trades.json")
	expectedResponse := new(GetTradesResponse)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
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
		lemonErr := LemonError{Message: errMessage}

		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		_, err := GetTrades(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetTrades(client, nil)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

/*
/ Integration tests below
*/

func TestGetInstrumentsIntegration(t *testing.T) {
	client := GetClient(t, DATA)

	_, err := GetInstruments(client, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetQuotesIntegration(t *testing.T) {
	client := GetClient(t, DATA)

	isins := []string{"DE000CBK1001"}
	query := GetQuotesQuery{ISIN: isins}

	_, err := GetQuotes(client, &query)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetOHLCIntegration(t *testing.T) {
	client := GetClient(t, DATA)
	isins := []string{"DE000CBK1001"}
	from := time.Now().AddDate(0, -1, 0)
	to := time.Now()
	query := GetOHLCQuery{ISIN: isins, From: from, To: to}

	t.Run("Minute", func(t *testing.T) {
		_, err := GetOHLCPerMinute(client, &query)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("Hour", func(t *testing.T) {
		_, err := GetOHLCPerHour(client, &query)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("Day", func(t *testing.T) {
		_, err := GetOHLCPerDay(client, &query)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}

func TestGetTradesIntegration(t *testing.T) {
	client := GetClient(t, DATA)
	from := time.Now().AddDate(0, -1, 0)
	to := time.Now()
	isins := []string{"DE000CBK1001"}
	query := GetTradesQuery{ISIN: isins, From: from, To: to}
	_, err := GetTrades(client, &query)
	if err != nil {
		t.Errorf(err.Error())
	}
}
