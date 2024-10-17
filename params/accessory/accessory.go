// Package accessory provides request configuration for methods of the [marksman.Client].
package accessory

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// RequestOptions contains possible options for requests made to the /accessories endpoints.
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
	OrderNumber string `url:"order_number,omitempty" json:"order_number,omitempty"`

	// Body params
	Name           string                 `json:"name,omitempty" url:"-"`
	Qty            int32                  `json:"qty,omitempty" url:"-"`
	PurchaseCost   float64                `json:"purchase_cost,omitempty" url:"-"`
	PurchaseDate   string                 `json:"purchase_date,omitempty" url:"-"`
	ModelNumber    string                 `json:"model_number,omitempty" url:"-"`
	CategoryID     snipeit.CategoryID     `json:"category_id,omitempty" url:"-"`
	CompanyID      snipeit.CompanyID      `json:"company_id,omitempty" url:"-"`
	LocationID     snipeit.LocationID     `json:"location_id,omitempty" url:"-"`
	ManufacturerID snipeit.ManufacturerID `json:"manufacturer_id,omitempty" url:"-"`
	SupplierID     snipeit.SupplierID     `json:"supplier_id,omitempty" url:"-"`
	AssignedUser   snipeit.UserID         `json:"assigned_user,omitempty" url:"-"`
	Note           string                 `json:"note,omitempty" url:"-"`
	CheckoutQty    int32                  `json:"checkout_qty,omitempty" url:"-"`
}

// Query marshals the [RequestOptions] into a [url.Values].
func (ro *RequestOptions) Query() (url.Values, error) {
	return ro.Resolver.Query()
}

// JSON encodes the [RequestOptions] as JSON.
func (ro *RequestOptions) JSON() (io.Reader, error) {
	return ro.Resolver.JSON()
}

// RequestOption configures a [RequestOptions].
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

// Expand toggles inclusion of detailed information on categories, etc.
func Expand() RequestOption {
	return func(ro *RequestOptions) {
		ro.Expand = "true"
	}
}

// OrderNumber sets the order number of an [snipeit.Accessory].
func OrderNumber(number string) RequestOption {
	return func(ro *RequestOptions) {
		ro.OrderNumber = number
	}
}

// Name sets the name of an [snipeit.Accessory].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// Qty sets the quantity of an [snipeit.Accessory].
func Qty(qty int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Qty = qty
	}
}

// PurchaseCost sets the cost of an [snipeit.Accessory].
func PurchaseCost(cost float64) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseCost = cost
	}
}

// PurchaseDate sets the date of purchase of an [snipeit.Accessory].
func PurchaseDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseDate = date
	}
}

// ModelNumber sets the model number of an [snipeit.Accessory].
func ModelNumber(number string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ModelNumber = number
	}
}

// CategoryID sets the [snipeit.CategoryID] of an [snipeit.Accessory].
func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

// CompanyID sets the [snipeit.CompanyID] of an [snipeit.Accessory].
func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

// LocationID sets the [snipeit.LocationID] of an [snipeit.Accessory].
func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

// Manufacturer sets the [snipeit.ManufacturerID] of an [snipeit.Accessory].
func ManufacturerID(id snipeit.ManufacturerID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManufacturerID = id
	}
}

// SupplierID sets the [snipeit.SupplierID] of an [snipeit.Accessory].
func SupplierID(id snipeit.SupplierID) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupplierID = id
	}
}

// AssignedUser sets the [snipeit.User] an [snipeit.Accessory] is assigned to.
func AssignedUser(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedUser = id
	}
}

// Note sets the note of an [snipeit.Accessory].
func Note(note string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Note = note
	}
}

// CheckoutQty sets the checkout quantity of an [snipeit.Accessory].
func CheckoutQty(qty int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.CheckoutQty = qty
	}
}
