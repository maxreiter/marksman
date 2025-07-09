package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Accessory struct {
	ID                int32           `json:"id"`
	Name              string          `json:"name"`
	Image             nullable.String `json:"image"`
	Company           *Company        `json:"company"`
	Manufacturer      *Manufacturer   `json:"manufacturer"`
	Supplier          *Supplier       `json:"supplier"`
	ModelNumber       nullable.String `json:"model_number"`
	Category          Category        `json:"category"`
	Location          *Location       `json:"location"`
	Notes             nullable.String `json:"notes"`
	Quantity          int             `json:"qty"`
	PurchaseDate      Time            `json:"purchase_date"`
	PurchaseCost      nullable.String `json:"purchase_cost"`
	OrderNumber       nullable.String `json:"order_number"`
	MinimumQuantity   nullable.Int    `json:"min_qty"`
	MinimumAmount     nullable.Int    `json:"min_amt"`
	RemainingQuantity nullable.Int    `json:"remaining_qty"`
	Remaining         nullable.Int    `json:"remaining"`
	CheckoutsCount    int             `json:"checkouts_count"`
	CreatedBy         *User           `json:"created_by"`
	CreatedAt         Time            `json:"created_at"`
	UpdatedAt         Time            `json:"updated_at"`
	// AvailableActions
	UserCanCheckout bool `json:"user_can_checkout"`
}
