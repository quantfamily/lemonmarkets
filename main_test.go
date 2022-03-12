package main

import (
	"io/ioutil"
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

func GetClient(t *testing.T, env Envrionment) Client {
	t.Helper()
	apiKey, isSet := os.LookupEnv("LEMON_API_KEY")
	if !isSet {
		t.Skip("missing environment variable LEMON_API_KEY")
	}
	c := LemonClient{Envrionment: env, ApiKey: apiKey}
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
