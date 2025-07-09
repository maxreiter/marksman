package request

type FetchCompanies struct {
	Name string `url:"name,omitempty"`
}

type CreateCompany struct {
	Name string `json:"name"`
}

type UpdateCompany struct {
	Name string `json:"name"`
}

type PartiallyUpdateCompany struct {
	Name string `json:"name"`
}
