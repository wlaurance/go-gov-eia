package eia

type EIACategory struct {
	CategoryId int    `json:"category_id"`
	Name       string `json:"category_id"`
}

type EIATLDCategory struct {
	ChildCategories []EIACategory `json:"childcategories"`
	EIACategory
}

type EIACategoryRequest struct {
	EIARequest `json:"request"`
	CategoryId int `json:"category_id"`
}

type EIACategoryResponse struct {
	Category EIATLDCategory `json:"category"`
	Request  EIACategoryRequest
}
