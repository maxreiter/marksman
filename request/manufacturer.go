package request

type FetchManufacturers struct {
	Name         string `url:"name,omitempty"`
	URL          string `url:"url,omitemtpy"`
	SupportURL   string `url:"support_url,omitempty"`
	SupportPhone string `url:"support_phone,omitempty"`
	SupportEmail string `url:"support_email,omitempty"`
}

type CreateManufacturer struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitempty"`
	SupportURL   string `json:"support_url,omitempty"`
	SupportPhone string `json:"support_phone,omitempty"`
	SupportEmail string `json:"support_email,omitempty"`
}

type UpdateManufacturer struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitempty"`
	SupportURL   string `json:"support_url,omitempty"`
	SupportPhone string `json:"support_phone,omitempty"`
	SupportEmail string `json:"support_email,omitempty"`
}

type PartiallyUpdateManufacturer struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	SupportURL   string `json:"support_url,omitempty"`
	SupportPhone string `json:"support_phone,omitempty"`
	SupportEmail string `json:"support_email,omitempty"`
}
