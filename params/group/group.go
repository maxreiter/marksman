package group

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
)

type RequestOptions struct {
	*params.BaseResolver

	// Compound params
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Body params
	Permissions string `json:"permissions,omitempty"`
}

func (ro *RequestOptions) Values() (url.Values, error) {
	return ro.BaseResolver.Values()
}

func (ro *RequestOptions) Marshal() (io.Reader, error) {
	return ro.BaseResolver.Marshal()
}

type RequestOption func(*RequestOptions)

func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

func Permissions(permissions string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Permissions = permissions
	}
}
