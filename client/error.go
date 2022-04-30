package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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
