package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAccount(t *testing.T) {
	accountBytes := ParseFile(t, "test_data/account.json")
	account := new(Account)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(accountBytes, &account); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = accountBytes
		client.ReturnError = nil
		_, err := GetAccount(client)
		if err != nil {
			t.Errorf(err.Error())
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
	placingOrder := Order{ISIN: "123123"}

	orderBytes := ParseFile(t, "test_data/create_order.json")
	order := new(Order)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(orderBytes, &order); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = orderBytes
		client.ReturnError = nil

		_, err := PlaceOrder(client, &placingOrder)
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
		_, err := PlaceOrder(client, &placingOrder)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := PlaceOrder(client, &placingOrder)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestGetOrder(t *testing.T) {
	orderBytes := ParseFile(t, "test_data/get_orders.json")
	order := new(Order)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(orderBytes, &order); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = orderBytes
		client.ReturnError = nil

		_, err := GetOrders(client)
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
		_, err := GetOrders(client)
		if err.Error() != errMessage {
			t.Errorf("Expected %s as error- message, got: %s", errMessage, err.Error())
		}
	})
	t.Run("fail to decode struct", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = []byte("bad")
		client.ReturnError = nil
		_, err := GetOrders(client)
		if err == nil {
			t.Errorf("expected error, got, nil")
		}
	})
}

func TestGetOrders(t *testing.T) {
	orderBytes := ParseFile(t, "test_data/get_order.json")
	order := new(Order)

	t.Run("parse struct", func(t *testing.T) {
		if err := json.Unmarshal(orderBytes, &order); err != nil {
			t.Errorf("error parsing struct: %w", err)
		}
	})
	t.Run("normal api response", func(t *testing.T) {
		client := GetMockedClient(t)
		client.ReturnData = orderBytes
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
		placed, err := PlaceOrder(client, &order)
		if err != nil {
			t.Errorf(err.Error())
		}
		orderID = placed.Results.ID
	})
	t.Run("Get Orders", func(t *testing.T) {
		orders, err := GetOrders(client)
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
