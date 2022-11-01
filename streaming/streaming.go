package streaming

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BASEURL = "https://realtime.lemon.markets/v1"

// DataTypes
type DataTypes interface {
	AuthenticationToken
}

// Item
type Item[data DataTypes, err error] struct {
	Data  data
	Error err
}

type StreamingClient struct {
	apiKey  string
	baseURL string
}

func (sc *StreamingClient) GetToken() *Item[AuthenticationToken, error] {
	resp := &http.Response{}
	token := &Item[AuthenticationToken, error]{}

	url := fmt.Sprintf("%s/%s", sc.baseURL, "auth")
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		token.Error = err
		return token
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sc.apiKey))

	client := http.Client{}
	resp, token.Error = client.Do(request)
	if err != nil {
		token.Error = err
		return token
	}

	if resp.StatusCode != 200 {
		token.Error = fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
		return token
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	token.Error = json.Unmarshal(data, &token.Data)
	return token
}

type AuthenticationToken struct {
	Token     string `json:"token"`
	UserID    string `json:"user_id"`
	ExpiresAt int64  `json:"expires_at"`
}

func NewClient(APIKey string) *StreamingClient {
	sc := StreamingClient{apiKey: APIKey, baseURL: BASEURL}
	return &sc
}
