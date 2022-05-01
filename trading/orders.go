package trading

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/quantfamily/lemonmarkets/client"
)

/*
Order information for a instrument
*/
type Order struct {
	ID                    string                 `json:"id,omitempty"`
	ISIN                  string                 `json:"isin,omitempty"`
	ISINTitle             string                 `json:"isin_title,omitempty"`
	ExpiresAt             time.Time              `json:"expires_at,omitempty"`
	CreatedAt             time.Time              `json:"created_at,omitempty"`
	Side                  string                 `json:"side,omitempty"`
	Quantity              int                    `json:"quantity,omitempty"`
	StopPrice             int                    `json:"stop_price,omitempty"`
	LimitPrice            int                    `json:"limit_price,omitempty"`
	EstimatedPrice        int                    `json:"estimated_price,omitempty"`
	EstimatedPriceTotal   int                    `json:"estimated_price_total,omitempty"`
	Venue                 string                 `json:"venue,omitempty"`
	Status                string                 `json:"status,omitempty"`
	Type                  string                 `json:"type,omitempty"`
	ExecutedQuantity      int                    `json:"executed_quantity,omitempty"`
	ExecutedPrice         int                    `json:"executed_price,omitempty"`
	ExecutedPriceTotal    int                    `json:"executed_price_total,omitempty"`
	ExecutedAt            time.Time              `json:"executed_at,omitempty"`
	RejectedAt            time.Time              `json:"rejected_at,omitempty"`
	Notes                 string                 `json:"notes,omitempty"`
	Charge                float64                `json:"charge,omitempty"`
	ChargeableAt          time.Time              `json:"chargeable_at,omitempty"`
	KeyCreationID         string                 `json:"key_creation_id,omitempty"`
	KeyActivationID       string                 `json:"key_activation_id,omitempty"`
	RegulatoryInformation *RegulatoryInformation `json:"regulatory_information,omitempty"`
	Idempotency           string                 `json:"idempotency,omitempty"`
}

/*
RegulatoryInformation information for an order
*/
type RegulatoryInformation struct {
	CostsEntry                      float64 `json:"costs_entry"`
	CostsEntryPct                   string  `json:"costs_entry_pct"`
	CostsRunning                    float64 `json:"costs_running"`
	CostsRunningPct                 string  `json:"costs_running_pct"`
	CostsProduct                    float64 `json:"costs_product"`
	CostsProductPct                 string  `json:"costs_product_pct"`
	CostsExit                       float64 `json:"costs_exit"`
	CostsExitPct                    string  `json:"costs_exit_pct"`
	YieldReductionYear              float64 `json:"yield_reduction_year"`
	YieldReductionYearPct           string  `json:"yield_reduction_year_pct"`
	YieldReductionYearFollowing     float64 `json:"yield_reduction_year_following"`
	YieldReductionYearFollowingPct  string  `json:"yield_reduction_year_following_pct"`
	YieldReductionYearExit          float64 `json:"yield_reduction_year_exit"`
	YieldReductionYearExitPct       string  `json:"yield_reduction_year_exit_pct"`
	EstimatedHoldingDurationYears   string  `json:"estimated_holding_duration_years"`
	EstimatedYieldReductionTotal    float64 `json:"estimated_yield_reduction_total"`
	EstimatedYieldReductionTotalPct string  `json:"estimated_yield_reduction_total_pct"`
	KIID                            string  `json:"KIID"`
	LegalDisclaimer                 string  `json:"legal_disclaimer"`
}

/*
CreateOrder places a order on LemonMarkets and returns response from the backend
*/
func CreateOrder(client *client.Client, order *Order) *Item[Order, error] {
	item := &Item[Order, error]{}

	orderData, err := json.Marshal(order)
	if err != nil {
		item.Error = err
		return item
	}

	response, err := client.Do("POST", "orders", nil, orderData)
	if err != nil {
		item.Error = err
		return item
	}
	item.Error = json.Unmarshal(response.Results, &item.Data)
	return item
}

/*
ActivateOrder activates a placed order on LemonMarkets to go into execution
*/
func ActivateOrder(client *client.Client, orderID string) error {
	_, err := client.Do("POST", fmt.Sprintf("orders/%s/activate", orderID), nil, nil)
	return err
}

/*
GetOrdersQuery is used to filter order when trying to received a list of placed orders
Read more at: https://docs.lemon.markets/trading/orders#get-orders
*/
type GetOrdersQuery struct {
	From          time.Time `json:"from,omitempty"`
	To            time.Time `json:"to,omitempty"`
	ISIN          string    `json:"isin,omitempty"`
	Side          string    `json:"side,omitempty"`
	Status        string    `json:"status,omitempty"`
	Type          string    `json:"type,omitempty"`
	KeyCreationID string    `json:"key_creation_id,omitempty"`
	Limit         int       `json:"limit,omitempty"`
	Page          int       `json:"page,omitempty"`
}

/*
GetOrders can take a query parameters and return one or more orders embedded a result in Response- object
*/
func GetOrders(client *client.Client, query *GetOrdersQuery) <-chan Item[Order, error] {
	ch := make(chan Item[Order, error])
	go returnOrders(client, query, ch)
	return ch
}

func returnOrders(client *client.Client, query *GetOrdersQuery, ch chan<- Item[Order, error]) {
	defer close(ch)
	response, err := client.Do("GET", "orders", query, nil)
	if err != nil {
		order := Item[Order, error]{}
		order.Error = err
		ch <- order
		return
	}
	for {
		var orders []Order
		order := Item[Order, error]{}
		order.Error = json.Unmarshal(response.Results, &orders)
		if order.Error != nil {
			ch <- order
			return
		}
		for _, order := range orders {
			ch <- Item[Order, error]{order, nil}
		}
		if response.Next == "" {
			return
		}
		response, order.Error = client.Do("GET", response.Next, nil, nil)
		if order.Error != nil {
			ch <- order
			return
		}
	}
}

/*
GetOrder returns a placed order based on a specific orderID
*/
func GetOrder(client *client.Client, orderID string) *Item[Order, error] {
	order := &Item[Order, error]{}
	response, err := client.Do("GET", fmt.Sprintf("orders/%s", orderID), nil, nil)
	if err != nil {
		order.Error = err
		return order
	}
	order.Error = json.Unmarshal(response.Results, &order.Data)
	return order
}

/*
DeleteOrder deletes a placed order and makes unable to be activated and executed
*/
func DeleteOrder(client *client.Client, orderID string) error {
	_, err := client.Do("DELETE", fmt.Sprintf("orders/%s", orderID), nil, nil)
	return err
}
