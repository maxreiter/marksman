// Package params provides universal parameters for use in subpackages.
package params

// SortType represents the different columns one may sort by in a request.
type SortType string

// The various columns a request response may be sorted by.
const (
	SortID              SortType = "id"
	SortName            SortType = "name"
	SortAssetTag        SortType = "asset_tag"
	SortSerial          SortType = "serial"
	SortModel           SortType = "model"
	SortModelNumber     SortType = "model_number"
	SortLastCheckout    SortType = "last_checkout"
	SortCategory        SortType = "category"
	SortManufacturer    SortType = "manufacturer"
	SortNotes           SortType = "notes"
	SortExpectedCheckin SortType = "expected_checkin"
	SortOrderNumber     SortType = "order_number"
	SortCompanyName     SortType = "companyName"
	SortLocation        SortType = "location"
	SortImage           SortType = "image"
	SortStatusLabel     SortType = "status_label"
	SortAssignedTo      SortType = "assigned_to"
	SortCreatedAt       SortType = "created_at"
	SortPurchaseDate    SortType = "purchase_date"
	SortPurchaseCost    SortType = "purchase_cost"
)

// OrderType represents the way one may order results from a request.
type OrderType string

// Different directions a request response may be ordered.
const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)
