package trading

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/quantfamily/lemonmarkets/common"
	"github.com/quantfamily/lemonmarkets/common/helpers"
)

func TestPositions(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_positions.json")
	expectedResponse := new(common.Response)
	expectedPositions := new([]Position)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedPositions); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil

		pChan, err := GetPositions(client)
		client.ReturnResponse = nil
		if err != nil {
			t.Errorf(err.Error())
		}
		positions := make([]Position, 0)
		for p := range pChan {
			positions = append(positions, p)
		}
		if !reflect.DeepEqual(&positions, expectedPositions) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error placing order"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnResponse = nil
		client.ReturnError = lemonErr
		_, err := GetPositions(client)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}
