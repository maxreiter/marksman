package department

import (
	"io"
	"net/url"

	params "github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type RequestOptions struct {
	*params.BaseResolver

	// Query params
	CompanyID  snipeit.CompanyID  `url:"company_id,omitempty"`
	ManagerID  snipeit.UserID     `url:"manager_id,omitempty"`
	LocationID snipeit.LocationID `url:"location_id,omitempty"`

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

func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

func ManagerID(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManagerID = id
	}
}

func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}
