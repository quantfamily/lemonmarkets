package client

import (
	"encoding/json"
	"fmt"
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

		client := Client{BaseURL: server.URL}
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

		client := Client{BaseURL: server.URL}
		_, err := client.Do("demo", "demo", nil, nil)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
	t.Run("Do, normal request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{"status": "ok"}`)
		}))
		defer server.Close()

		client := Client{BaseURL: server.URL}
		_, err := client.Do("demo", "demo", nil, nil)
		if err != nil {
			t.Errorf("Expected nil, got error: %v", err)
		}
	})
}
