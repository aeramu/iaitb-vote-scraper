package entity

type Response struct {
	Meta Meta `json:"meta"`
	Data []Alumnee `json:"data"`
}

type Alumnee struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	KeywordName string
	Generation  string
	Major       string
	Status      string `json:"verificationStatus"`
	Error       string
}

type Meta struct {
	TotalItems int `json:"totalItems"`
}
