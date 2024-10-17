// Package fieldset provides request configuration for methods of the marksman client.
package fieldset

// RequestOptions contains possible options for requests made to the /fieldsets endpoints.
type RequestOptions struct {
	Name string `json:"name,omitempty" url:"-"`
}

// RequestOption is used to configure a [RequestOptions].
type RequestOption func(*RequestOptions)

// Name sets the name of a [snipeit.Fieldset].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}
