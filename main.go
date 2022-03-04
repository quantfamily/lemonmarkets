package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Envrionment string

const (
	PAPER Envrionment = "https://paper-trading.lemon.markets/v1"
	LIVE  Envrionment = "https://trading.lemon.markets/v1"
	DATA  Envrionment = "https://data.lemon.markets/v1"
)

func NewClient(env Envrionment, apiky string) *Client {
	c := Client{Envrionment: env, ApiKey: apiky}
	return &c
}

type Client struct {
	Envrionment Envrionment
	ApiKey      string
}

func (c *Client) Do(method string, endpoint string, data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.Envrionment, endpoint)
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
	Reply
	previous string `json:"previous"`
	next     string `json:"next"`
	Total    int    `json:"total"`
	Page     int    `json:"page"`
	Pages    int    `json:"pages"`
}

func (lr *ListReply) Next(client *Client) error {
	request, err := http.NewRequest("GET", lr.next, nil)
	if err != nil {
		return nil
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.ApiKey))
	hclient := http.Client{}
	resp, err := hclient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(responseBody, lr)
}
