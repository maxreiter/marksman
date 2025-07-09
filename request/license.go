package request

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type FetchLicenses struct {
	Name           string        `url:"name,omitempty"`
	ProductKey     string        `url:"product_key,omitempty"`
	Limit          int32         `url:"limit,omitempty"`
	Offset         int32         `url:"offset,omitempty"`
	Search         string        `url:"search,omitempty"`
	OrderNumber    string        `url:"order_number,omitempty"`
	Sort           string        `url:"sort,omitempty"`
	Order          SortOrderType `url:"order,omitempty"`
	Expand         bool          `url:"expand,omitempty"`
	PurchaseOrder  string        `url:"purchase_order,omitempty"`
	LicenseName    string        `url:"license_name,omitempty"`
	LicenseEmail   string        `url:"license_email,omitempty"`
	ManufacturerID int32         `url:"manufacturer_id,omitempty"`
	SupplierID     int32         `url:"supplier_id,omitempty"`
	CategoryID     int32         `url:"category_id,omitempty"`
	DepreciationID int32         `url:"deprepciation_id,omitempty"`
	Maintained     bool          `url:"maintained,omitempty"`
	Deleted        bool          `url:"deleted,omitempty"`
}

type CreateLicense struct {
	Name       string `json:"name"`
	Seats      int32  `json:"seats"`
	CategoryID int32  `json:"category_id"`
	CompanyID  int32  `json:"company_id,omitempty"`
	// ExpirationDate
	LicenseEmail string  `json:"license_email,omitempty"`
	LicenseName  string  `json:"license_name,omitempty"`
	Serial       string  `json:"serial"`
	Maintained   bool    `json:"maintained,omitempty"`
	Notes        string  `json:"notes,omitempty"`
	OrderNumber  string  `json:"order_number,omitempty"`
	PurchaseCost float64 `json:"purchase_cost,omitempty"`
	// PurchaseDate
	PurchaseOrder string `json:"purchase_order,omitempty"`
	Reassignable  bool   `json:"reassignable,omitempty"`
	SupplierID    int32  `json:"supplier_id,omitempty"`
	// TerminationDate
}

type UpdateLicense struct {
	Name       string `json:"name"`
	Seats      int32  `json:"seats"`
	CategoryID int32  `json:"category_id"`
	CompanyID  int32  `json:"company_id,omitempty"`
	// ExpirationDate
	LicenseEmail string  `json:"license_email,omitempty"`
	LicenseName  string  `json:"license_name,omitempty"`
	Serial       string  `json:"serial"`
	Maintained   bool    `json:"maintained,omitempty"`
	Notes        string  `json:"notes,omitempty"`
	OrderNumber  string  `json:"order_number,omitempty"`
	PurchaseCost float64 `json:"purchase_cost,omitempty"`
	// PurchaseDate
	PurchaseOrder string `json:"purchase_order,omitempty"`
	Reassignable  bool   `json:"reassignable,omitempty"`
	SupplierID    int32  `json:"supplier_id,omitempty"`
	// TerminationDate
}

type PartiallyUpdateLicense struct {
	Name       string `json:"name,omitempty"`
	Seats      int32  `json:"seats,omitempty"`
	CategoryID int32  `json:"category_id,omitempty"`
	CompanyID  int32  `json:"company_id,omitempty"`
	// ExpirationDate
	LicenseEmail string  `json:"license_email,omitempty"`
	LicenseName  string  `json:"license_name,omitempty"`
	Serial       string  `json:"serial"`
	Maintained   bool    `json:"maintained,omitempty"`
	Notes        string  `json:"notes,omitempty"`
	OrderNumber  string  `json:"order_number,omitempty"`
	PurchaseCost float64 `json:"purchase_cost,omitempty"`
	// PurchaseDate
	PurchaseOrder string `json:"purchase_order,omitempty"`
	Reassignable  bool   `json:"reassignable,omitempty"`
	SupplierID    int32  `json:"supplier_id,omitempty"`
	// TerminationDate
}

type UpdateLicenseSeat struct {
	AssignedTo nullable.Int32 `json:"assigned_to"`
	AssetID    nullable.Int32 `json:"asset_id"`
	Notes      string         `json:"notes,omitempty"`
}
