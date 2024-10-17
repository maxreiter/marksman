package statuslabel

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
)

type StatusType string

const (
	StatusDeployable   StatusType = "deployable"
	StatusUndeployable StatusType = "undeployable"
	StatusPending      StatusType = "pending"
	StatusArchived     StatusType = "archived"
)

type RequestOptions struct {
	*params.BaseResolver

	// Query params
	Limit  int32            `url:"limit,omitempty" json:"-"`
	Offset int32            `url:"offset,omitempty" json:"-"`
	Search string           `url:"search,omitempty" json:"-"`
	Sort   params.SortType  `url:"sort,omitempty" json:"-"`
	Order  params.OrderType `url:"order,omitempty" json:"-"`

	// Compound params
	Name       string     `url:"name,omitempty" json:"name,omitempty"`
	StatusType StatusType `url:"status_type,omitempty" json:"type,omitempty"`

	// Body params
	Notes        string `json:"notes,omitempty" url:"-"`
	Color        string `json:"color,omitempty" url:"-"`
	ShowInNav    bool   `json:"show_in_nav,omitempty" url:"-"`
	DefaultLabel bool   `json:"default_label,omitempty" url:"-"`
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

func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

func Type(status StatusType) RequestOption {
	return func(ro *RequestOptions) {
		ro.StatusType = status
	}
}

func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

func Color(color string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Color = color
	}
}

func ShowInNav() RequestOption {
	return func(ro *RequestOptions) {
		ro.ShowInNav = true
	}
}

func DefaultLabel() RequestOption {
	return func(ro *RequestOptions) {
		ro.DefaultLabel = true
	}
}
