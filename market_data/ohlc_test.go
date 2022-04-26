package market_data

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/quantfamily/lemonmarkets/common"
	"github.com/quantfamily/lemonmarkets/common/helpers"
)

func TestGetOHLCPerMinute(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_ohlc.json")
	expectedResponse := new(common.Response)
	expectedOHL := new([]OHLC)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedOHL); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		ch, err := GetOHLCPerMinute(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		ohlcs := make([]OHLC, 0)
		for ohlc := range ch {
			ohlcs = append(ohlcs, ohlc)
		}
		if !reflect.DeepEqual(&ohlcs, expectedOHL) {
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
	expectedOHL := new([]OHLC)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedOHL); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		ch, err := GetOHLCPerHour(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		ohlcs := make([]OHLC, 0)
		for ohlc := range ch {
			ohlcs = append(ohlcs, ohlc)
		}
		if !reflect.DeepEqual(&ohlcs, expectedOHL) {
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
	expectedOHL := new([]OHLC)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedOHL); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		ch, err := GetOHLCPerDay(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		ohlcs := make([]OHLC, 0)
		for ohlc := range ch {
			ohlcs = append(ohlcs, ohlc)
		}
		if !reflect.DeepEqual(&ohlcs, expectedOHL) {
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
