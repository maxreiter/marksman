package category

import (
	"io"
	"net/url"

	params "github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type CategoryType string

const (
	CategoryAsset      CategoryType = "asset"
	CategoryAccessory  CategoryType = "accessory"
	CategoryConsumable CategoryType = "consumable"
	CategoryComponent  CategoryType = "component"
	CategoryLicense    CategoryType = "license"
)

type RequestOptions struct {
	*params.BaseResolver

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

func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

func Type(categoryType CategoryType) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryType = categoryType
	}
}

func UseDefaultEULA() RequestOption {
	return func(ro *RequestOptions) {
		ro.UseDefaultEULA = true
	}
}

func RequireAcceptance() RequestOption {
	return func(ro *RequestOptions) {
		ro.RequireAcceptance = true
	}
}

func CheckinEmail() RequestOption {
	return func(ro *RequestOptions) {
		ro.CheckinEmail = true
	}
}
