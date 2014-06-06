package eia

import (
	"encoding/json"
)

type EIAName struct {
	Name string `json:"name"`
}

type EIACategoryId struct {
	CategoryId int `json:"category_id"`
}

type EIASeriesId struct {
	SeriesId string `json:"series_id"`
}

type EIACategory struct {
	EIACategoryId
	EIAName
}

type EIASeries struct {
	EIASeriesId
	EIAName
	Updated string `json:"updated"`
	F       string `json:"f"`
	Units   string `json:"units"`
}

type EIATLDCategory struct {
	ChildCategories []EIACategory `json:"childcategories"`
	ChildSeries     []EIASeries   `json:"childseries"`
	EIACategory
}

type EIATLDRequest struct {
	Request EIARequest `json:"request"`
}

type EIACategoryRequest struct {
	EIARequest
	EIACategoryId
}

type EIACategoryResponse struct {
	Category EIATLDCategory `json:"category"`
	Request  EIACategoryRequest
}

type EIASeriesRequest struct {
	EIARequest
	EIASeriesId
}

type EIASeriesResponse struct {
	EIASeriesRequest
	Series []EIASeriesExtended `json:"series"`
}

type EIASeriesExtended struct {
	EIASeries
	Data        []EIAPoint `json:"data" datastore:",noindex"`
	Description string     `json:"description"`
	UnitsShort  string     `json:"unitsshort"`
	Geography   string     `json:"geography"`
}

type EIASeriesExtendedPre struct {
	EIASeriesExtended
	Data [][]interface{} `json:"data"`
}

type EIAPoint struct {
	Date  string
	Price float64
}

func (ep *EIAPoint) UnmarshalJSON(b []byte) (err error) {
	var a []interface{}
	err = json.Unmarshal(b, &a)
	if err != nil {
		return
	}
	*ep = EIAPoint{Date: a[0].(string), Price: a[1].(float64)}
	return
}
