package request

import (
	"github.com/maxreiter/marksman/snipeit"
)

type AssetStatusType string

const (
	AssetStatusTypeRTD          AssetStatusType = "rtd"
	AssetStatusTypeDeployed     AssetStatusType = "deployed"
	AssetStatusTypeUndeployable AssetStatusType = "undeployable"
	AssetStatusTypeDeleted      AssetStatusType = "deleted"
	AssetStatusTypeArchived     AssetStatusType = "archived"
	AssetStatuysTypeRequestable AssetStatusType = "requestable"
)

type FetchAssets struct {
	Limit          int32           `url:"limit,omitempty"`
	Offset         int32           `url:"offset,omitempty"`
	Search         string          `url:"search,omitempty"`
	OrderNumber    string          `url:"order_number,omitempty"`
	Sort           string          `url:"sort,omitempty"`
	Order          SortOrderType   `url:"order,omitempty"`
	ModelID        int32           `url:"model_id,omitempty"`
	CategoryID     int32           `url:"category_id,omitempty"`
	ManufacturerID int32           `url:"manufacturer_id,omitempty"`
	CompanyID      int32           `url:"company_id,omitempty"`
	LocationID     int32           `url:"location_id,omitempty"`
	Status         AssetStatusType `url:"status,omitempty"`
	StatusID       int32           `url:"status_id,omitempty"`
}

type CreateAsset struct {
	AssetTag       string       `json:"asset_tag"`
	StatusID       int32        `json:"status_id"`
	ModelID        int32        `json:"model_id"`
	Name           string       `json:"name,omitempty"`
	Image          string       `json:"image,omitempty"`
	Serial         string       `json:"serial,omitempty"`
	PurchaseDate   snipeit.Date `json:"purchase_date,omitzero"`
	PurchaseCost   float64      `json:"purchase_cost,omitempty"`
	OrderNumber    string       `json:"order_number,omitempty"`
	Notes          string       `json:"notes,omitempty"`
	Archived       bool         `json:"archived,omitempty"`
	WarrantyMonths int32        `json:"warranty_months,omitempty"`
	Depreciate     bool         `json:"depreciate,omitempty"`
	SupplierID     int32        `json:"supplier_id,omitempty"`
	Requestable    bool         `json:"requestable,omitempty"`
	RTDLocationID  int32        `json:"rtd_location_id,omitempty"`
	LastAuditDate  snipeit.Date `json:"last_audit_date,omitzero"`
	LocationID     int32        `json:"location_id,omitempty"`
	BYOD           bool         `json:"byod,omitempty"`
}

type FetchDeletedAsset struct {
	Deleted bool `url:"deleted,omitempty"`
}

type UpdateAsset struct {
	AssetTag         string       `json:"asset_tag"`
	StatusID         int32        `json:"status_id"`
	ModelID          int32        `json:"model_id"`
	Notes            string       `json:"notes,omitempty"`
	LastCheckout     snipeit.Date `json:"last_checkout,omitzero"`
	AssignedUser     int32        `json:"assigned_user,omitempty"`
	AssignedLocation int32        `json:"assigned_location,omitempty"`
	AssignedAsset    int32        `json:"assigned_asset,omitempty"`
	CompanyID        int32        `json:"company_id,omitempty"`
	Serial           string       `json:"serial,omitempty"`
	OrderNumber      string       `json:"order_number,omitempty"`
	WarrantyMonths   int32        `json:"warranty_months,omitempty"`
	PurchaseCost     float64      `json:"purchase_cost,omitempty"`
	PurchaseDate     snipeit.Date `json:"purchase_date,omitzero"`
	Requestable      bool         `json:"requestable,omitempty"`
	Archived         bool         `json:"archived,omitempty"`
	RTDLocationID    int32        `json:"rtd_location_id,omitempty"`
	Name             string       `json:"name,omitempty"`
	LocationID       int32        `json:"location_id,omitempty"`
	Image            string       `json:"image,omitempty"`
	BYOD             bool         `json:"byod,omitempty"`
}

type PartiallyUpdateAsset struct {
	AssetTag         string       `json:"asset_tag,omitempty"`
	StatusID         int32        `json:"status_id,omitempty"`
	ModelID          int32        `json:"model_id,omitempty"`
	Notes            string       `json:"notes,omitempty"`
	LastCheckout     snipeit.Date `json:"last_checkout,omitzero"`
	AssignedUser     int32        `json:"assigned_user,omitempty"`
	AssignedLocation int32        `json:"assigned_location,omitempty"`
	AssignedAsset    int32        `json:"assigned_asset,omitempty"`
	CompanyID        int32        `json:"company_id,omitempty"`
	Serial           string       `json:"serial,omitempty"`
	OrderNumber      string       `json:"order_number,omitempty"`
	WarrantyMonths   int32        `json:"warranty_months,omitempty"`
	PurchaseCost     float64      `json:"purchase_cost,omitempty"`
	PurchaseDate     snipeit.Date `json:"purchase_date,omitzero"`
	Requestable      bool         `json:"requestable,omitempty"`
	Archived         bool         `json:"archived,omitempty"`
	RTDLocationID    int32        `json:"rtd_location_id,omitempty"`
	Name             string       `json:"name,omitempty"`
	LocationID       int32        `json:"location_id,omitempty"`
	Image            string       `json:"image,omitempty"`
	BYOD             bool         `json:"byod,omitempty"`
}

type CheckOutToType string

const (
	CheckOutToTypeUser     CheckOutToType = "user"
	CheckOutToTypeAsset    CheckOutToType = "asset"
	CheckOutToTypeLocation CheckOutToType = "location"
)

type CheckOutAsset struct {
	StatusID         int32          `json:"status_id"`
	CheckOutToType   CheckOutToType `json:"checkout_to_type"`
	AssignedUser     int32          `json:"assigned_user,omitempty"`
	AssignedAsset    int32          `json:"assigned_asset,omitempty"`
	AssignedLocation int32          `json:"assigned_location,omitempty"`
	// ExpectedCheckin
	// CheckoutAt
	Name  string `json:"name,omitempty"`
	Notes string `json:"notes,omitempty"`
}

type CheckInAsset struct {
	StatusID   int32  `json:"status_id"`
	Name       string `json:"name,omitempty"`
	Note       string `json:"note,omitempty"`
	LocationID int32  `json:"location_id,omitempty"`
}

type AuditAsset struct {
	AssetTag      string       `json:"asset_tag"`
	LocationID    int32        `json:"location_id,omitempty"`
	NextAuditDate snipeit.Date `json:"next_audit_date,omitzero"`
}

type GenerateAssetLabels struct {
	AssetTags []string `json:"asset_tags"`
}
