package eia

import (
	"encoding/json"
	"io/ioutil"
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

func (e *EIAClient) Categories() (EIATLDCategory, error) {
	return e.CategoriesById("371")
}

func (e *EIAClient) CategoriesById(id string) (cats EIATLDCategory, err error) {
	values := url.Values{}
	values.Set("category_id", id)
	resp, err := e.makeRequest("category", values)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r EIACategoryResponse
	json.Unmarshal(body, &r)
	cats = r.Category
	return
}
