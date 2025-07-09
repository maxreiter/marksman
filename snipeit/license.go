package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type License struct {
	ID                  int32           `json:"id"`
	Name                string          `json:"name"`
	Manufacturer        *Manufacturer   `json:"manufacturer"`
	ProductKey          nullable.String `json:"product_key"`
	OrderNumber         nullable.String `json:"order_number"`
	PurchaseOrder       nullable.String `json:"purchase_order"`
	PurchaseDate        Date            `json:"purchase_date"`
	TerminationDate     Date            `json:"termination_date"`
	Depreciation        *Depreciation   `json:"depreciation"`
	PurchaseCost        nullable.String `json:"purchase_cost"`
	PurchaseCostNumeric nullable.String `json:"purchase_cost_numeric"`
	Notes               nullable.String `json:"notes"`
	ExpirationDate      Date            `json:"expiration_date"`
	Seats               int             `json:"seats"`
	FreeSeatsCount      int             `json:"free_seats_count"`
	Remaining           int             `json:"remaining"`
	MinimumAmount       nullable.Int    `json:"min_amt"`
	LicenseName         nullable.String `json:"license_name"`
	LicenseEmail        nullable.String `json:"license_email"`
	Reassignable        bool            `json:"reassignable"`
	Maintained          bool            `json:"maintained"`
	Supplier            *Supplier       `json:"supplier"`
	Category            *Category       `json:"category"`
	CreatedBy           *User           `json:"created_by"`
	CreatedAt           Time            `json:"created_at"`
	UpdatedAt           Time            `json:"updated_at"`
	DeletedAt           Time            `json:"deleted_at"`
	UserCanCheckout     bool            `json:"user_can_checkout"`
	// AvailableActions
}
