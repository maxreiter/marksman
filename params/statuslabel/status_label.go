// Package statuslabel provides request configuration for methods of the marksman client.
package statuslabel

import (
	"github.com/maxreiter/marksman/params"
)

// StatusType represents the status of an [snipeit.StatusLabel].
type StatusType string

// Possible statuses that a [snipeit.StatusLabel] may be.
const (
	StatusDeployable   StatusType = "deployable"
	StatusUndeployable StatusType = "undeployable"
	StatusPending      StatusType = "pending"
	StatusArchived     StatusType = "archived"
)

// RequestOptions contains possible options for requests made to the /statuslabels endpoints.
type RequestOptions struct {

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

// Name sets the name of a [snipeit.StatusLabel].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// Type sets the status type of a [snipeit.StatusLabel].
func Type(status StatusType) RequestOption {
	return func(ro *RequestOptions) {
		ro.StatusType = status
	}
}

// Notes set the notes of a [snipeit.StatusLabel].
func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

// Color sets the color of a [snipeit.StatusLabel].
func Color(color string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Color = color
	}
}

// ShowInNav marks a [snipeit.StatusLabel] to be shown in navigation.
func ShowInNav() RequestOption {
	return func(ro *RequestOptions) {
		ro.ShowInNav = true
	}
}

// DefaultLabel marks a [snipeit.StatusLabel] to use the default label.
func DefaultLabel() RequestOption {
	return func(ro *RequestOptions) {
		ro.DefaultLabel = true
	}
}
