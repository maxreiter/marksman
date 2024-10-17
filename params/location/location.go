package location

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type RequestOptions struct {
	*params.BaseResolver

	// Query params
	Limit  int32            `url:"limit,omitempty" json:"-"`
	Search string           `url:"search,omitempty" json:"-"`
	Sort   params.SortType  `url:"sort,omitempty" json:"-"`
	Order  params.OrderType `url:"order,omitempty" json:"-"`

	// Compound params
	Name     string `url:"name,omitempty" json:"name,omitempty"`
	Address  string `url:"address,omitempty" json:"address,omitempty"`
	Address2 string `url:"address2,omitempty" json:"address2,omitempty"`
	City     string `url:"city,omitempty" json:"city,omitempty"`
	Zip      string `url:"zip,omitempty" json:"zip,omitempty"`
	Country  string `url:"country,omitempty" json:"country,omitempty"`

	// Body params
	State     string             `json:"state,omitempty" url:"-"`
	LDAPOU    string             `json:"ldap_ou,omitempty" url:"-"`
	ParentID  snipeit.LocationID `json:"parent_id,omitempty" url:"-"`
	Currency  string             `json:"currency,omitempty" url:"-"`
	ManagerID snipeit.UserID     `json:"manager_id,omitempty" url:"-"`
}

func (ro *RequestOptions) Values() (url.Values, error) {
	return ro.BaseResolver.Values()
}

func (ro *RequestOptions) Marshal() (io.Reader, error) {
	return ro.BaseResolver.Marshal()
}

type RequestOption func(*RequestOptions)

func Limit(limit int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Limit = limit
	}
}

func Search(search string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Search = search
	}
}

func Sort(sort params.SortType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Sort = sort
	}
}

func Order(order params.OrderType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Order = order
	}
}

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

func State(state string) RequestOption {
	return func(ro *RequestOptions) {
		ro.State = state
	}
}

func LDAPOU(ou string) RequestOption {
	return func(ro *RequestOptions) {
		ro.LDAPOU = ou
	}
}

func ParentID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ParentID = id
	}
}

func Currency(currency string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Currency = currency
	}
}

func ManagerID(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManagerID = id
	}
}
