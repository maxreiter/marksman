package manufacturer

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
)

type RequestOptions struct {
	*params.BaseResolver

	// Query params
	URL          string `url:"url,omitempty"`
	SupportURL   string `url:"support_url,omitempty"`
	SupportPhone string `url:"support_phone,omitempty"`
	SupportEmail string `url:"support_email,omitempty"`

	// Compound params
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

func URL(baseURL string) RequestOption {
	return func(ro *RequestOptions) {
		ro.URL = baseURL
	}
}

func SupportURL(supportURL string) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupportURL = supportURL
	}
}

func SupportPhone(phone string) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupportPhone = phone
	}
}

func SupportEmail(email string) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupportEmail = email
	}
}
