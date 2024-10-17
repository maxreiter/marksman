package asset

import (
	"fmt"
	"io"
	"net/url"

	params "github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type StatusType string

const (
	StatusRTD          StatusType = "rtd"
	StatusDeployed     StatusType = "deployed"
	StatusUndeployable StatusType = "undeployable"
	StatusDeleted      StatusType = "deleted"
	StatusArchived     StatusType = "archived"
	StatusRequestable  StatusType = "requestable"
)

type CheckoutType string

const (
	CheckoutUser     CheckoutType = "user"
	CheckoutLocation CheckoutType = "location"
	CheckoutAsset    CheckoutType = "asset"
)

type RequestOption func(*RequestOptions)

type RequestOptions struct {
	*params.BaseResolver

	// Query params
	Limit          int32                  `url:"limit,omitempty" json:"-"`
	Offset         int32                  `url:"offset,omitempty" json:"-"`
	Search         string                 `url:"search,omitempty" json:"-"`
	Sort           params.SortType        `url:"sort,omitempty" json:"-"`
	Order          params.OrderType       `url:"order,omitempty" json:"-"`
	CategoryID     snipeit.CategoryID     `url:"category_id,omitempty" json:"-"`
	ManufacturerID snipeit.ManufacturerID `url:"manufacturer_id,omitempty" json:"-"`
	CompanyID      snipeit.CompanyID      `url:"company_id,omitempty" json:"-"`
	Status         StatusType             `url:"status,omitempty" json:"-"`
	Deleted        bool                   `url:"deleted,omitempty" json:"-"`

	// Compound params
	StatusID    snipeit.StatusLabelID `url:"status_id,omitempty" json:"status_id,omitempty"`
	ModelID     snipeit.ModelID       `url:"model_id,omitempty" json:"model_id,omitempty"`
	OrderNumber string                `url:"order_number,omitempty" json:"order_number,omitempty"`
	LocationID  snipeit.LocationID    `url:"location_id,omitempty" json:"location_id,omitempty"`

	// Body params
	AssetTag         string             `json:"asset_tag,omitempty" url:"-"`
	Name             string             `json:"name,omitempty" url:"-"`
	Image            string             `json:"image,omitempty" url:"-"`
	Serial           string             `json:"serial,omitempty" url:"-"`
	PurchaseDate     string             `json:"purchase_date,omitempty" url:"-"`
	PurchaseCost     float64            `json:"purchase_cost,omitempty" url:"-"`
	Notes            string             `json:"notes,omitempty" url:"-"`
	Archived         bool               `json:"archived,omitempty" url:"-"`
	WarrantyMonths   int32              `json:"warranty_months,omitempty" url:"-"`
	Depreciate       bool               `json:"depreciate,omitempty" url:"-"`
	SupplierID       int32              `json:"supplier_id,omitempty" url:"-"`
	Requestable      bool               `json:"requestable,omitempty" url:"-"`
	RTDLocationID    snipeit.LocationID `json:"rtd_location_id,omitempty" url:"-"`
	LastAuditDate    string             `json:"last_audit_date,omitempty" url:"-"`
	BYOD             bool               `json:"byod,omitempty" url:"-"`
	CheckoutToType   CheckoutType       `json:"checkout_to_type,omitempty" url:"-"`
	AssignedUser     snipeit.UserID     `json:"assigned_user,omitempty" url:"-"`
	AssignedAsset    snipeit.AssetID    `json:"assigned_asset,omitempty" url:"-"`
	AssignedLocation snipeit.LocationID `json:"assigned_location,omitempty" url:"-"`
	ExpectedCheckin  string             `json:"expected_checkin,omitempty" url:"-"`
	CheckoutAt       string             `json:"checkout_at,omitempty" url:"-"`
	Note             string             `json:"note,omitempty" url:"-"`
	NextAuditDate    string             `json:"next_audit_date,omitempty" url:"-"`
}

func (ro *RequestOptions) Values() (url.Values, error) {
	return ro.BaseResolver.Values()
}

func (ro *RequestOptions) Marshal() (io.Reader, error) {
	return ro.BaseResolver.Marshal()
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

func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

func ManufacturerID(id snipeit.ManufacturerID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManufacturerID = id
	}
}

func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

func Status(status StatusType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Status = status
	}
}

func Deleted() RequestOption {
	return func(ro *RequestOptions) {
		ro.Deleted = true
	}
}

func StatusID(id snipeit.StatusLabelID) RequestOption {
	return func(ro *RequestOptions) {
		ro.StatusID = id
	}
}

func ModelID(id snipeit.ModelID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ModelID = id
	}
}

func OrderNumber(number string) RequestOption {
	return func(ro *RequestOptions) {
		ro.OrderNumber = number
	}
}

func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

func AssetTag(tag string) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssetTag = tag
	}
}

func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

func Image(mime, data string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Image = fmt.Sprintf("data:@%s;base64,%s", mime, data)
	}
}

func Serial(serial string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Serial = serial
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

func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

func Archived() RequestOption {
	return func(ro *RequestOptions) {
		ro.Archived = true
	}
}

func WarrantyMonths(months int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.WarrantyMonths = months
	}
}

func Depreciate() RequestOption {
	return func(ro *RequestOptions) {
		ro.Depreciate = true
	}
}

func SupplierID(id int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupplierID = id
	}
}

func Requestable() RequestOption {
	return func(ro *RequestOptions) {
		ro.Requestable = true
	}
}

func RTDLocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.RTDLocationID = id
	}
}

func LastAuditDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.LastAuditDate = date
	}
}

func BYOD() RequestOption {
	return func(ro *RequestOptions) {
		ro.BYOD = true
	}
}

func CheckoutToType(checkout CheckoutType) RequestOption {
	return func(ro *RequestOptions) {
		ro.CheckoutToType = checkout
	}
}

func AssignedUser(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedUser = id
	}
}

func AssignedAsset(id snipeit.AssetID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedAsset = id
	}
}

func AssignedLocation(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedLocation = id
	}
}

func ExpectedCheckin(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ExpectedCheckin = date
	}
}

func CheckoutAt(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.CheckoutAt = date
	}
}

func Note(note string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Note = note
	}
}

func NextAuditDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.NextAuditDate = date
	}
}
