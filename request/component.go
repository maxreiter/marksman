package request

import "github.com/maxreiter/marksman/snipeit"

type FetchComponents struct {
	Name        string        `url:"name,omitempty"`
	Limit       int32         `url:"limit,omitempty"`
	Offset      int32         `url:"offset,omitempty"`
	Search      string        `url:"search,omitempty"`
	OrderNumber string        `url:"order_number,omitempty"`
	Sort        string        `url:"sort,omitempty"`
	Order       SortOrderType `url:"order,omitempty"`
	Expand      bool          `url:"expand,omitempty"`
}

type CreateComponent struct {
	Name        string `json:"name"`
	Quantity    int32  `json:"qty"`
	CategoryID  int32  `json:"category_id"`
	LocationID  int32  `json:"location_id,omitempty"`
	CompanyID   int32  `json:"company_id,omitempty"`
	OrderNumber string `json:"order_number,omitempty"`
	// PurchaseDate
	PurchaseCost  float64 `json:"purchase_cost,omitempty"`
	MinimumAmount int32   `json:"min_amt,omitempty"`
	Serial        string  `json:"serial,omitempty"`
	ModelNumber   string  `json:"model_number,omitempty"`
}

type UpdateComponent struct {
	Name          string       `json:"name"`
	Quantity      int32        `json:"qty"`
	CategoryID    int32        `json:"category_id"`
	LocationID    int32        `json:"location_id,omitempty"`
	CompanyID     int32        `json:"company_id,omitempty"`
	OrderNumber   string       `json:"order_number,omitempty"`
	PurchaseDate  snipeit.Date `json:"purchase_date,omitzero"`
	PurchaseCost  float64      `json:"purchase_cost,omitempty"`
	MinimumAmount int32        `json:"min_amt,omitempty"`
	Serial        string       `json:"serial,omitempty"`
	ModelNumber   string       `json:"model_number,omitempty"`
}

type PartiallyUpdateComponent struct {
	Name          string       `json:"name,omitempty"`
	Quantity      int32        `json:"qty,omitempty"`
	CategoryID    int32        `json:"category_id,omitempty"`
	LocationID    int32        `json:"location_id,omitempty"`
	CompanyID     int32        `json:"company_id,omitempty"`
	OrderNumber   string       `json:"order_number,omitempty"`
	PurchaseDate  snipeit.Date `json:"purchase_date,omitzero"`
	PurchaseCost  float64      `json:"purchase_cost,omitempty"`
	MinimumAmount int32        `json:"min_amt,omitempty"`
	Serial        string       `json:"serial,omitempty"`
	ModelNumber   string       `json:"model_number,omitempty"`
}

type CheckoutComponent struct {
	AssignedTo       int32 `json:"assigned_to"`
	AssignedQuantity int32 `json:"assigned_qty"`
}

type CheckinComponent struct {
	CheckinQuantity int32 `json:"checkin_qty"`
}
