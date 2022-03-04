package main

import (
	"testing"
	"time"
)

func TestGetInstruments(t *testing.T) {
	client := GetClient(t, DATA)

	_, err := GetInstruments(client)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetQuotes(t *testing.T) {
	client := GetClient(t, DATA)

	_, err := GetQuotes(client, "DE000CBK1001")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetOHLC(t *testing.T) {
	client := GetClient(t, DATA)
	isin := "DE000CBK1001"

	t.Run("Minute", func(t *testing.T) {
		from := time.Now().AddDate(0, -1, 0).Unix()
		to := time.Now().Unix()
		_, err := GetOHLCPerMinute(client, from, to, isin)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("Hour", func(t *testing.T) {
		from := time.Now().AddDate(0, -1, 0).Unix()
		to := time.Now().Unix()
		_, err := GetOHLCPerHour(client, from, to, isin)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("Day", func(t *testing.T) {
		from := time.Now().AddDate(0, -1, 0).Unix()
		to := time.Now().Unix()
		_, err := GetOHLCPerDay(client, from, to, isin)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}

func TestGetTrades(t *testing.T) {
	client := GetClient(t, DATA)
	from := time.Now().AddDate(0, -1, 0).Unix()
	to := time.Now().Unix()
	_, err := GetTrades(client, from, to, "DE000CBK1001")
	if err != nil {
		t.Errorf(err.Error())
	}
}
