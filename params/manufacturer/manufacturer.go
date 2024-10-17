// Package manufacturer provides request configuration for methods of the marksman client.
package manufacturer

// RequestOptions contains possible options for requests made to the /manufacturers endpoints.
type RequestOptions struct {

	// Query params
	URL          string `url:"url,omitempty"`
	SupportURL   string `url:"support_url,omitempty"`
	SupportPhone string `url:"support_phone,omitempty"`
	SupportEmail string `url:"support_email,omitempty"`

	// Compound params
	Name string `url:"name,omitempty" json:"name,omitempty"`
}

// RequestOption is used to configure a [RequestOptions].
type RequestOption func(*RequestOptions)

// Name sets the name of a [snipeit.Manufacturer].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// URL sets the URL of a [snipeit.Manufacturer].
func URL(baseURL string) RequestOption {
	return func(ro *RequestOptions) {
		ro.URL = baseURL
	}
}

// SupportURL sets the support URL of a [snipeit.Manufacturer].
func SupportURL(supportURL string) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupportURL = supportURL
	}
}

// SupportPhone sets the support phone number of a [snipeit.Manufacturer].
func SupportPhone(phone string) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupportPhone = phone
	}
}

// SupportEmail sets the support email address of a [snipeit.Manufacturer].
func SupportEmail(email string) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupportEmail = email
	}
}
