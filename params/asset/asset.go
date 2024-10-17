// Package asset provides request configuration for methods of the [marksman.Client].
package asset

import (
	"fmt"
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// StatusType represents the status of an [snipeit.Asset].
type StatusType string

// The different types of statuses an [snipeit.Asset] may be marked as.
const (
	StatusRTD          StatusType = "rtd"
	StatusDeployed     StatusType = "deployed"
	StatusUndeployable StatusType = "undeployable"
	StatusDeleted      StatusType = "deleted"
	StatusArchived     StatusType = "archived"
	StatusRequestable  StatusType = "requestable"
)

// CheckoutType represents the different models an [snipeit.Asset] may be checked out to.
type CheckoutType string

// The different kinds of models an [snipeit.Asset] may be checked out to.
const (
	CheckoutUser     CheckoutType = "user"
	CheckoutLocation CheckoutType = "location"
	CheckoutAsset    CheckoutType = "asset"
)

// RequestOption is used to configure a [RequestOptions].
type RequestOption func(*RequestOptions)

// RequestOptions contains possible options for requests made to the /assets endpoints.
type RequestOptions struct {
	*params.Resolver

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

// Query marshals the [RequestOptions] into a [url.Values].
func (ro *RequestOptions) Query() (url.Values, error) {
	return ro.Resolver.Query()
}

// JSON encodes the [RequestOptions] as JSON.
func (ro *RequestOptions) JSON() (io.Reader, error) {
	return ro.Resolver.JSON()
}

// Limit sets the return limit on a request.
func Limit(limit int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Limit = limit
	}
}

// Offset sets the pagination offset on a request.
func Offset(offset int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Offset = offset
	}
}

// Search sets the search string on a request.
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

// CategoryID sets the [snipeit.CategoryID] of an [snipeit.Asset].
func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

// ManufacturerID sets the [snipeit.ManufacturerID] of an [snipeit.Asset].
func ManufacturerID(id snipeit.ManufacturerID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManufacturerID = id
	}
}

// CompanyID sets the [snipeit.CompanyID] of an [snipeit.Asset].
func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

// Status sets the status type of an [snipeit.Asset].
func Status(status StatusType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Status = status
	}
}

// Deleted toggles the viewing of deleted [snipeit.Asset].
func Deleted() RequestOption {
	return func(ro *RequestOptions) {
		ro.Deleted = true
	}
}

// StatusID sets the [snipeit.StatusLabelID] of an [snipeit.Asset].
func StatusID(id snipeit.StatusLabelID) RequestOption {
	return func(ro *RequestOptions) {
		ro.StatusID = id
	}
}

// ModelID sets the [snipeit.ModelID] of an [snipeit.Asset].
func ModelID(id snipeit.ModelID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ModelID = id
	}
}

// OrderNumber sets the order number of an [snipeit.Asset].
func OrderNumber(number string) RequestOption {
	return func(ro *RequestOptions) {
		ro.OrderNumber = number
	}
}

// LocationID sets the [snipeit.LocationID] of an [snipeit.Asset].
func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

// AssetTag sets the asset tag of an [snipeit.Asset].
func AssetTag(tag string) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssetTag = tag
	}
}

// Name sets the name of an [snipeit.Asset].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// Image sets the image data of an [snipeit.Asset].
func Image(mime, data string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Image = fmt.Sprintf("data:@%s;base64,%s", mime, data)
	}
}

// Serial sets the serial of an [snipeit.Asset].
func Serial(serial string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Serial = serial
	}
}

// PurchaseDate sets the date an [snipeit.Asset] was purchased.
func PurchaseDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseDate = date
	}
}

// PurchaseCost sets the cost of an [snipeit.Asset].
func PurchaseCost(cost float64) RequestOption {
	return func(ro *RequestOptions) {
		ro.PurchaseCost = cost
	}
}

// Notes sets the notes of an [snipeit.Asset].
func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

// Archived marks an [snipeit.Asset] as archived.
func Archived() RequestOption {
	return func(ro *RequestOptions) {
		ro.Archived = true
	}
}

// WarrantyMonths sets the number of months an [snipeit.Asset] has left on its warranty.
func WarrantyMonths(months int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.WarrantyMonths = months
	}
}

// Depreciate marks an [snipeit.Asset] as depreciated.
func Depreciate() RequestOption {
	return func(ro *RequestOptions) {
		ro.Depreciate = true
	}
}

// SupplierID sets the [snipeit.SupplierID] of an [snipeit.Asset].
func SupplierID(id int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupplierID = id
	}
}

// Requestable marks an [snipeit.Asset] as requesable.
func Requestable() RequestOption {
	return func(ro *RequestOptions) {
		ro.Requestable = true
	}
}

// RTDLocationID sets the RTD Location ID of an [snipeit.Asset].
func RTDLocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.RTDLocationID = id
	}
}

// LastAuditDate sets the last date an [snipeit.Asset] was audited.
func LastAuditDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.LastAuditDate = date
	}
}

// BYOD marks an [snipeit.Asset] as BYOD.
func BYOD() RequestOption {
	return func(ro *RequestOptions) {
		ro.BYOD = true
	}
}

// CheckoutToType sets the checkout type of an [snipeit.Asset].
func CheckoutToType(checkout CheckoutType) RequestOption {
	return func(ro *RequestOptions) {
		ro.CheckoutToType = checkout
	}
}

// AssignedUser sets the assigned [snipeit.User] of an [snipeit.Asset].
func AssignedUser(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedUser = id
	}
}

// AssignedAsset sets the assigned [snipeit.Asset] of an asset.
func AssignedAsset(id snipeit.AssetID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedAsset = id
	}
}

// AssignedLocation sets the assigned [snipeit.Location] of an [snipeit.Asset].
func AssignedLocation(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssignedLocation = id
	}
}

// ExpectedCheckin sets the expected checkin date of an [snipeit.Asset].
func ExpectedCheckin(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ExpectedCheckin = date
	}
}

// CheckoutAt sets the date an [snipeit.Asset] was checked out.
func CheckoutAt(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.CheckoutAt = date
	}
}

// Note sets the note of an [snipeit.Asset].
func Note(note string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Note = note
	}
}

// NextAuditDate sets the next date an audit should be performed on an [snipeit.Asset].
func NextAuditDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.NextAuditDate = date
	}
}
