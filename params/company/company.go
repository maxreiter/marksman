package company

import (
	"io"
	"net/url"

	params "github.com/maxreiter/marksman/params"
)

type RequestOptions struct {
	*params.BaseResolver

	Name string `url:"name,omitempty" json:"name,omitempty"`
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
