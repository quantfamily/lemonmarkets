package market_data

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/quantfamily/lemonmarkets/common"
	"github.com/quantfamily/lemonmarkets/common/helpers"
)

func TestGetInstruments(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_instruments.json")
	expectedResponse := new(common.Response)
	expectedInstruments := new([]Instrument)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedInstruments); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		ch, err := GetInstruments(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		instruments := make([]Instrument, 0)
		for i := range ch {
			instruments = append(instruments, i)
		}
		if !reflect.DeepEqual(&instruments, expectedInstruments) {
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

func TestGetInstrumentsIntegration(t *testing.T) {
	client := NewClient(helpers.APIKey(t))

	_, err := GetInstruments(client, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
