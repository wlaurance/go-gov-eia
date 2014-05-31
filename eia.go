package eia

import (
	"net/http"
)

func Client(apiKey string, client *http.Client) EIAClient {
	return EIAClient{apiKey, client}
}
