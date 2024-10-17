// Package license provides request configuration for methods of the marksman client.
package license

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// RequestOptions contains possible options for requests made to the /licenses endpoints.
type RequestOptions struct {
	*params.Resolver

	// Query params
	ProductKey     string                 `url:"product_key,omitempty"`
	Limit          int32                  `url:"limit,omitempty" json:"-"`
	Offset         int32                  `url:"offset,omitempty" json:"-"`
	Search         string                 `url:"search,omitempty" json:"-"`
	Sort           params.SortType        `url:"sort,omitempty" json:"-"`
	Order          params.OrderType       `url:"order,omitempty" json:"-"`
	Expand         string                 `url:"expand,omitempty" json:"-"`
	DepreciationID snipeit.DepreciationID `url:"depreciation_id,omitempty"`
	Deleted        string                 `url:"deleted,omitempty"`

	// Compound params
	Name           string                 `url:"name,omitempty" json:"name,omitempty"`
	CategoryID     snipeit.CategoryID     `url:"category_id,omitempty" json:"category_id,omitempty"`
	LicenseEmail   string                 `url:"license_email,omitempty" json:"license_email,omitempty"`
	LicenseName    string                 `url:"license_name,omitempty" json:"license_name,omitempty"`
	Maintained     bool                   `url:"maintained,omitempty" json:"maintained,omitempty"`
	ManufacturerID snipeit.ManufacturerID `url:"manufacturer_id,omitempty" json:"manufacturer_id,omitempty"`
	OrderNumber    string                 `url:"order_number,omitempty" json:"order_number,omitempty"`
	PurchaseOrder  string                 `url:"purchase_order,omitempty" json:"purchase_order,omitempty"`
	SupplierID     snipeit.SupplierID     `url:"supplier_id,omitempty" json:"supplier_id,omitempty"`

	// Body params
	Seats           int32             `json:"seats,omitempty" url:"-"`
	CompanyID       snipeit.CompanyID `json:"company_id,omitempty" url:"-"`
	ExpirationDate  string            `json:"expiration_date"`
	Serial          string            `json:"serial,omitempty" url:"-"`
	Notes           string            `json:"notes,omitempty" url:"-"`
	PurchaseCost    float64           `json:"purchase_cost,omitempty" url:"-"`
	Reassignable    bool              `json:"reassignable,omitempty" url:"-"`
	TerminationDate string            `json:"termination_date,omitempty" url:"-"`
	PurchaseDate    string            `json:"purchase_date,omitempty" url:"-"`
	AssignedTo      snipeit.UserID    `json:"assigned_to,omitempty" url:"-"`
	AssetID         snipeit.AssetID   `json:"asset_id,omitempty" url:"-"`
	Note            string            `json:"note,omitempty" url:"-"`
}

// Query marshals the [RequestOptions] to a [url.Values].
func (ro *RequestOptions) Query() (url.Values, error) {
	return ro.Resolver.Query()
}

// JSON encodes the [RequestOptions] as JSON.
func (ro *RequestOptions) JSON() (io.Reader, error) {
	return ro.Resolver.JSON()
}

// RequestOption is used to configure a [RequestOptions]/
type RequestOption func(*RequestOptions)

// ProducyKey sets the product key of a [snipeit.License].
func ProductKey(key string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ProductKey = key
	}
}

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

// Expand toggles the inclusion of detailed information on a [snipeit.License].
func Expand() RequestOption {
	return func(ro *RequestOptions) {
		ro.Expand = "true"
	}
}

// DepreciationID sets the [snipeit.DepreciationID] of a [snipeit.License].
func DepreciationID(id snipeit.DepreciationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.DepreciationID = id
	}
}

// Deleted toggles including deleted results.
func Deleted() RequestOption {
	return func(ro *RequestOptions) {
		ro.Deleted = "true"
	}
}

// Name sets the name of a [snipeit.License].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// CategoryID sets the [snipeit.CategoryID] of a [snipeit.License].
func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

// LicenseEmail sets the associated email address of a [snipeit.License].
func LicenseEmail(email string) RequestOption {
	return func(ro *RequestOptions) {
		ro.LicenseEmail = email
	}
}

// LicenseName sets the associated name of a [snipeit.License].
func LicenseName(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.LicenseName = name
	}
}

// Maintained marks the [snipeit.License] as maintained.
func Maintained() RequestOption {
	return func(ro *RequestOptions) {
		ro.Maintained = true
	}
}

// ManufacturerID sets the [snipeit.ManufacturerID] of a [snipeit.License].
func ManufacturerID(id snipeit.ManufacturerID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManufacturerID = id
	}
}

// OrderNumber sets the order number of a [snipeit.License].
func OrderNumber(num string) RequestOption {
	return func(ro *RequestOptions) {
		ro.OrderNumber = num
	}
}

// PurchaseOrder sets the purchase order of a [snipeit.License].
func PurchaseOrder(order string) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseOrder = order
	}
}

// SupplierID sets the [snipeit.SupplierID] of a [snipeit.License].
func SupplierID(id snipeit.SupplierID) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupplierID = id
	}
}

// Seats sets the number of seats a [snipeit.License] has.
func Seats(seats int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Seats = seats
	}
}

// CompanyID sets the [snipeit.CompanyID] of a [snipeit.License].
func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

// ExpirationDate sets the date a [snipeit.License] expires.
func ExpirationDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ExpirationDate = date
	}
}

// Serial sets the serial of a [snipeit.License].
func Serial(serial string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Serial = serial
	}
}

// Notes sets the notes of a [snipeit.License].
func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

// PurchaseCost sets the cost of a [snipeit.License].
func PurchaseCost(cost float64) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseCost = cost
	}
}

// Reassignable marks a [snipeit.License] as reassignable.
func Reassignable() RequestOption {
	return func(ro *RequestOptions) {
		ro.Reassignable = true
	}
}

// TermindationDate sets the date of termination of a [snipeit.License].
func TerminationDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.TerminationDate = date
	}
}

// PurchaseDate sets the date of purchase of a [snipeit.License].
func PurchaseDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseDate = date
	}
}

// AssignedTo assigns a [snipeit.User] to a [snipeit.License].
func AssignedTo(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedTo = id
	}
}

// AssetID sets the [snipeit.AssetID] of a [snipeit.License].
func AssetID(id snipeit.AssetID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssetID = id
	}
}

// Note sets the note of a [snipeit.License].
func Note(note string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Note = note
	}
}
