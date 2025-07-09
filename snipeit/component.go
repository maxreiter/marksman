package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Component struct {
	ID              int32           `json:"id"`
	Name            string          `json:"name"`
	Image           nullable.String `json:"image"`
	Serial          nullable.String `json:"serial"`
	Location        *Location       `json:"location"`
	Quantity        nullable.Int    `json:"qty"`
	MinimumAmount   nullable.Int    `json:"min_amt"`
	Category        *Category       `json:"category"`
	Supplier        *Supplier       `json:"supplier"`
	Manufacturer    *Manufacturer   `json:"manufacturer"`
	ModelNumber     nullable.String `json:"model_number"`
	OrderNumber     nullable.String `json:"order_number"`
	PurchaseDate    Date            `json:"purchase_date"`
	PurchaseCost    nullable.String `json:"purchase_cost"`
	Remaining       nullable.Int    `json:"remaining"`
	Company         *Company        `json:"company"`
	Notes           nullable.String `json:"notes"`
	CreatedBy       *User           `json:"created_by"`
	CreatedAt       Time            `json:"created_at"`
	UpdatedAt       Time            `json:"updated_at"`
	UserCanCheckout int             `json:"user_can_checkout"`
	// AvailableActions
}
