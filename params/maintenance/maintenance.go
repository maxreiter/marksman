// Package maintenance provides request configuration for methods of the [marksman.Client].
package maintenance

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// RequestOptions contains possible options for requests made to the /maintenances endpoints.
type RequestOptions struct {
	*params.Resolver

	// Query params
	Limit  int32            `url:"limit,omitempty" json:"-"`
	Offset int32            `url:"offset,omitempty" json:"-"`
	Search string           `url:"search,omitempty" json:"-"`
	Sort   params.SortType  `url:"sort,omitempty" json:"-"`
	Order  params.OrderType `url:"order,omitempty" json:"-"`

	// Compound params
	AssetID snipeit.AssetID `url:"asset_id,omitempty" json:"asset_id,omitempty"`

	// Body params
	Title                string                       `json:"title,omitempty" url:"-"`
	SupplierID           snipeit.SupplierID           `json:"supplier_id,omitempty" url:"-"`
	IsWarranty           bool                         `json:"is_warranty,omitempty" url:"-"`
	Cost                 float64                      `json:"cost,omitempty" url:"-"`
	Notes                string                       `json:"notes,omitempty" url:"-"`
	AssetMaintenanceType snipeit.AssetMaintenanceType `json:"asset_maintenance_type,omitempty" url:"-"`
	StartDate            string                       `json:"start_date,omitempty" url:"-"`
	CompletionDate       string                       `json:"completion_date,omitempty" url:"-"`
}

// Query marshals the [RequestOptions] as a [url.Query].
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

// AssetID sets the [snipeit.AssetID] of a [snipeit.Maintenance].
func AssetID(id snipeit.AssetID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssetID = id
	}
}

// Title sets the title of a [snipeit.Maintenance].
func Title(title string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Title = title
	}
}

// SupplierID sets the [snipeit.SupplierID] of a [snipeit.Maintenance].
func SupplierID(id snipeit.SupplierID) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupplierID = id
	}
}

// IsWarranty marks a [snipeit.Maintenance] as under warranty.
func IsWarranty() RequestOption {
	return func(ro *RequestOptions) {
		ro.IsWarranty = true
	}
}

// Cost sets the cost of a [snipeit.Maintenance].
func Cost(cost float64) RequestOption {
	return func(ro *RequestOptions) {
		ro.Cost = cost
	}
}

// Notes sets the notes of a [snipeit.Maintenance].
func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

// MaintenanceType sets the type of a [snipeit.Maintenance].
func MaintenanceType(maintenanceType snipeit.AssetMaintenanceType) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssetMaintenanceType = maintenanceType
	}
}

// StartDate sets the start date of a [snipeit.Maintenance].
func StartDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.StartDate = date
	}
}

// CompletionDate sets the competion date of a [snipeit.Maintenance].
func CompletionDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompletionDate = date
	}
}
