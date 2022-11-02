package streaming

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/quantfamily/lemonmarkets/client/helpers"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	rawFileBytes := helpers.ParseFile(t, "get_token.json")

	t.Run("fail to get response", func(t *testing.T) {
		expectedErr := fmt.Errorf("Unexpected status code: 400")
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "", 400)
		}))
		defer server.Close()
		client := StreamingClient{apiKey: "", baseURL: server.URL}
		token := client.GetToken()
		assert.NotNil(t, token.Error)
		assert.Equal(t, expectedErr, token.Error)
	})
	t.Run("Fail to decode results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `really odd response`)
		}))
		defer server.Close()
		client := StreamingClient{apiKey: "", baseURL: server.URL}
		token := client.GetToken()
		assert.NotNil(t, token.Error)
		assert.ObjectsAreEqual(&json.SyntaxError{}, token.Error)
	})
	t.Run("Successful test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(rawFileBytes))
		}))
		defer server.Close()
		client := StreamingClient{apiKey: "", baseURL: server.URL}
		token := client.GetToken()
		assert.Nil(t, token.Error)
		assert.Equal(t, int64(1655856000084), token.Data.ExpiresAt)
		assert.Equal(t, "bsX3hsLNcjrGaaIc", token.Data.Token)
		assert.Equal(t, "usr_qyJD", token.Data.UserID)

	})
}

func IntegrationClient(t *testing.T) *StreamingClient {
	t.Helper()
	apiKey, isSet := os.LookupEnv("STREAMING_API_KEY")
	if !isSet {
		t.Skip("missing environment variable STREAMING_API_KEY")
	}
	return NewClient(apiKey)
}

func TestTokenIntegration(t *testing.T) {
	client := IntegrationClient(t)

	token := client.GetToken()
	assert.Nil(t, token.Error)
	assert.NotEmpty(t, token.Data.ExpiresAt)
	assert.NotEmpty(t, token.Data.UserID)
	assert.NotEmpty(t, token.Data.Token)
}
