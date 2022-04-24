package common

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

/*
Reply is basic values that are included from the backend to know if the request is good or not
*/
type Response struct {
	Time     time.Time       `json:"time"`
	Status   string          `json:"status"`
	Mode     string          `json:"mode"`
	Previous string          `json:"previous"`
	Next     string          `json:"next"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	Pages    int             `json:"pages"`
	Results  json.RawMessage `json:"results"`
}

/*
Next will call the backend for the next list of items of a particular collection and Update its structure
*/
func (lr *Response) QueryNext(client Client, s interface{}) error {
	if len(lr.Next) == 0 {
		return fmt.Errorf("end of list")
	}
	splitted := strings.Split(lr.Next, "/v1/")
	if len(splitted) != 2 {
		return fmt.Errorf("url is not correct")
	}
	response, err := client.Do("GET", splitted[1], nil, nil)
	if err != nil {
		return err
	}
	return json.Unmarshal(response.Results, s)
}
