package eia

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

type EIACategoryRequest struct {
	EIARequest `json:"request"`
	EIACategoryId
}

type EIACategoryResponse struct {
	Category EIATLDCategory `json:"category"`
	Request  EIACategoryRequest
}
