package request

type FetchSuppliers struct {
	Name     string `url:"name,omitempty"`
	Address  string `url:"address,omitempty"`
	Address2 string `url:"address2,omitempty"`
	City     string `url:"city,omitempty"`
	Zip      string `url:"zip,omitempty"`
	Country  string `url:"country,omitempty"`
	Fax      string `url:"fax,omitempty"`
	Email    string `url:"email,omitempty"`
	URL      string `url:"url,omitempty"`
	Notes    string `url:"notes,omitempty"`
}
