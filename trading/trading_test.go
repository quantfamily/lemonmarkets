package trading

import (
	"os"
	"testing"
)

func IntegrationClient(t *testing.T) *TradingClient {
	t.Helper()
	apiKey, isSet := os.LookupEnv("TRADING_API_KEY")
	if !isSet {
		t.Skip("missing environment variable TRADING_API_KEY")
	}
	return NewClient(apiKey, PAPER)
}
