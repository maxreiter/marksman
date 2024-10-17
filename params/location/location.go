// Package location provides request configuration for methods of the marksman client.
package location

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// RequestOptions contains possible options for requests made to the /locations endpoints.
type RequestOptions struct {
	*params.Resolver

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

// Limit sets the return limit of a request.
func Limit(limit int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Limit = limit
	}
}

// Search sets the search string of a request.
func Search(search string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Search = search
	}
}

// Sort sets the return sort of a request.
func Sort(sort params.SortType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Sort = sort
	}
}

// Order sets the return order of a request.
func Order(order params.OrderType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Order = order
	}
}

// Name sets the name of a [snipeit.Location].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// Address sets the address of a [snipeit.Location].
func Address(address string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Address = address
	}
}

// Address2 sets the alternate address of a [snipeit.Location].
func Address2(address string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Address2 = address
	}
}

// City sets the city of a [snipeit.Location].
func City(city string) RequestOption {
	return func(ro *RequestOptions) {
		ro.City = city
	}
}

// Zip sets the zip code of a [snipeit.Location].
func Zip(zip string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Zip = zip
	}
}

// Country sets the country of a [snipeit.Location].
func Country(country string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Country = country
	}
}

// State sets the state of a [snipeit.Location].
func State(state string) RequestOption {
	return func(ro *RequestOptions) {
		ro.State = state
	}
}

// LDAPOU sets the LDAP Organizational Unit of a [snipeit.Location].
func LDAPOU(ou string) RequestOption {
	return func(ro *RequestOptions) {
		ro.LDAPOU = ou
	}
}

// ParentID sets the parent ID of a [snipeit.Location].
func ParentID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ParentID = id
	}
}

// Currency sets the currency of a [snipeit.Location].
func Currency(currency string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Currency = currency
	}
}

// ManagerID sets the manager ID of a [snipeit.Location].
func ManagerID(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManagerID = id
	}
}
