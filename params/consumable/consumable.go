// Package consumable provides request configuration for methods of the [marksman.Client].
package consumable

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// RequestOptions contains possible options for requests made to the /consumables endpoints.
type RequestOptions struct {
	*params.Resolver

	// Query params
	Limit     int32             `url:"limit,omitempty" json:"-"`
	Offset    int32             `url:"offset,omitempty" json:"-"`
	Search    string            `url:"search,omitempty" json:"-"`
	Sort      params.SortType   `url:"sort,omitempty" json:"-"`
	Order     params.OrderType  `url:"order,omitempty" json:"-"`
	Expand    string            `url:"expand,omitempty" json:"-"`
	CompanyID snipeit.CompanyID `url:"company_id,omitempty" json:"-"`

	// Compound params
	Name           string                 `url:"name,omitempty" json:"name,omitempty"`
	CategoryID     snipeit.CategoryID     `url:"category_id,omitempty" url:"category_id,omitempty"`
	OrderNumber    string                 `url:"order_number,omitempty" url:"order_number,omitempty"`
	ManufacturerID snipeit.ManufacturerID `url:"manufacturer_id,omitempty" url:"manufacturer_id,omitempty"`

	// Body params
	Qty          int32              `json:"qty,omitempty" url:"-"`
	LocationID   snipeit.LocationID `json:"location_id,omitempty" url:"-"`
	Requestable  bool               `json:"requestable,omitempty" url:"-"`
	PurchaseDate string             `json:"purchase_date,omitempty" url:"-"`
	MinAmt       int32              `json:"min_amt,omitempty" url:"-"`
	ModelNumber  string             `json:"model_number,omitempty" url:"-"`
	ItemNo       string             `json:"item_no,omitempty" url:"-"`
	AssignedTo   snipeit.UserID     `json:"assigned_to,omitempty" url:"-"`
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

// Expand toggles inclusion of detailed information on a [snipeit.Consumable].
func Expand() RequestOption {
	return func(ro *RequestOptions) {
		ro.Expand = "true"
	}
}

// CompanyID sets the [snipeit.CompanyID] of a [snipeit.Consumable].
func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

// Name sets the name of a [snipeit.Consumable].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// CategoryID sets the [snipeit.CategoryID] of a [snipeit.Consumable].
func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

// OrderNumber sets the order number of a [snipeit.Consumable].
func OrderNumber(number string) RequestOption {
	return func(ro *RequestOptions) {
		ro.OrderNumber = number
	}
}

// ManufacturerID sets the [snipeit.ManufacturerID] of a [snipeit.Consumable].
func ManufacturerID(id snipeit.ManufacturerID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManufacturerID = id
	}
}

// Qty sets the quantity of [snipeit.Consumable].
func Qty(qty int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Qty = qty
	}
}

// LocationID sets the [snipeit.LocationID] of a [snipeit.Consumable].
func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

// Requestable marks a [snipeit.Consumable] as requestable.
func Requestable() RequestOption {
	return func(ro *RequestOptions) {
		ro.Requestable = true
	}
}

// PurchaseDate sets the date a [snipeit.Consumable] was purchased.
func PurchaseDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseDate = date
	}
}

// MinAmt sets the minimum amount of a [snipeit.Consumable].
func MinAmt(amt int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.MinAmt = amt
	}
}

// ModelNumber sets the model number of a [snipeit.Consumable].
func ModelNumber(num string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ModelNumber = num
	}
}

// ItemNo sets the item number of a [snipeit.Consumable].
func ItemNo(no string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ItemNo = no
	}
}

// AssignedTo sets the [snipeit.User] a [snipeit.Consumable] is assigned to.
func AssignedTo(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedTo = id
	}
}
