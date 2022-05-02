package trading

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/quantfamily/lemonmarkets/client"
	"github.com/quantfamily/lemonmarkets/client/helpers"
	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_account.json")

	t.Run("fail to get response", func(t *testing.T) {
		expectedErr := client.LemonError{
			Time:    time.Time{},
			Mode:    "paper",
			Status:  "error",
			Code:    "order_total_price_limit_exceeded",
			Message: "cannot place/activate buy order if estimated total price is greater than 25k Euro",
		}
		errRsp, _ := json.Marshal(&expectedErr)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, string(errRsp), 400)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		account := client.GetAccount()
		assert.NotNil(t, account.Error)
		assert.Equal(t, &expectedErr, account.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		account := client.GetAccount()
		assert.NotNil(t, account.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, account.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		account := client.GetAccount()
		assert.Nil(t, account.Error)
		assert.Equal(t, "basic", account.Data.TradingPlan)
		assert.Equal(t, "K2057263187", account.Data.DepositID)
		assert.Equal(t, "m_burry@tradingapi.com", account.Data.EMail)

	})
}

func TestCreateWithdrawal(t *testing.T) {
	t.Run("fail to get response", func(t *testing.T) {
		expectedErr := client.LemonError{
			Time:    time.Time{},
			Mode:    "paper",
			Status:  "error",
			Code:    "order_total_price_limit_exceeded",
			Message: "cannot place/activate buy order if estimated total price is greater than 25k Euro",
		}
		errRsp, _ := json.Marshal(&expectedErr)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, string(errRsp), 400)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		err := client.CreateWithdrawal(&Withdrawal{Amount: 10})
		assert.NotNil(t, err)
		assert.Equal(t, &expectedErr, err)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{"status": "ok"}`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		err := client.CreateWithdrawal(&Withdrawal{Amount: 10})
		assert.Nil(t, err)
	})
}

func TestGetWithdrawals(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_withdrawals.json")

	t.Run("fail to get response", func(t *testing.T) {
		expectedErr := client.LemonError{
			Time:    time.Time{},
			Mode:    "paper",
			Status:  "error",
			Code:    "order_total_price_limit_exceeded",
			Message: "cannot place/activate buy order if estimated total price is greater than 25k Euro",
		}
		errRsp, _ := json.Marshal(&expectedErr)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, string(errRsp), 400)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		withdrawalCh := client.GetWithdrawals()
		withdrawal := <-withdrawalCh
		assert.NotNil(t, withdrawal.Error)
		assert.Equal(t, &expectedErr, withdrawal.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		withdrawalCh := client.GetWithdrawals()
		withdrawal := <-withdrawalCh
		assert.NotNil(t, withdrawal.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, withdrawal.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		withdrawalCh := client.GetWithdrawals()
		withdrawal := <-withdrawalCh
		assert.Nil(t, withdrawal.Error)
		assert.Equal(t, 1000000, withdrawal.Data.Amount)
	})
}

func TestGetBankStatements(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_bankstatements.json")

	t.Run("fail to get response", func(t *testing.T) {
		expectedErr := client.LemonError{
			Time:    time.Time{},
			Mode:    "paper",
			Status:  "error",
			Code:    "order_total_price_limit_exceeded",
			Message: "cannot place/activate buy order if estimated total price is greater than 25k Euro",
		}
		errRsp, _ := json.Marshal(&expectedErr)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, string(errRsp), 400)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		bankstatementCh := client.GetBankStatements()
		bankstatement := <-bankstatementCh
		assert.NotNil(t, bankstatement.Error)
		assert.Equal(t, &expectedErr, bankstatement.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		bankstatementCh := client.GetBankStatements()
		bankstatement := <-bankstatementCh
		assert.NotNil(t, bankstatement.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, bankstatement.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		bankstatementCh := client.GetBankStatements()
		bankstatement := <-bankstatementCh
		assert.Nil(t, bankstatement.Error)
		assert.Equal(t, 100000, bankstatement.Data.Amount)
	})
}

func TestGetDocuments(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_documents.json")

	t.Run("fail to get response", func(t *testing.T) {
		expectedErr := client.LemonError{
			Time:    time.Time{},
			Mode:    "paper",
			Status:  "error",
			Code:    "order_total_price_limit_exceeded",
			Message: "cannot place/activate buy order if estimated total price is greater than 25k Euro",
		}
		errRsp, _ := json.Marshal(&expectedErr)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, string(errRsp), 400)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		documentCh := client.GetDocuments()
		document := <-documentCh
		assert.NotNil(t, document.Error)
		assert.Equal(t, &expectedErr, document.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		documentCh := client.GetDocuments()
		document := <-documentCh
		assert.NotNil(t, document.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, document.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		backend := client.Backend{BaseURL: server.URL}
		client := TradingClient{backend: &backend}
		documentCh := client.GetDocuments()
		document := <-documentCh
		assert.Nil(t, document.Error)
		assert.Equal(t, "account_opening.pdf", document.Data.Name)
	})
}

func TestGetAccountIntegration(t *testing.T) {
	client := IntegrationClient(t)
	account := client.GetAccount()

	assert.Nil(t, account.Error)
}

func TestGetWithdrawalIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ch := client.GetWithdrawals()

	withdrawal := <-ch

	assert.Nil(t, withdrawal.Error)
}

func TestGetWithdrawalsIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ch := client.GetWithdrawals()

	withdrawal := <-ch
	assert.Nil(t, withdrawal.Error)
}

func TestGetBankStatementsIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ch := client.GetBankStatements()

	bankstatement := <-ch
	assert.Nil(t, bankstatement.Error)
}

func TestGetDocumentsIntegration(t *testing.T) {
	client := IntegrationClient(t)
	ch := client.GetDocuments()

	document := <-ch
	assert.Nil(t, document.Error)
}
