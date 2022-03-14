package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type MockedClient struct {
	CalledMethod   string
	CalledEndpoint string
	CalledQuery    interface{}
	CalledData     []byte
	ReturnData     []byte
	ReturnError    error
}

func (mc *MockedClient) Do(method string, endpoint string, q interface{}, data []byte) ([]byte, error) {
	mc.CalledMethod = method
	mc.CalledEndpoint = endpoint
	mc.CalledQuery = q
	mc.CalledData = data
	return mc.ReturnData, mc.ReturnError
}

func GetMockedClient(t *testing.T) *MockedClient {
	t.Helper()
	return new(MockedClient)
}

func GetClient(t *testing.T, env Environment) Client {
	t.Helper()
	apiKey, isSet := os.LookupEnv("LEMON_API_KEY")
	if !isSet {
		t.Skip("missing environment variable LEMON_API_KEY")
	}
	c := LemonClient{Environment: env, ApiKey: apiKey}
	return &c
}

func ParseFile(t *testing.T, fileName string) []byte {
	t.Helper()
	file, err := os.Open(fileName)
	if err != nil {
		t.Errorf("Error opening file for test: %w", err)
		return nil
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("Error reading bytes from file: %w", err)
		return nil
	}
	return bytes
}

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
