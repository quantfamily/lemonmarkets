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

type Withdrawal struct {
	ID          string    `json:"id,omitempty"`
	Amount      int       `json:"amount,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Idempotency string    `json:"idempotency,omitempty"`
}

func CreateWithdrawal(client *client.Client, withdrawal *Withdrawal) error {
	withdrawData, err := json.Marshal(withdrawal)
	if err != nil {
		return err
	}
	_, err = client.Do("POST", "withdrawal", nil, withdrawData)
	return err
}

func GetWithdrawals(client *client.Client) <-chan Item[Withdrawal, error] {
	ch := make(chan Item[Withdrawal, error])
	go returnWithdrawals(client, ch)
	return ch
}

func returnWithdrawals(client *client.Client, ch chan<- Item[Withdrawal, error]) {
	defer close(ch)
	response, err := client.Do("GET", "withdrawals", nil, nil)
	if err != nil {
		withdrawal := Item[Withdrawal, error]{}
		withdrawal.Error = err
		ch <- withdrawal
		return
	}
	for {
		var withdrawals []Withdrawal
		withdrawal := Item[Withdrawal, error]{}
		withdrawal.Error = json.Unmarshal(response.Results, &withdrawals)
		if withdrawal.Error != nil {
			ch <- withdrawal
			return
		}
		for _, withdrawal := range withdrawals {
			ch <- Item[Withdrawal, error]{withdrawal, nil}
		}
		if response.Next == "" {
			return
		}
		response, withdrawal.Error = client.Do("GET", response.Next, nil, nil)
		if withdrawal.Error != nil {
			ch <- withdrawal
			return
		}
	}
}

type BankStatement struct {
	ID        string    `json:"id,omitempty"`
	AcountID  string    `json:"account_id,omitempty"`
	Type      string    `json:"type,omitempty"`
	Date      string    `json:"date,omitempty"` // TODO: Get this formatted to time.Time (YYYY-MM-DD)
	Amount    int       `json:"amount,omitempty"`
	ISIN      string    `json:"isin,omitempty"`
	ISINTitle string    `json:"isin_title,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func GetBankStatements(client *client.Client) <-chan Item[BankStatement, error] {
	ch := make(chan Item[BankStatement, error])
	go returnBankStatements(client, ch)
	return ch
}

func returnBankStatements(client *client.Client, ch chan<- Item[BankStatement, error]) {
	defer close(ch)
	response, err := client.Do("GET", "bankstatements", nil, nil)
	if err != nil {
		bankstatement := Item[BankStatement, error]{}
		bankstatement.Error = err
		ch <- bankstatement
		return
	}
	for {
		var bankstatements []BankStatement
		bankstatement := Item[BankStatement, error]{}
		bankstatement.Error = json.Unmarshal(response.Results, &bankstatements)
		if bankstatement.Error != nil {
			ch <- bankstatement
			return
		}
		for _, bankstatement := range bankstatements {
			ch <- Item[BankStatement, error]{bankstatement, nil}
		}
		if response.Next == "" {
			return
		}
		response, bankstatement.Error = client.Do("GET", response.Next, nil, nil)
		if bankstatement.Error != nil {
			ch <- bankstatement
			return
		}
	}
}

type Document struct {
	ID            string    `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	Category      string    `json:"category,omitempty"`
	Link          string    `json:"link,omitempty"`
	ViewedFirstAt time.Time `json:"viewed_first_at,omitempty"`
	ViewedLastAt  time.Time `json:"viewed_last_at,omitempty"`
}

func GetDocuments(client *client.Client) <-chan Item[Document, error] {
	ch := make(chan Item[Document, error])
	go returnDocuments(client, ch)
	return ch
}

func returnDocuments(client *client.Client, ch chan<- Item[Document, error]) {
	defer close(ch)
	response, err := client.Do("GET", "documents", nil, nil)
	if err != nil {
		document := Item[Document, error]{}
		document.Error = err
		ch <- document
		return
	}
	for {
		var documents []Document
		document := Item[Document, error]{}
		document.Error = json.Unmarshal(response.Results, &documents)
		if document.Error != nil {
			ch <- document
			return
		}
		for _, document := range documents {
			ch <- Item[Document, error]{document, nil}
		}
		if response.Next == "" {
			return
		}
		response, document.Error = client.Do("GET", response.Next, nil, nil)
		if document.Error != nil {
			ch <- document
			return
		}
	}
}
