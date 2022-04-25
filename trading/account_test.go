package trading

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/quantfamily/lemonmarkets/common"
	"github.com/quantfamily/lemonmarkets/common/helpers"
)

func TestAccount(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_account.json")
	expectedResponse := new(common.Response)
	expectedAccount := new(Account)
	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedAccount); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}

	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		clientResponse, err := GetAccount(client)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(clientResponse, expectedAccount) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error getting account"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnResponse = nil
		client.ReturnError = lemonErr
		_, err := GetAccount(client)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}

func TestAccountIntegration(t *testing.T) {
	client := NewClient(helpers.APIKey(t), PAPER)
	_, err := GetAccount(client)
	if err != nil {
		t.Errorf("Failure to get account %v", err)
	}
}
