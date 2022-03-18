package lemonmarkets

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {
	t.Run("Do, get LemonError", func(t *testing.T) {
		errMessage := "generic error message"
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			lemonError := LemonError{Message: errMessage}
			response, _ := json.Marshal(lemonError)
			w.Write(response)
		}))
		defer server.Close()

		client := LemonClient{Environment: Environment(server.URL)}
		_, err := client.Do("demo", "demo", nil, nil)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err.Error() != errMessage {
			t.Errorf("expected message to be %s, received: %s", errMessage, err.Error())
		}
	})
	t.Run("Do, get unknown error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte("very bad"))
		}))
		defer server.Close()

		client := LemonClient{Environment: Environment(server.URL)}
		_, err := client.Do("demo", "demo", nil, nil)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
	t.Run("Do, normal request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			lemonError := Reply{Status: "very ok"}
			response, _ := json.Marshal(lemonError)
			w.Write(response)
		}))
		defer server.Close()

		client := LemonClient{Environment: Environment(server.URL)}
		_, err := client.Do("demo", "demo", nil, nil)
		if err != nil {
			t.Errorf("Expected nil, got error: %w", err)
		}
	})
	t.Run("Do, with querystring", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			lemonError := Reply{Status: "very ok"}
			response, _ := json.Marshal(lemonError)
			expectedQuery := "isin=isin1&isin=isin2&mic=my_mic"
			if r.URL.RawQuery != expectedQuery {
				t.Errorf("Expected query %s, got %s", expectedQuery, r.URL.RawQuery)
			}
			w.Write(response)
		}))
		defer server.Close()

		client := LemonClient{Environment: Environment(server.URL)}
		isins := []string{"isin1", "isin2"}
		query := GetInstrumentsQuery{MIC: "my_mic", ISIN: isins}
		_, err := client.Do("demo", "demo", &query, nil)
		if err != nil {
			t.Errorf("Expected nil, got error: %w", err)
		}
	})
}

func TestListReply(t *testing.T) {
	t.Run("Next, next is empty", func(t *testing.T) {
		lr := ListReply{}
		client := NewClient(PAPER, "")

		err := lr.Next(client)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err.Error() != "end of list" {
			t.Errorf("expected end of list, got: %w", err)
		}
	})
	t.Run("Next, url is not correct", func(t *testing.T) {
		lr := ListReply{next: "http://bad_url.com/no/v2/no"}
		client := NewClient(PAPER, "")

		err := lr.Next(client)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err.Error() != "url is not correct" {
			t.Errorf("expected url is not correct, got: %w", err)
		}
	})
	t.Run("Next client error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte("very bad"))
		}))
		defer server.Close()

		lr := ListReply{next: "http://bad_url.com/works/v1/omg_call"}
		client := LemonClient{Environment: Environment(server.URL)}

		err := lr.Next(&client)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
	t.Run("Next, all OK", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			lr := ListReply{Reply: Reply{Status: "OK"}}
			responseBytes, _ := json.Marshal(&lr)
			w.Write(responseBytes)
		}))
		defer server.Close()

		lr := ListReply{next: "http://good_url.com/works/v1/omg_call"}
		client := LemonClient{Environment: Environment(server.URL)}

		err := lr.Next(&client)
		if err != nil {
			t.Errorf("expected nil, got error: %w", err)
		}
	})
}
