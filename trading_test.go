package main

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestAccount(t *testing.T) {
	rawFileBytes := ParseFile(t, "test_data/get_account.json")
	expectedResponse := new(GetAccountResponse)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
		if expectedResponse.Status != "ok" {
			t.Errorf("Expected status to be ok, got: %s", expectedResponse.Status)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
		client.ReturnError = nil
		clientResponse, err := GetAccount(client)
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
		_, err := GetAccount(client)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetAccount(client)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestCreateOrder(t *testing.T) {
	orderToCreate := Order{ISIN: "123123"}
	rawFileBytes := ParseFile(t, "test_data/create_order.json")
	expectedResponse := new(CreateOrderResponse)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
		if expectedResponse.Status != "ok" {
			t.Errorf("Expected status to be ok, got: %s", expectedResponse.Status)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
		client.ReturnError = nil
		clientResponse, err := CreateOrder(client, &orderToCreate)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(clientResponse, expectedResponse) {
			t.Errorf("Not equal")
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error creating order"
		lemonErr := LemonError{Message: errMessage}

		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		_, err := CreateOrder(client, &orderToCreate)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := CreateOrder(client, &orderToCreate)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestActivateOrder(t *testing.T) {
	t.Run("normal response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = nil
		err := ActivateOrder(client, "abc123")
		if err != nil {
			t.Errorf("Expected nil error, got %w", err)
		}
	})
	t.Run("error response", func(t *testing.T) {
		errMessage := "error deleting order"
		lemonErr := LemonError{Message: errMessage}
		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		err := ActivateOrder(client, "abc123")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestGetOrders(t *testing.T) {
	rawFileBytes := ParseFile(t, "test_data/get_orders.json")
	expectedResponse := new(GetOrdersResponse)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
		if expectedResponse.Status != "ok" {
			t.Errorf("Expected status to be ok, got: %s", expectedResponse.Status)
		}
	})
	t.Run("normal api response, nil query", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
		client.ReturnError = nil

		_, err := GetOrders(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("normal api response, with query", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
		client.ReturnError = nil
		query := GetOrdersQuery{Side: "buy"}

		_, err := GetOrders(client, &query)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error placing order"
		lemonErr := LemonError{Message: errMessage}

		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		_, err := GetOrders(client, nil)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetOrders(client, nil)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestGetOrder(t *testing.T) {
	rawFileBytes := ParseFile(t, "test_data/get_order.json")
	expectedResponse := new(GetOrderResponse)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
		if expectedResponse.Status != "ok" {
			t.Errorf("Expected status to be ok, got: %s", expectedResponse.Status)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
		client.ReturnError = nil

		_, err := GetOrder(client, "123")
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error placing order"
		lemonErr := LemonError{Message: errMessage}

		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		_, err := GetOrder(client, "123")
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetOrder(client, "123")
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestDeleteOrder(t *testing.T) {
	t.Run("delete, normal", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = nil
		err := DeleteOrder(client, "abc123")
		if err != nil {
			t.Errorf("Expected nil error, got %w", err)
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error deleting order"
		lemonErr := LemonError{Message: errMessage}
		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		err := DeleteOrder(client, "abc123")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestPortfolio(t *testing.T) {
	rawFileBytes := ParseFile(t, "test_data/get_portfolio.json")
	expectedResponse := new(GetPortfolioResult)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(rawFileBytes, expectedResponse); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
		if expectedResponse.Status != "ok" {
			t.Errorf("Expected status to be ok, got: %s", expectedResponse.Status)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = rawFileBytes
		client.ReturnError = nil

		_, err := GetPortfolio(client)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("err api response", func(t *testing.T) {
		errMessage := "error placing order"
		lemonErr := LemonError{Message: errMessage}

		client := GetMockedClient(t)
		client.ReturnData = nil
		client.ReturnError = lemonErr
		_, err := GetPortfolio(client)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetPortfolio(client)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

/*
Integration tests below
*/
func TestAccountIntegration(t *testing.T) {
	client := GetClient(t, PAPER)
	_, err := GetAccount(client)
	if err != nil {
		t.Errorf("Failure to get account %w", err)
	}
}

func TestOrderIntegration(t *testing.T) {
	ISIN := "DE000CBK1001"
	var orderID string
	client := GetClient(t, PAPER)
	t.Run("Place Order", func(t *testing.T) {
		expires_at := time.Now().AddDate(0, 0, 14)

		order := Order{ISIN: ISIN, Side: "buy", ExpiresAt: expires_at, Quantity: 1, Venue: "XMUN"}
		placed, err := CreateOrder(client, &order)
		if err != nil {
			t.Errorf(err.Error())
		}
		orderID = placed.Results.ID
	})
	t.Run("Get Orders", func(t *testing.T) {
		orders, err := GetOrders(client, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
		var found bool
		for _, order := range orders.Results {
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
