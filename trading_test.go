package main

import (
	"os"
	"testing"
	"time"
)

func GetClient(t *testing.T, env Envrionment) *Client {
	t.Helper()
	apiKey, isSet := os.LookupEnv("LEMON_API_KEY")
	if !isSet {
		t.Skip("missing environment variable LEMON_API_KEY")
	}
	c := Client{Envrionment: env, ApiKey: apiKey}
	return &c
}

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
