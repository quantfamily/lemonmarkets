package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/quantfamily/lemonmarkets/common"
)

type MockedClient struct {
	CalledMethod   string
	CalledEndpoint string
	CalledQuery    interface{}
	CalledData     []byte
	ReturnResponse *common.Response
	ReturnError    error
}

func (mc *MockedClient) Do(method string, endpoint string, query interface{}, data []byte) (*common.Response, error) {
	mc.CalledMethod = method
	mc.CalledEndpoint = endpoint
	mc.CalledQuery = query
	mc.CalledData = data
	return mc.ReturnResponse, mc.ReturnError
}

func GetMockedClient(t *testing.T) *MockedClient {
	t.Helper()
	return new(MockedClient)
}

func ParseFile(t *testing.T, fileName string) []byte {
	t.Helper()
	filePath, _ := filepath.Abs(fmt.Sprintf("test_data/%s", fileName))
	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Error opening file for test: %v", err)
		return nil
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("Error reading bytes from file: %v", err)
		return nil
	}
	return bytes
}

func APIKey(t *testing.T) string {
	t.Helper()
	apiKey, isSet := os.LookupEnv("LEMON_API_KEY")
	if !isSet {
		t.Skip("missing environment variable LEMON_API_KEY")
	}
	return apiKey
}
