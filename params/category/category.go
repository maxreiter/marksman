// Package category provides request configuration for methods of the marksman client.
package category

import (
	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// CategoryType defines the types of models a [snipeit.Category] may represent.
type CategoryType string

// The types of models a [snipeit.Category] may represent.
const (
	CategoryAsset      CategoryType = "asset"
	CategoryAccessory  CategoryType = "accessory"
	CategoryConsumable CategoryType = "consumable"
	CategoryComponent  CategoryType = "component"
	CategoryLicense    CategoryType = "license"
)

// RequestOptions contains possible options for requests made to the /categories endpoints.
type RequestOptions struct {

	// Query params
	Limit      int32              `url:"limit,omitempty" json:"-"`
	Offset     int32              `url:"offset,omitempty" json:"-"`
	Search     string             `url:"search,omitempty" json:"-"`
	Sort       params.SortType    `url:"sort,omitempty" json:"-"`
	Order      params.OrderType   `url:"order,omitempty" json:"-"`
	CategoryID snipeit.CategoryID `url:"category_id,omitempty" json:"-"`

	// Compound params
	Name              string       `url:"name,omitempty" json:"name,omitempty"`
	CategoryType      CategoryType `url:"category_type,omitempty" json:"category_type,omitempty"`
	UseDefaultEULA    bool         `url:"use_default_eula,omitempty" json:"use_default_eula,omitempty"`
	RequireAcceptance bool         `url:"require_acceptance,omitempty" json:"require_acceptance,omitempty"`
	CheckinEmail      bool         `url:"checkin_email,omitempty" json:"checkin_email,omitempty"`
}

// RequestOption is used to configure a [RequestOptions]
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

// CategoryID sets the [snipeit.CategoryID] of a [snipeit.Category].
func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

// Name sets the name of a [snipeit.Category].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// Type sets the type of [snipeit.Category].
func Type(categoryType CategoryType) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryType = categoryType
	}
}

// UseDefaultEULA forces a [snipeit.Category] to use the default EULA.
func UseDefaultEULA() RequestOption {
	return func(ro *RequestOptions) {
		ro.UseDefaultEULA = true
	}
}

// RequireAcceptance marks acceptance as required.
func RequireAcceptance() RequestOption {
	return func(ro *RequestOptions) {
		ro.RequireAcceptance = true
	}
}

// CheckinEmail toggles email checkin on a [snipeit.Category].
func CheckinEmail() RequestOption {
	return func(ro *RequestOptions) {
		ro.CheckinEmail = true
	}
}
