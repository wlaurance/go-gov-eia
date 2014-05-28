package eia

import (
	"net/http"
	"os"
	"testing"
)

func getApiKey() string {
	key := os.Getenv("API_KEY")
	if key == "" {
		panic("Please set API_KEY in the shell environment.")
	}
	return key
}

func getClient() EIAClient {
	return Client(getApiKey(), &http.Client{})
}

func TestCategories(t *testing.T) {
	client := getClient()
	categories, err := client.Categories()
	if err != nil {
		t.Errorf("Categories error %s", err.Error())
	} else {
		el := 1
		if len(categories) < el {
			t.Errorf("Categories had length %d, not %d", len(categories), el)
		}
	}
}
