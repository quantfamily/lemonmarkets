package trading

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/quantfamily/lemonmarkets/common"
)

/*
Order information for a instrument
*/
type Order struct {
	CreatedAt             time.Time              `json:"created_at,omitempty"`
	ID                    string                 `json:"id,omitempty"`
	Status                string                 `json:"status,omitempty"`
	ISIN                  string                 `json:"isin,omitempty"`
	ExpiresAt             time.Time              `json:"expires_at,omitempty"`
	Side                  string                 `json:"side,omitempty"`
	Quantity              int                    `json:"quantity,omitempty"`
	StopPrice             float64                `json:"stop_price,omitempty"`
	LimitPrice            float64                `json:"limit_price,omitempty"`
	Venue                 string                 `json:"venue,omitempty"`
	EstimatedPrice        float64                `json:"estimated_price,omitempty"`
	Notes                 string                 `json:"notes,omitempty"`
	Idempotency           string                 `json:"idempotency,omitempty"`
	Charge                float64                `json:"charge,omitempty"`
	ChargeableAt          time.Time              `json:"chargeable_at,omitempty"`
	KeyCreationID         string                 `json:"key_creation_id,omitempty"`
	KeyActivationID       string                 `json:"key_activation_id,omitempty"`
	RegulatoryInformation *RegulatoryInformation `json:"regulatory_information,omitempty"`
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
func CreateOrder(client common.Client, order *Order) (*Order, error) {
	orderData, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	response, err := client.Do("POST", "orders", nil, orderData)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response.Results, order)
	return order, err
}

/*
ActivateOrder activates a placed order on LemonMarkets to go into execution
*/
func ActivateOrder(client common.Client, orderID string) error {
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
func GetOrders(client common.Client, query *GetOrdersQuery) (<-chan Order, error) {
	response, err := client.Do("GET", "orders", query, nil)
	if err != nil {
		return nil, err

	}
	var orders []Order
	err = json.Unmarshal(response.Results, &orders)
	if err != nil {
		return nil, err
	}
	ch := make(chan Order)
	go returnOrders(client, response, ch)
	return ch, nil
}

func returnOrders(client common.Client, response *common.Response, outchan chan<- Order) {
	defer close(outchan)
	var orders []Order
	for {
		err := json.Unmarshal(response.Results, &orders)
		if err != nil {
			return
		}
		for _, order := range orders {
			outchan <- order
		}
		if response.Next == "" {
			return
		}
		response, err = client.Do("GET", response.Next, nil, nil)
		if err != nil {
			return
		}
	}
}

/*
GetOrder returns a placed order based on a specific orderID
*/
func GetOrder(client common.Client, orderID string) (*Order, error) {
	response, err := client.Do("GET", fmt.Sprintf("orders/%s", orderID), nil, nil)
	if err != nil {
		return nil, err
	}
	var order Order
	err = json.Unmarshal(response.Results, &order)
	return &order, err
}

/*
DeleteOrder deletes a placed order and makes unable to be activated and executed
*/
func DeleteOrder(client common.Client, orderID string) error {
	_, err := client.Do("DELETE", fmt.Sprintf("orders/%s", orderID), nil, nil)
	return err
}
