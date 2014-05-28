package eia

import (
	"errors"
	"net/http"
)

const base = "http://api.eia.gov/"

type EIAClient struct {
	apiKey     string
	httpClient *http.Client
}

func (e *EIAClient) Categories() (cats []EIACategory, err error) {
	cats = make([]EIACategory, 1)
	err = errors.New("BAD CATEGORIES")
	return
}
