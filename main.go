package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

func getErrorResponse(resp *http.Response) error {
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	lemonError := new(LemonError)
	err = json.Unmarshal(responseBody, lemonError)
	if err != nil {
		fmt.Println("Unmarshal errro: ", string(responseBody))
		return err
	}
	return lemonError
}

type LemonError struct {
	Time    time.Time `json:"time"`
	Mode    string    `json:"mode"`
	Status  string    `json:"status"`
	Code    string    `json:"error_code"`
	Message string    `json:"error_message"`
}

func (e LemonError) Error() string {
	return e.Message
}

type Environment string

const (
	PAPER Environment = "https://paper-trading.lemon.markets/v1"
	LIVE  Environment = "https://trading.lemon.markets/v1"
	DATA  Environment = "https://data.lemon.markets/v1"
)

func NewClient(env Environment, apiky string) Client {
	lc := LemonClient{Environment: env, ApiKey: apiky}
	return &lc
}

type Client interface {
	Do(string, string, interface{}, []byte) ([]byte, error)
}

type LemonClient struct {
	Environment Environment
	ApiKey      string
}

func (c *LemonClient) Do(method string, endpoint string, q interface{}, data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.Environment, endpoint)
	if q != nil {
		queryString, err := query.Values(q)
		if err != nil {
			return nil, err
		}
		url = fmt.Sprintf("%s?%s", url, queryString)
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))

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
	responseBody, err := ioutil.ReadAll(resp.Body)

	return responseBody, err
}

type Reply struct {
	Time   time.Time `json:"time"`
	Status string    `json:"status"`
	Mode   string    `json:"mode"`
}

type ListReply struct {
	previous string `json:"previous"`
	next     string `json:"next"`
	Total    int    `json:"total"`
	Page     int    `json:"page"`
	Pages    int    `json:"pages"`
	Reply
}

func (lr *ListReply) Next(client Client) error {
	if len(lr.next) == 0 {
		return fmt.Errorf("end of list")
	}
	splitted := strings.Split(lr.next, "/v1/")
	if len(splitted) != 2 {
		return fmt.Errorf("url is not correct")
	}
	responseData, err := client.Do("GET", splitted[1], nil, nil)
	if err != nil {
		return err
	}
	return json.Unmarshal(responseData, lr)
}
