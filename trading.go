package main

//
//Withdrawal
//
// TODO: Fix reset of account stuff

import (
	"encoding/json"
	"fmt"
	"time"
)

type AccountResult struct {
	CreatedAt         time.Time `json:"created_at"`
	AccountID         string    `json:"account_id"`
	Firstname         string    `json:"firstname"`
	Lastname          string    `json:"Lastname"`
	EMail             string    `json:"email"`
	Phone             string    `json:"phone"`
	Address           string    `json:"address"`
	BillingAddress    string    `json:"billing_address"`
	BillingEMail      string    `json:"billing_email"`
	BillingName       string    `json:"billing_name"`
	BillingVAT        string    `json:"billing_vat"`
	Mode              string    `json:"mode"`
	DepositID         string    `json:"deposit_id"`
	ClientID          string    `json:"client_id"`
	AccountNumber     string    `json:"account_number"`
	IBANBrokerage     string    `json:"iban_brokerage"`
	IBANOrigin        string    `json:"iban_origin"`
	BankNameOrigin    string    `json:"bank_name_origin"`
	Balance           float32   `json:"balance"`
	CashToInvest      float32   `json:"cash_to_invest"`
	CashToWithdraw    float32   `json:"cash_to_withdraw"`
	TradingPlan       string    `json:"basic"`
	DataPlan          string    `json:"data_plan"`
	TaxAllowance      string    `json:"tax_allowance"`
	TaxAllowanceStart time.Time `json:"tax_allowance_start"`
	TaxAllowanceEnd   time.Time `json:"tax_allowance_end"`
}

type Account struct {
	Reply
	Results AccountResult `json:"results"`
}

func GetAccount(client Client) (*Account, error) {
	responseData, err := client.Do("GET", "account", nil)
	if err != nil {
		return nil, err
	}

	account := Account{}
	err = json.Unmarshal(responseData, &account)
	return &account, err
}

func PlaceOrder(client Client, order *Order) (*GetOrderResult, error) {
	orderData, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	responseData, err := client.Do("POST", "orders", orderData)
	if err != nil {
		return nil, err
	}
	placedOrder := new(GetOrderResult)
	err = json.Unmarshal(responseData, placedOrder)
	return placedOrder, err
}

func ActivateOrder(client Client, orderID string) error {
	_, err := client.Do("POST", fmt.Sprintf("orders/%s/activate", orderID), nil)
	return err
}

func GetOrders(client Client) (*GetOrdersResult, error) {
	responseData, err := client.Do("GET", "orders", nil)
	if err != nil {
		return nil, err

	}
	orderResult := new(GetOrdersResult)
	err = json.Unmarshal(responseData, orderResult)
	return orderResult, err
}

func GetOrder(client Client, orderID string) (*GetOrderResult, error) {
	responseData, err := client.Do("GET", fmt.Sprintf("orders/%s", orderID), nil)
	if err != nil {
		return nil, err
	}
	order := new(GetOrderResult)
	err = json.Unmarshal(responseData, order)
	return order, err
}

func DeleteOrder(client Client, orderID string) error {
	_, err := client.Do("DELETE", fmt.Sprintf("orders/%s", orderID), nil)
	return err
}

type GetOrdersResult struct {
	ListReply
	Results []ActivatedOrder `json:"results"`
}

type GetOrderResult struct {
	Reply
	Results ActivatedOrder `json:"results"`
}

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
	RegulatoryInformation *RegulatoryInformation `json:"regulatory_information,omitempty"`
}

type ActivatedOrder struct {
	Order
	KeyActivationID string `json:"key_activation_id"`
}

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

func GetPortfolio(client Client) (*GetPortfolioResult, error) {
	responseData, err := client.Do("GET", "portfolio", nil)
	if err != nil {
		return nil, err
	}
	portfolioResult := new(GetPortfolioResult)
	err = json.Unmarshal(responseData, portfolioResult)
	return portfolioResult, err
}

type GetPortfolioResult struct {
	ListReply
	Results PortfolioPosition `json:"results"`
}

type PortfolioPosition struct {
	ISIN                string  `json:"isin"`
	ISINTitle           string  `json:"isin_title"`
	Quantity            int     `json:"quantity"`
	BuyPriceAverage     float64 `json:"buy_price_avg"`
	EstimatedPriceTotal float64 `json:"estimated_price_total"`
	EstimatedPrice      float64 `json:"estimated_price"`
}
