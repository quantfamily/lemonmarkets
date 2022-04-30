package client

import (
	"encoding/json"
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
