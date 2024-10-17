// Package component provides request configuration for methods of the [marksman.Client].
package component

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// RequestOptions contains possible options for requests made to the /components endpoints.
type RequestOptions struct {
	*params.Resolver

	// Query params
	Limit  int32            `url:"limit,omitempty" json:"-"`
	Offset int32            `url:"offset,omitempty" json:"-"`
	Search string           `url:"search,omitempty" json:"-"`
	Sort   params.SortType  `url:"sort,omitempty" json:"-"`
	Order  params.OrderType `url:"order,omitempty" json:"-"`
	Expand string           `url:"expand,omitempty" json:"-"`

	// Compound params
	OrderNumber string `json:"order_number,omitempty" url:"order_number,omitempty"`

	// Body params
	Name         string             `json:"name,omitempty" url:"-"`
	Qty          int32              `json:"qty,omitempty" url:"-"`
	CategoryID   snipeit.CategoryID `json:"category_id,omitempty" url:"-"`
	LocationID   snipeit.LocationID `json:"location_id,omitempty" url:"-"`
	CompanyID    snipeit.CompanyID  `json:"company_id,omitempty" url:"-"`
	PurchaseDate string             `json:"purchase_date,omitempty" url:"-"`
	PurchaseCost float64            `json:"purchase_cost,omitempty" url:"-"`
	MinAmt       int32              `json:"min_amt,omitempty" url:"-"`
	Serial       string             `json:"serial,omitempty" url:"-"`
	AssignedTo   snipeit.UserID     `json:"assigned_to,omitempty" url:"-"`
	AssignedQty  int32              `json:"assigned_qty,omitempty" url:"-"`
	CheckinQty   int32              `json:"checkin_qty,omitempty" url:"-"`
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

// Expand toggles inclusion of detailed information on components, etc.
func Expand() RequestOption {
	return func(ro *RequestOptions) {
		ro.Expand = "true"
	}
}

// OrderNumber sets the order number of a [snipeit.Component].
func OrderNumber(num string) RequestOption {
	return func(ro *RequestOptions) {
		ro.OrderNumber = num
	}
}

// Name sets the name of a [snipeit.Component].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// Qty sets the quantity of a [snipeit.Component].
func Qty(qty int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Qty = qty
	}
}

// CategoryID sets the [snipeit.CategoryID] of a [snipeit.Component].
func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

// LocationID sets the [snipeit.LocationID] of a [snipeit.Component].
func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

// CompanyID sets the [snipeit.CompanyID] of a [snipeit.Component].
func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

// PurchaseDate sets the date of purchase of a [snipeit.Component].
func PurchaseDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseDate = date
	}
}

// PurchaseCost sets the cost of a [snipeit.Component].
func PurchaseCost(cost float64) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseCost = cost
	}
}

// MinAmt sets the minimum amount of a [snipeit.Component].
func MinAmt(amt int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.MinAmt = amt
	}
}

// Serial sets the serial of a [snipeit.Component].
func Serial(serial string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Serial = serial
	}
}

// AssignedTo sets the assigned [snipeit.User] of a [snipeit.Component].
func AssignedTo(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedTo = id
	}
}

// AssignedQty sets the assigned quantity of a [snipeit.Component].
func AssignedQty(qty int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedQty = qty
	}
}

// CheckinQty sets the quanity of a [snipeit.Component] upon checkin.
func CheckinQty(qty int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.CheckinQty = qty
	}
}
