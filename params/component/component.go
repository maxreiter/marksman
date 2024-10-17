package component

import (
	"io"
	"net/url"

	params "github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type RequestOptions struct {
	*params.BaseResolver

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

func Offset(offset int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Offset = offset
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

func Expand() RequestOption {
	return func(ro *RequestOptions) {
		ro.Expand = "true"
	}
}

func OrderNumber(num string) RequestOption {
	return func(ro *RequestOptions) {
		ro.OrderNumber = num
	}
}

func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

func Qty(qty int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Qty = qty
	}
}

func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

func PurchaseDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseDate = date
	}
}

func PurchaseCost(cost float64) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseCost = cost
	}
}

func MinAmt(amt int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.MinAmt = amt
	}
}

func Serial(serial string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Serial = serial
	}
}

func AssignedTo(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedTo = id
	}
}

func AssignedQty(qty int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedQty = qty
	}
}

func CheckinQty(qty int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.CheckinQty = qty
	}
}
