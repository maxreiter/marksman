package maintenance

import (
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type AssetMaintenanceType string

const (
	MaintenanceGeneral         AssetMaintenanceType = "Maintenance"
	MaintenanceRepair          AssetMaintenanceType = "Repair"
	MaintenancePATTest         AssetMaintenanceType = "PAT test"
	MaintenanceUpgrade         AssetMaintenanceType = "Upgrade"
	MaintenanceHardwareSupport AssetMaintenanceType = "Hardware Support"
	MaintenanceSoftwareSupport AssetMaintenanceType = "Software Support"
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
	AssetID snipeit.AssetID `url:"asset_id,omitempty" json:"asset_id,omitempty"`

	// Body params
	Title                string               `json:"title,omitempty" url:"-"`
	SupplierID           int32                `json:"supplier_id,omitempty" url:"-"`
	IsWarranty           bool                 `json:"is_warranty,omitempty" url:"-"`
	Cost                 float64              `json:"cost,omitempty" url:"-"`
	Notes                string               `json:"notes,omitempty" url:"-"`
	AssetMaintenanceType AssetMaintenanceType `json:"asset_maintenance_type,omitempty" url:"-"`
	StartDate            string               `json:"start_date,omitempty" url:"-"`
	CompletionDate       string               `json:"completion_date,omitempty" url:"-"`
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

func AssetID(id snipeit.AssetID) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssetID = id
	}
}

func Title(title string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Title = title
	}
}

func SupplierID(id int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.SupplierID = id
	}
}

func IsWarranty() RequestOption {
	return func(ro *RequestOptions) {
		ro.IsWarranty = true
	}
}

func Cost(cost float64) RequestOption {
	return func(ro *RequestOptions) {
		ro.Cost = cost
	}
}

func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

func MaintenanceType(maintenanceType AssetMaintenanceType) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssetMaintenanceType = maintenanceType
	}
}

func StartDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.StartDate = date
	}
}

func CompletionDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompletionDate = date
	}
}
