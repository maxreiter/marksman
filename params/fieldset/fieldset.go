// Package fieldset provides request configuration for methods of the [marksman.Client].
package fieldset

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
)

// RequestOptions contains possible options for requests made to the /fieldsets endpoints.
type RequestOptions struct {
	*params.Resolver

	Name string `json:"name,omitempty" url:"-"`
}

// Query marshals the [RequestOptions] as a [url.Values].
func (ro *RequestOptions) Query() (url.Values, error) {
	return ro.Resolver.Query()
}

// JSON encodes the [RequestOptions] as JSON.
func (ro *RequestOptions) JSON() (io.Reader, error) {
	return ro.Resolver.JSON()
}

// RequestOption is used to configure a [RequestOptions].
type RequestOption func(*RequestOptions)

// Name sets the name of a [snipeit.Fieldset].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}
