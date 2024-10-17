// Package report provides request configuration for methods of the marksman client.
package report

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// RequestOptions contains possible options for requests made to the /reports endpoints.
type RequestOptions struct {
	*params.Resolver

	Limit      int32              `url:"limit,omitempty" json:"-"`
	Offset     int32              `url:"offset,omitempty" json:"-"`
	Search     string             `url:"search,omitempty" json:"-"`
	TargetType string             `url:"target_type,omitempty" json:"-"`
	TargetID   snipeit.ID         `url:"target_id,omitempty" json:"-"`
	ItemType   string             `url:"item_type,omitempty" json:"-"`
	ItemID     snipeit.ID         `url:"item_id,omitempty" json:"-"`
	ActionType snipeit.ActionType `url:"action_type,omitempty" json:"-"`
	Order      params.OrderType   `url:"order,omitempty" json:"-"`
	Sort       params.SortType    `url:"sort,omitempty" json:"-"`
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

// TargetType sets the target type of a [snipeit.ActionLog].
func TargetType(targetType string) RequestOption {
	return func(ro *RequestOptions) {
		ro.TargetType = targetType
	}
}

// TargetID sets the target ID of a [snipeit.ActionLog].
func TargetID(id snipeit.ID) RequestOption {
	return func(ro *RequestOptions) {
		ro.TargetID = id
	}
}

// ItemType sets the item type of a [snipeit.ActionLog].
func ItemType(itemType string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ItemType = itemType
	}
}

// ItemID sets the item ID of a [snipeit.ActionLog].
func ItemID(id snipeit.ID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ItemID = id
	}
}

// ActionType sets the action type of a [snipeit.ActionLog].
func ActionType(actionType snipeit.ActionType) RequestOption {
	return func(ro *RequestOptions) {
		ro.ActionType = actionType
	}
}

// Order sets the return order of a request.
func Order(order params.OrderType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Order = order
	}
}

// Sort sets the return sort of a request.
func Sort(sort params.SortType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Sort = sort
	}
}
