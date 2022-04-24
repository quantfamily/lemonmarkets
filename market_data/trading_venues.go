package market_data

/*
Venue of where the tradeable instrument is located
*/
type Venue struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Mic      string `json:"mic"`
	IsOpen   bool   `json:"is_open"`
	Tradable bool   `json:"tradable"`
	Currency string `json:"currency"`
}
