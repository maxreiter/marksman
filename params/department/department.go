// Package department provides request configuration for methods of the marksman client.
package department

import (
	"github.com/maxreiter/marksman/snipeit"
)

// RequestOptions contains options for requests made to the /departments endpoints.
type RequestOptions struct {

	// Query params
	CompanyID  snipeit.CompanyID  `url:"company_id,omitempty"`
	ManagerID  snipeit.UserID     `url:"manager_id,omitempty"`
	LocationID snipeit.LocationID `url:"location_id,omitempty"`

	// Compound params
	Name string `url:"name,omitempty" json:"name,omitempty"`
}

// RequestOption is used to configure a [RequestOptions].
type RequestOption func(*RequestOptions)

// CompanyID sets the [snipeit.CompanyID] of a [snipeit.Department].
func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

// ManagerID sets the manager ID of a [snipeit.Department].
func ManagerID(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManagerID = id
	}
}

// LocationID sets the [snipeit.LocationID] of a [snipeit.Department].
func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

// Name sets the name of a [snipeit.Department].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}
