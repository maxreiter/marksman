package request

import "github.com/maxreiter/marksman/snipeit"

type FetchConsumables struct {
	Name           string        `url:"name,omitempty"`
	Limit          int32         `url:"limit,omitempty"`
	Offset         int32         `url:"offset,omitempty"`
	Search         string        `url:"search,omitempty"`
	OrderNumber    string        `url:"order_number,omitempty"`
	Sort           string        `url:"sort,omitempty"`
	Order          SortOrderType `url:"order,omitempty"`
	Expand         bool          `url:"expand,omitempty"`
	CategoryID     int32         `url:"category_id,omitempty"`
	CompanyID      int32         `url:"company_id,omitempty"`
	ManufacturerID int32         `url:"manufacturer_id,omitempty"`
}

type CreateConsumable struct {
	Name           string       `json:"name"`
	Quantity       int32        `json:"qty"`
	CategoryID     int32        `json:"category_id"`
	CompanyID      int32        `json:"company_id,omitempty"`
	OrderNumber    int32        `json:"order_number,omitempty"`
	ManufacturerID int32        `json:"manufacturer_id,omitempty"`
	LocationID     int32        `json:"location_id,omitempty"`
	Requestable    bool         `json:"requestable,omitempty"`
	PurchaseDate   snipeit.Date `json:"purchase_date,omitzero"`
	MinimumAmount  int32        `json:"min_amt,omitempty"`
	ModelNumber    string       `json:"model_number,omitempty"`
	ItemNumber     string       `json:"item_no,omitempty"`
}

type UpdateConsumable struct {
	Name           string       `json:"name"`
	Quantity       int32        `json:"qty"`
	CategoryID     int32        `json:"category_id"`
	CompanyID      int32        `json:"company_id,omitempty"`
	OrderNumber    int32        `json:"order_number,omitempty"`
	ManufacturerID int32        `json:"manufacturer_id,omitempty"`
	LocationID     int32        `json:"location_id,omitempty"`
	Requestable    bool         `json:"requestable,omitempty"`
	PurchaseDate   snipeit.Date `json:"purchase_date,omitzero"`
	MinimumAmount  int32        `json:"min_amt,omitempty"`
	ModelNumber    string       `json:"model_number,omitempty"`
	ItemNumber     string       `json:"item_no,omitempty"`
}

type PartiallyUpdateConsumable struct {
	Name           string       `json:"name,omitempty"`
	Quantity       int32        `json:"qty,omitempty"`
	CategoryID     int32        `json:"category_id,omitempty"`
	CompanyID      int32        `json:"company_id,omitempty"`
	OrderNumber    int32        `json:"order_number,omitempty"`
	ManufacturerID int32        `json:"manufacturer_id,omitempty"`
	LocationID     int32        `json:"location_id,omitempty"`
	Requestable    bool         `json:"requestable,omitempty"`
	PurchaseDate   snipeit.Date `json:"purchase_date,omitzero"`
	MinimumAmount  int32        `json:"min_amt,omitempty"`
	ModelNumber    string       `json:"model_number,omitempty"`
	ItemNumber     string       `json:"item_no,omitempty"`
}

type CheckoutConsumable struct {
	AssignedTo       int32 `json:"assigned_to"`
	CheckoutQuantity int32 `json:"checkout_qty"`
}
