package trading

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/quantfamily/lemonmarkets/common"
	"github.com/quantfamily/lemonmarkets/common/helpers"
)

func TestCreateOrder(t *testing.T) {
	orderToCreate := Order{ISIN: "123123"}
	rawFileBytes := helpers.ParseFile(t, "create_order.json")
	expectedResponse := new(common.Response)
	expectedOrder := new(Order)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedOrder); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		clientResponse, err := CreateOrder(client, &orderToCreate)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(clientResponse, expectedOrder) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error creating order"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnResponse = nil
		client.ReturnError = lemonErr
		_, err := CreateOrder(client, &orderToCreate)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}

func TestActivateOrder(t *testing.T) {
	t.Run("normal response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = nil
		client.ReturnError = nil
		err := ActivateOrder(client, "abc123")
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
	t.Run("error response", func(t *testing.T) {
		errMessage := "error deleting order"
		lemonErr := common.LemonError{Message: errMessage}
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = nil
		client.ReturnError = lemonErr
		err := ActivateOrder(client, "abc123")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestGetOrders(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_orders.json")
	expectedResponse := new(common.Response)
	expectedOrders := new([]Order)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedOrders); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response, nil query", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil

		_, err := GetOrders(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("normal api response, with query", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil
		query := GetOrdersQuery{Side: "buy"}

		_, err := GetOrders(client, &query)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error placing order"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnResponse = nil
		client.ReturnError = lemonErr
		_, err := GetOrders(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}

func TestGetOrder(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_order.json")
	expectedResponse := new(common.Response)
	expectedOrder := new(Order)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
		if err := json.Unmarshal(expectedResponse.Results, expectedOrder); err != nil {
			t.Errorf("error parsing struct: %v", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = expectedResponse
		client.ReturnError = nil

		order, err := GetOrder(client, "123")
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(order, expectedOrder) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error placing order"
		lemonErr := common.LemonError{Message: errMessage}

		client := helpers.GetMockedClient(t)
		client.ReturnResponse = nil
		client.ReturnError = lemonErr
		_, err := GetOrder(client, "123")
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
}

func TestDeleteOrder(t *testing.T) {
	t.Run("delete, normal", func(t *testing.T) {
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = nil
		client.ReturnError = nil
		err := DeleteOrder(client, "abc123")
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error deleting order"
		lemonErr := common.LemonError{Message: errMessage}
		client := helpers.GetMockedClient(t)
		client.ReturnResponse = nil
		client.ReturnError = lemonErr
		err := DeleteOrder(client, "abc123")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestOrderIntegration(t *testing.T) {
	ISIN := "DE000CBK1001"
	var orderID string
	client := NewClient(helpers.APIKey(t), PAPER)
	t.Run("Place Order", func(t *testing.T) {
		expiresAt := time.Now().AddDate(0, 0, 14)

		order := Order{ISIN: ISIN, Side: "buy", ExpiresAt: expiresAt, Quantity: 1, Venue: "XMUN"}
		placed, err := CreateOrder(client, &order)
		if err != nil {
			t.Errorf(err.Error())
		}
		orderID = placed.ID
	})
	t.Run("Get Orders", func(t *testing.T) {
		orders, err := GetOrders(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		var found bool
		for order := range orders {
			if order.ID == orderID {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("did not find order in order group")
		}
	})
	t.Run("Get Order", func(t *testing.T) {
		_, err := GetOrder(client, orderID)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("Delete order", func(t *testing.T) {
		err := DeleteOrder(client, orderID)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
