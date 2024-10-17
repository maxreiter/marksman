package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

// Component represents a component in the SnipeIT interface.
type Component struct {
	ID               ComponentID         `json:"id"`
	Name             string              `json:"name"`
	Image            null.NullableString `json:"image,omitempty"`
	Serial           null.NullableString `json:"serial,omitempty"`
	Location         *Location           `json:"location,omitempty"`
	Qty              null.NullableInt    `json:"qty,omitempty"`
	MinAmt           null.NullableInt    `json:"min_amt,omitempty"`
	Category         *Category           `json:"category,omitempty"`
	Supplier         *Supplier           `json:"supplier,omitempty"`
	OrderNumber      string              `json:"order_number,omitempty"`
	PurchaseDate     Datetime            `json:"purchase_date,omitempty"`
	PurchaseCost     string              `json:"purchase_cost,omitempty"`
	Remaining        int                 `json:"remaining,omitempty"`
	Company          *Company            `json:"company,omitempty"`
	Notes            null.NullableString `json:"notes,omitempty"`
	CreatedBy        *User               `json:"created_by,omitempty"`
	CreatedAt        Datetime            `json:"created_at,omitempty"`
	UpdatedAt        Datetime            `json:"updated_at,omitempty"`
	UserCanCheckout  bool                `json:"user_can_checkout,omitempty"`
	AvailableActions *AvailableActions   `json:"available_actions,omitempty"`
}
