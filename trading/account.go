package trading

import (
	"encoding/json"
	"time"

	"github.com/quantfamily/lemonmarkets/client"
)

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
func GetAccount(client *client.Client) *Item[Account, error] {
	account := &Item[Account, error]{}
	responseData, err := client.Do("GET", "account", nil, nil)
	if err != nil {
		account.Error = err
		return account
	}
	account.Error = json.Unmarshal(responseData.Results, &account.Data)
	return account
}
