package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

/*
LemonClient holding the Base Path of http address as Environment as well as corresponding API Key
*/
type Backend struct {
	BaseURL string
	APIKey  string
}

/*
Do preforms request towards the backend service.
Method as Restful method (GET, POST, etc)
Endpoint as where the call should go (OHLC, Account etc)
Q as struct holding query- parameters that should be include, eg for filtering
Data as request body that should be posted
*/
func (c *Backend) Do(method string, endpoint string, q interface{}, data []byte) (*Response, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
	if q != nil {
		queryString, err := query.Values(q)
		if err != nil {
			return nil, err
		}
		url = fmt.Sprintf("%s?%s", url, queryString.Encode())
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 400 {
		return nil, getErrorResponse(resp)
	}
	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("unknown http error from backend: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	response := new(Response)
	data, err = io.ReadAll(resp.Body)
	err = json.Unmarshal(data, response)
	return response, err
}
