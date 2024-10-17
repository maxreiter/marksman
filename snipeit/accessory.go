package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

// Accessory represents an accessory in the SnipeIT interface.
type Accessory struct {
	ID               AccessoryID         `json:"id"`
	Name             string              `json:"name"`
	Image            null.NullableString `json:"image,omitempty"`
	Company          *Company            `json:"company,omitempty"`
	Manufacturer     *Manufacturer       `json:"manufacturer,omitempty"`
	Supplier         *Supplier           `json:"supplier,omitempty"`
	ModelNumber      null.NullableString `json:"model_number,omitempty"`
	Category         *Category           `json:"category,omitempty"`
	Location         *Location           `json:"location,omitempty"`
	Notes            null.NullableString `json:"notes,omitempty"`
	Qty              null.NullableInt    `json:"qty,omitempty"`
	PurchaseDate     Datetime            `json:"purchase_date,omitempty"`
	PurchaseCost     string              `json:"purchase_cost,omitempty"`
	OrderNumber      null.NullableString `json:"order_number,omitempty"`
	MinQty           null.NullableInt    `json:"min_qty,omitempty"`
	RemainingQty     null.NullableInt    `json:"remaining_qty,omitempty"`
	CheckoutsCount   int                 `json:"checkouts_count,omitempty"`
	CreatedBy        *User               `json:"created_by,omitempty"`
	CreatedAt        Datetime            `json:"created_at,omitempty"`
	UpdatedAt        Datetime            `json:"updated_at,omitempty"`
	AvailableActions *AvailableActions   `json:"available_actions,omitempty"`
}
