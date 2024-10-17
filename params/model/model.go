// Package model provides request configuration for methods of the [marksman.Client].
package model

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// RequestOptions contains possible options for requests made to the /models endpoints.
type RequestOptions struct {
	*params.Resolver

	Limit  int32            `url:"limit,omitempty" json:"-"`
	Offset int32            `url:"offset,omitempty" json:"-"`
	Search string           `url:"search,omitempty" json:"-"`
	Sort   params.SortType  `url:"sort,omitempty" json:"-"`
	Order  params.OrderType `url:"order,omitempty" json:"-"`

	Name           string                 `json:"name,omitempty" url:"-"`
	ModelNumber    string                 `json:"model_number,omitempty" url:"-"`
	CategoryID     snipeit.CategoryID     `json:"category_id,omitempty" url:"-"`
	ManufacturerID snipeit.ManufacturerID `json:"manufacturer_id,omitempty" url:"-"`
	EOL            int32                  `json:"eol,omitempty" url:"-"`
	FieldsetID     snipeit.FieldsetID     `json:"fieldset_id,omitempty" url:"-"`
	DepreciationID snipeit.DepreciationID `json:"depreciation_id,omitempty" url:"-"`
	Notes          string                 `json:"notes,omitempty" url:"-"`
	Requestable    bool                   `json:"requestable,omitempty" url:"-"`
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

// Offset sets the pagination offset of a request.
func Offset(offset int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Offset = offset
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

// Name sets the name of a [snipeit.Model].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// ModelNumber sets the model number of a [snipeit.Model].
func ModelNumber(num string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ModelNumber = num
	}
}

// CategoryID sets the [snipeit.CategoryID] of a [snipeit.Model].
func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

// ManufacturerID sets the [snipeit.ManufacturerID] of a [snipeit.Model].
func ManufacturerID(id snipeit.ManufacturerID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManufacturerID = id
	}
}

// EOL sets the number of months until a [snipeit.Model] reaches the end of its life.
func EOL(eol int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.EOL = eol
	}
}

// FieldsetID sets the [snipeit.FieldsetID] of a [snipeit.Model].
func FieldsetID(id snipeit.FieldsetID) RequestOption {
	return func(ro *RequestOptions) {
		ro.FieldsetID = id
	}
}

// DepreciationID sets the [snipeit.DepreciationID] of a [snipeit.Model].
func DepreciationID(id snipeit.DepreciationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.DepreciationID = id
	}
}

// Notes sets the notes of a [snipeit.Model].
func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

// Requestable marks a [snipeit.Model] as requestable.
func Requestable() RequestOption {
	return func(ro *RequestOptions) {
		ro.Requestable = true
	}
}
