package eia

import (
	"net/http"
	"net/url"
)

const base = "http://api.eia.gov/"

type EIAClient struct {
	apiKey     string
	httpClient *http.Client
}

func (e *EIAClient) makeRequest(setName string, qs url.Values) (resp *http.Response, err error) {
	qs.Set("api_key", e.apiKey)
	url := base + setName + "/?" + qs.Encode()
	resp, err = http.Get(url)
	return
}

func (e *EIAClient) Categories() (cats []EIACategory, err error) {
	values := url.Values{}
	values.Set("category_id", "371")
	resp, err := e.makeRequest("category", values)
	defer resp.Body.Close()
	return
}
