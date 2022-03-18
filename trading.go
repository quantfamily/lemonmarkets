package lemonmarkets

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
GetAccountResponse response regarding account information from LemonMarkets
Read more at: https://docs.lemon.markets/trading/account
*/
type GetAccountResponse struct {
	Reply
	Results Account `json:"results"`
}

/*
Account details about registered account
*/
type Account struct {
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

/*
GetAccount returns account information from the used, based on the API Key
*/
func GetAccount(client Client) (*GetAccountResponse, error) {
	responseData, err := client.Do("GET", "account", nil, nil)
	if err != nil {
		return nil, err
	}

	account := GetAccountResponse{}
	err = json.Unmarshal(responseData, &account)
	return &account, err
}

/*
CreateOrderResponse from placing a order to LemonMarkets
Read more at: https://docs.lemon.markets/trading/orders#placing-an-order
*/
type CreateOrderResponse struct {
	Reply
	Results Order `json:"results"`
}

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
func CreateOrder(client Client, order *Order) (*CreateOrderResponse, error) {
	orderData, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	responseData, err := client.Do("POST", "orders", nil, orderData)
	if err != nil {
		return nil, err
	}
	createdOrder := new(CreateOrderResponse)
	err = json.Unmarshal(responseData, createdOrder)
	return createdOrder, err
}

/*
ActivateOrder activates a placed order on LemonMarkets to go into execution
*/
func ActivateOrder(client Client, orderID string) error {
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
ActivatedOrder is an addition to the normal order, with a key_activation_id embedded
*/
type ActivatedOrder struct {
	Order
	KeyActivationID string `json:"key_activation_id"`
}

/*
GetOrdersResponse returns a list of Activated orders from LemonMarkets
Read more at: https://docs.lemon.markets/trading/orders#get-orders
*/
type GetOrdersResponse struct {
	ListReply
	Results []ActivatedOrder `json:"results"`
}

/*
GetOrders can take a query paramters and return one or more orders embedded a result in Response- object
*/
func GetOrders(client Client, query *GetOrdersQuery) (*GetOrdersResponse, error) {
	responseData, err := client.Do("GET", "orders", query, nil)
	if err != nil {
		return nil, err

	}
	orderResult := new(GetOrdersResponse)
	err = json.Unmarshal(responseData, orderResult)
	return orderResult, err
}

/*
GetOrderResponse response for a specific order
Read more at: https://docs.lemon.markets/trading/orders#get-ordersorder_id
*/
type GetOrderResponse struct {
	Reply
	Results ActivatedOrder `json:"results"`
}

/*
GetOrder returns a placed order based on a specific orderID
*/
func GetOrder(client Client, orderID string) (*GetOrderResponse, error) {
	responseData, err := client.Do("GET", fmt.Sprintf("orders/%s", orderID), nil, nil)
	if err != nil {
		return nil, err
	}
	order := new(GetOrderResponse)
	err = json.Unmarshal(responseData, order)
	return order, err
}

/*
DeleteOrder deletes a placed order and makes unable to be activated and executed
*/
func DeleteOrder(client Client, orderID string) error {
	_, err := client.Do("DELETE", fmt.Sprintf("orders/%s", orderID), nil, nil)
	return err
}

/*
GetPortfolioResult returns information about Portfolios status on LemonMarkets
Read more at: https://docs.lemon.markets/trading/portfolio
*/
type GetPortfolioResult struct {
	ListReply
	Results []PortfolioPosition `json:"results"`
}

/*
PortfolioPosition is information about Positions inside the Portfolio
*/
type PortfolioPosition struct {
	ISIN                string  `json:"isin"`
	ISINTitle           string  `json:"isin_title"`
	Quantity            int     `json:"quantity"`
	BuyPriceAverage     float64 `json:"buy_price_avg"`
	EstimatedPriceTotal float64 `json:"estimated_price_total"`
	EstimatedPrice      float64 `json:"estimated_price"`
}

/*
GetPortfolio returns current portfolio in LemonMarkets
*/
func GetPortfolio(client Client) (*GetPortfolioResult, error) {
	responseData, err := client.Do("GET", "portfolio", nil, nil)
	if err != nil {
		return nil, err
	}
	portfolioResult := new(GetPortfolioResult)
	err = json.Unmarshal(responseData, portfolioResult)
	return portfolioResult, err
}
