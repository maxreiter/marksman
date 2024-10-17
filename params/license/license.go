package license

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type RequestOptions struct {
	*params.BaseResolver

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

func (ro *RequestOptions) Values() (url.Values, error) {
	return ro.BaseResolver.Values()
}

func (ro *RequestOptions) Marshal() (io.Reader, error) {
	return ro.BaseResolver.Marshal()
}

type RequestOption func(*RequestOptions)

func ProductKey(key string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ProductKey = key
	}
}

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

func DepreciationID(id snipeit.DepreciationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.DepreciationID = id
	}
}

func Deleted() RequestOption {
	return func(ro *RequestOptions) {
		ro.Deleted = "true"
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

func LicenseEmail(email string) RequestOption {
	return func(ro *RequestOptions) {
		ro.LicenseEmail = email
	}
}

func LicenseName(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.LicenseName = name
	}
}

func Maintained() RequestOption {
	return func(ro *RequestOptions) {
		ro.Maintained = true
	}
}

func ManufacturerID(id snipeit.ManufacturerID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManufacturerID = id
	}
}

func OrderNumber(num string) RequestOption {
	return func(ro *RequestOptions) {
		ro.OrderNumber = num
	}
}

func PurchaseOrder(order string) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseOrder = order
	}
}

func SupplierID(id snipeit.SupplierID) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupplierID = id
	}
}

func Seats(seats int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Seats = seats
	}
}

func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

func ExpirationDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ExpirationDate = date
	}
}

func Serial(serial string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Serial = serial
	}
}

func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

func PurchaseCost(cost float64) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseCost = cost
	}
}

func Reassignable() RequestOption {
	return func(ro *RequestOptions) {
		ro.Reassignable = true
	}
}

func TerminationDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.TerminationDate = date
	}
}

func PurchaseDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseDate = date
	}
}

func AssignedTo(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedTo = id
	}
}

func AssetID(id snipeit.AssetID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssetID = id
	}
}

func Note(note string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Note = note
	}
}
