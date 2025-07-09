package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Consumable struct {
	ID              int32           `json:"id"`
	Name            string          `json:"name"`
	Image           nullable.String `json:"image"`
	Category        *Category       `json:"category"`
	Company         *Company        `json:"company"`
	ItemNumber      nullable.String `json:"item_no"`
	Location        *Location       `json:"location"`
	Manufacturer    *Manufacturer   `json:"manufacturer"`
	Supplier        *Supplier       `json:"supplier"`
	MinimumAmount   nullable.Int    `json:"min_amt"`
	ModelNumber     nullable.String `json:"model_number"`
	Remaining       nullable.Int    `json:"remaining"`
	OrderNumber     nullable.String `json:"order_number"`
	PurchaseCost    nullable.String `json:"purchase_cost"`
	PurchaseDate    Date            `json:"purchase_date"`
	Quantity        nullable.Int    `json:"qty"`
	Notes           nullable.String `json:"notes"`
	CreatedBy       *User           `json:"created_by"`
	CreatedAt       Time            `json:"created_at"`
	UpdatedAt       Time            `json:"updated_at"`
	UserCanCheckout bool            `json:"user_can_checkout"`
	//AvailableActions
}
