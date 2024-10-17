package supplier

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
)

type RequestOptions struct {
	*params.BaseResolver

	Name     string `url:"name,omitempty" json:"-"`
	Address  string `url:"address,omitempty" json:"-"`
	Address2 string `url:"address2,omitempty" json:"-"`
	City     string `url:"city,omitempty" json:"-"`
	Zip      string `url:"zip,omitempty" json:"-"`
	Country  string `url:"country,omitempty" json:"-"`
	Fax      string `url:"fax,omitempty" json:"-"`
	Email    string `url:"email,omitempty" json:"-"`
	URL      string `url:"url,omitempty" json:"-"`
	Notes    string `url:"notes,omitempty" json:"-"`
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

func Address(address string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Address = address
	}
}

func Address2(address string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Address2 = address
	}
}

func City(city string) RequestOption {
	return func(ro *RequestOptions) {
		ro.City = city
	}
}

func Zip(zip string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Zip = zip
	}
}

func Country(country string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Country = country
	}
}

func Fax(fax string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Fax = fax
	}
}

func Email(email string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Email = email
	}
}

func URL(baseURL string) RequestOption {
	return func(ro *RequestOptions) {
		ro.URL = baseURL
	}
}

func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}
