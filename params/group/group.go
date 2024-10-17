// Package group provides request configuration for methods of the marksman client.
package group

// RequestOptions contains possible options for requests made to the /groups endpoints.
type RequestOptions struct {

	// Compound params
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Body params
	Permissions string `json:"permissions,omitempty"`
}

// RequestOption is used to configure a [RequestOptions].
type RequestOption func(*RequestOptions)

// Name sets the name of a [snipeit.Group].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// Permissions sets the permissions of a [snipeit.Group].
func Permissions(permissions string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Permissions = permissions
	}
}
