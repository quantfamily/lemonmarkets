package lemonmarkets

/*
Environment reference a type of LemonMarkets Environment to the corresponding base url
*/
type Environment string

const (
	// PAPER trading without using real cash
	PAPER Environment = "https://paper-trading.lemon.markets/v1"
	// LIVE trading using real cash
	LIVE Environment = "https://trading.lemon.markets/v1"
	// DATA to request market data to do analysis on
	DATA Environment = "https://data.lemon.markets/v1"
)
