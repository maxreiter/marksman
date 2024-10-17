package model

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type RequestOptions struct {
	*params.BaseResolver

	Limit  int32            `url:"limit,omitempty" json:"-"`
	Offset int32            `url:"offset,omitempty" json:"-"`
	Search string           `url:"search,omitempty" json:"-"`
	Sort   params.SortType  `url:"sort,omitempty" json:"-"`
	Order  params.OrderType `url:"order,omitempty" json:"-"`

	Name           string                 `json:"name,omitempty" url:"-"`
	ModelNumber    string                 `json:"model_number,omitempty" url:"-"`
	CategoryID     snipeit.CategoryID     `json:"category_id,omitempty" url:"-"`
	ManufacturerID snipeit.ManufacturerID `json:"manufacturer_id,omitempty" url:"-"`
	EOL            int32                  `json:"eol,omitempty" url:"-"`
	FieldsetID     snipeit.FieldsetID     `json:"fieldset_id,omitempty" url:"-"`
	DepreciationID snipeit.DepreciationID `json:"depreciation_id,omitempty" url:"-"`
	Notes          string                 `json:"notes,omitempty" url:"-"`
	Requestable    bool                   `json:"requestable,omitempty" url:"-"`
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

func ModelNumber(num string) RequestOption {
	return func(ro *RequestOptions) {
		ro.ModelNumber = num
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

func EOL(eol int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.EOL = eol
	}
}

func FieldsetID(id snipeit.FieldsetID) RequestOption {
	return func(ro *RequestOptions) {
		ro.FieldsetID = id
	}
}

func DepreciationID(id snipeit.DepreciationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.DepreciationID = id
	}
}

func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

func Requestable() RequestOption {
	return func(ro *RequestOptions) {
		ro.Requestable = true
	}
}
