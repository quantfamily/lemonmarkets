package lemonmarkets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

func getErrorResponse(resp *http.Response) error {
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	lemonError := new(LemonError)
	err = json.Unmarshal(responseBody, lemonError)
	if err != nil {
		fmt.Println("Unmarshal error: ", string(responseBody))
		return err
	}
	return lemonError
}

/*
LemonError is a type of error parsed from the error- response given during 400- error.
Read more at: https://docs.lemon.markets/error-handling
*/
type LemonError struct {
	Time    time.Time `json:"time"`
	Mode    string    `json:"mode"`
	Status  string    `json:"status"`
	Code    string    `json:"error_code"`
	Message string    `json:"error_message"`
}

/*
Error will return the error- message coming from LemonMarkets backend
*/
func (e LemonError) Error() string {
	return e.Message
}

/*
Environment reference a type of LemonMarkets Environment to the corresponding base url
*/
type Environment string

const (
	// PAPER trading without using real cash
	PAPER Environment = "https://paper-trading.lemon.markets/v1"
	// LIVE trading using real cash
	LIVE Environment = "https://trading.lemon.markets/v1"
	// DATA to request market data to do analysis on
	DATA Environment = "https://data.lemon.markets/v1"
)

/*
NewClient initializes a new client towards LemonMarkets.
The client object will be used when using API calls to the different endpoints
*/
func NewClient(env Environment, apikey string) Client {
	lc := LemonClient{Environment: env, APIKey: apikey}
	return &lc
}

/*
Client is interface to do operations towards backend service
*/
type Client interface {
	Do(method string, endpoint string, query interface{}, data []byte) ([]byte, error)
}

/*
LemonClient holding the Base Path of http address as Environment as well as corresponding API Key
*/
type LemonClient struct {
	Environment Environment
	APIKey      string
}

/*
Do preforms request towards the backend service.
Method as Restful method (GET, POST, etc)
Endpoint as where the call should go (OHLC, Account etc)
Q as struct holding query- parameters that should be include, eg for filtering
Data as request body that should be posted
*/
func (c *LemonClient) Do(method string, endpoint string, q interface{}, data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.Environment, endpoint)
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
	responseBody, err := io.ReadAll(resp.Body)

	return responseBody, err
}

/*
Reply is basic values that are included from the backend to know if the request is good or not
*/
type Reply struct {
	Time   time.Time `json:"time"`
	Status string    `json:"status"`
	Mode   string    `json:"mode"`
}

/*
ListReply includes the information that can be used to get further iterations of the values in a bigger collection
*/
type ListReply struct {
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Total    int    `json:"total"`
	Page     int    `json:"page"`
	Pages    int    `json:"pages"`
	Reply
}

/*
Next will call the backend for the next list of items of a particular collection and Update its structure
*/
func (lr *ListReply) QueryNext(client Client, s interface{}) error {
	if len(lr.Next) == 0 {
		return fmt.Errorf("end of list")
	}
	splitted := strings.Split(lr.Next, "/v1/")
	if len(splitted) != 2 {
		return fmt.Errorf("url is not correct")
	}
	responseData, err := client.Do("GET", splitted[1], nil, nil)
	if err != nil {
		return err
	}
	return json.Unmarshal(responseData, s)
}
