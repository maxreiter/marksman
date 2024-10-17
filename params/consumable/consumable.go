package consumable

import (
	"io"
	"net/url"

	params "github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type RequestOptions struct {
	*params.BaseResolver

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

func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

func OrderNumber(number string) RequestOption {
	return func(ro *RequestOptions) {
		ro.OrderNumber = number
	}
}

func ManufacturerID(id snipeit.ManufacturerID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManufacturerID = id
	}
}

func Qty(qty int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Qty = qty
	}
}

func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

func Requestable() RequestOption {
	return func(ro *RequestOptions) {
		ro.Requestable = true
	}
}

func PurchaseDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseDate = date
	}
}

func MinAmt(amt int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.MinAmt = amt
	}
}

func ModelNumber(num string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ModelNumber = num
	}
}

func ItemNo(no string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ItemNo = no
	}
}

func AssignedTo(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedTo = id
	}
}
