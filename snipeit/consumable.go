package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type Consumable struct {
	ID               ConsumableID        `json:"id"`
	Name             string              `json:"name"`
	Image            null.NullableString `json:"image"`
	Category         *Category           `json:"category"`
	Company          *Company            `json:"company"`
	ItemNo           string              `json:"item_no"`
	Location         *Location           `json:"location"`
	Manufacturer     *Manufacturer       `json:"manufacturer"`
	Supplier         *Supplier           `json:"supplier"`
	MinAmt           int                 `json:"min_amt"`
	ModelNumber      null.NullableString `json:"model_number"`
	Remaining        int                 `json:"remaining"`
	OrderNumber      string              `json:"order_number"`
	PurchaseCost     string              `json:"purchase_cost"`
	PurchaseDate     Datetime            `json:"purchase_date"`
	Qty              int                 `json:"qty"`
	Notes            null.NullableString `json:"notes"`
	CreatedBy        *User               `json:"created_by"`
	CreatedAt        Datetime            `json:"created_at"`
	UpdatedAt        Datetime            `json:"updated_at"`
	AvailableActions *AvailableActions   `json:"available_actions"`
}
