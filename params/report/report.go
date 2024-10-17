package report

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type RequestActionType string

const (
	ActionAddSeats    RequestActionType = "add seats"
	ActionCheckinFrom RequestActionType = "checkin from"
	ActionCheckout    RequestActionType = "checkout"
	ActionUpdate      RequestActionType = "update"
)

type RequestOptions struct {
	*params.BaseResolver

	Limit      int32             `url:"limit,omitempty" json:"-"`
	Offset     int32             `url:"offset,omitempty" json:"-"`
	Search     string            `url:"search,omitempty" json:"-"`
	TargetType string            `url:"target_type,omitempty" json:"-"`
	TargetID   snipeit.ID        `url:"target_id,omitempty" json:"-"`
	ItemType   string            `url:"item_type,omitempty" json:"-"`
	ItemID     snipeit.ID        `url:"item_id,omitempty" json:"-"`
	ActionType RequestActionType `url:"action_type,omitempty" json:"-"`
	Order      params.OrderType  `url:"order,omitempty" json:"-"`
	Sort       params.SortType   `url:"sort,omitempty" json:"-"`
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

func TargetType(targetType string) RequestOption {
	return func(ro *RequestOptions) {
		ro.TargetType = targetType
	}
}

func TargetID(id snipeit.ID) RequestOption {
	return func(ro *RequestOptions) {
		ro.TargetID = id
	}
}

func ItemType(itemType string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ItemType = itemType
	}
}

func ItemID(id snipeit.ID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ItemID = id
	}
}

func ActionType(actionType RequestActionType) RequestOption {
	return func(ro *RequestOptions) {
		ro.ActionType = actionType
	}
}

func Order(order params.OrderType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Order = order
	}
}

func Sort(sort params.SortType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Sort = sort
	}
}
