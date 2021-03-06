package eia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const base = "http://api.eia.gov/"

type EIAClient struct {
	ApiKey     string
	HttpClient *http.Client
}

func (e *EIAClient) makeRequest(setName string, qs url.Values) (resp *http.Response, err error) {
	qs.Set("api_key", e.ApiKey)
	url := base + setName + "/?" + qs.Encode()
	resp, err = e.HttpClient.Get(url)
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

func (e *EIAClient) SeriesById(id string) (series []EIASeriesExtended, err error) {
	values := url.Values{}
	values.Set("series_id", id)
	resp, err := e.makeRequest("series", values)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var s EIASeriesResponse
	json.Unmarshal(body, &s)
	series = s.Series
	return
}
