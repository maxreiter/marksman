package request

import "github.com/maxreiter/marksman/snipeit"

type FetchAccessories struct {
	Limit       int32         `url:"limit,omitempty"`
	Offset      int32         `url:"offset,omitempty"`
	Search      string        `url:"search,omitempty"`
	OrderNumber string        `url:"order_number,omitempty"`
	Sort        string        `url:"sort,omitempty"`
	Order       SortOrderType `url:"order,omitempty"`
	Expand      bool          `url:"expand,omitempty"`
}

type CreateAccessory struct {
	Name           string       `json:"name"`
	Quantity       int32        `json:"qty"`
	OrderNumber    string       `json:"order_number,omitempty"`
	PurchaseCost   float64      `json:"purchase_cost,omitempty"`
	PurchaseDate   snipeit.Date `json:"purchase_date"`
	ModelNumber    string       `json:"model_number,omitempty"`
	CategoryID     int32        `json:"category_id"`
	CompanyID      int32        `json:"company_id,omitempty"`
	LocationID     int32        `json:"location_id,omitempty"`
	ManufacturerID int32        `json:"manufacturer_id,omitempty"`
	SupplierID     int32        `json:"supplier_id,omitempty"`
}

type UpdateAccessory struct {
	Name           string       `json:"name"`
	Quantity       int32        `json:"qty"`
	CategoryID     int32        `json:"category_id"`
	OrderNumber    string       `json:"order_number,omitempty"`
	PurchaseCost   float64      `json:"purchase_cost,omitempty"`
	PurchaseDate   snipeit.Date `json:"purchase_date,omitzero"`
	ModelNumber    string       `json:"model_number,omitempty"`
	CompanyID      int32        `json:"company_id,omitempty"`
	LocationID     int32        `json:"location_id,omitempty"`
	ManufacturerID int32        `json:"manufacturer_id,omitempty"`
	SupplierID     int32        `json:"supplier_id,omitempty"`
}

type PartiallyUpdateAccessory struct {
	Name           string       `json:"name,omitempty"`
	Quantity       int32        `json:"qty,omitempty"`
	CategoryID     int32        `json:"category_id,omitempty"`
	OrderNumber    string       `json:"order_number,omitempty"`
	PurchaseCost   float64      `json:"purchase_cost,omitempty"`
	PurchaseDate   snipeit.Date `json:"purchase_date,omitzero"`
	ModelNumber    string       `json:"model_number,omitempty"`
	CompanyID      int32        `json:"company_id,omitempty"`
	LocationID     int32        `json:"location_id,omitempty"`
	ManufacturerID int32        `json:"manufacturer_id,omitempty"`
	SupplierID     int32        `json:"supplier_id,omitempty"`
}

type FetchCheckedOutAccessories struct {
	Limit  int32 `url:"limit,omitempty"`
	Offset int32 `url:"offset,omitempty"`
}

type CheckOutAccessory struct {
	AssignedUser     int32  `json:"assigned_user"`
	Notes            string `json:"note,omitempty"`
	CheckoutQuantity int32  `json:"checkout_qty,omitempty"`
}
