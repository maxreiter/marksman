package request

import "github.com/maxreiter/marksman/snipeit"

type FetchMaintenances struct {
	Limit   int32         `url:"limit,omitempty"`
	Offset  int32         `url:"offset,omitempty"`
	Search  string        `url:"search,omitempty"`
	Sort    string        `url:"sort,omitempty"`
	Order   SortOrderType `url:"order,omitempty"`
	AssetID int32         `url:"asset_id,omitempty"`
}

type CreateMaintenance struct {
	Title                string                       `json:"title"`
	AssetID              int32                        `json:"asset_id"`
	SupplierID           int32                        `json:"supplier_id"`
	IsWarranty           bool                         `json:"is_warranty,omitempty"`
	Cost                 float64                      `json:"cost,omitempty"`
	Notes                string                       `json:"notes,omitempty"`
	AssetMaintenanceType snipeit.AssetMaintenanceType `json:"asset_maintenance_type"`
	// StartDate
	// CompletionDate
}

type UpdateMaintenance struct {
	Title                string                       `json:"title"`
	AssetID              int32                        `json:"asset_id"`
	SupplierID           int32                        `json:"supplier_id"`
	IsWarranty           bool                         `json:"is_warranty,omitempty"`
	Cost                 float64                      `json:"cost,omitempty"`
	Notes                string                       `json:"notes,omitempty"`
	AssetMaintenanceType snipeit.AssetMaintenanceType `json:"asset_maintenance_type"`
	// StartDate
	// CompletionDate
}

type PartiallyUpdateMaintenance struct {
	Title                string                       `json:"title,omitempty"`
	AssetID              int32                        `json:"asset_id,omitempty"`
	SupplierID           int32                        `json:"supplier_id,omitempty"`
	IsWarranty           bool                         `json:"is_warranty,omitempty"`
	Cost                 float64                      `json:"cost,omitempty"`
	Notes                string                       `json:"notes,omitempty"`
	AssetMaintenanceType snipeit.AssetMaintenanceType `json:"asset_maintenance_type"`
	// StartDate
	// CompletionDate
}
