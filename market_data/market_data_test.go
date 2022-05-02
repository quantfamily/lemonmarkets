package market_data

import (
	"os"
	"testing"
)

func IntegrationClient(t *testing.T) *MarketDataClient {
	t.Helper()
	apiKey, isSet := os.LookupEnv("LEMON_API_KEY")
	if !isSet {
		t.Skip("missing environment variable LEMON_API_KEY")
	}
	return NewClient(apiKey)
}
