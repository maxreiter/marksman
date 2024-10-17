// Package company provides request configuration for methods of the marksman client.
package company

// RequestOptions contains possible options for requests made to the /companies endpoints.
type RequestOptions struct {
	Name string `url:"name,omitempty" json:"name,omitempty"`
}

// RequestOption is used to configure a [RequestOptions].
type RequestOption func(*RequestOptions)

// Name sets the name of a [snipeit.Company].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}
