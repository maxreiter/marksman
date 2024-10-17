package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type LicenseSeat struct {
	ID               LicenseSeatID       `json:"id"`
	LicenseID        LicenseID           `json:"license_id"`
	AssignedUser     *User               `json:"assigned_user,omitempty"`
	AssignedAsset    *Asset              `json:"assigned_asset,omitempty"`
	Location         *Location           `json:"location,omitempty"`
	Reassignable     bool                `json:"reassignable,omitempty"`
	Notes            null.NullableString `json:"notes,omitempty"`
	UserCanCheckout  bool                `json:"user_can_checkout,omitempty"`
	AvailableActions *AvailableActions   `json:"available_actions,omitempty"`
}

type License struct {
	ID                  LicenseID           `json:"id"`
	Name                string              `json:"name"`
	Company             *Company            `json:"company,omitempty"`
	Manufacturer        *Manufacturer       `json:"manufacturer,omitempty"`
	ProductKey          string              `json:"product_key,omitempty"`
	OrderNumber         null.NullableString `json:"order_number,omitempty"`
	PurchaseOrder       null.NullableString `json:"purchase_order,omitempty"`
	TerminationDate     Datetime            `json:"termination_date,omitempty"`
	Depreciation        *Depreciation       `json:"depreciation,omitempty"`
	PurchaseCost        string              `json:"purchase_cost,omitempty"`
	PurchaseCostNumeric int                 `json:"purchase_cost_numeric,omitempty"`
	Notes               null.NullableString `json:"notes,omitempty"`
	ExpirationDate      Datetime            `json:"expiration_date,omitempty"`
	Seats               int                 `json:"seats,omitempty"`
	FreeSeatsCount      int                 `json:"free_seats_count,omitempty"`
	MinAmt              null.NullableInt    `json:"min_amt,omitempty"`
	LicenseName         null.NullableString `json:"license_name,omitempty"`
	LicenseEmail        null.NullableString `json:"license_email,omitempty"`
	Reassignable        bool                `json:"reassignable,omitempty"`
	Maintained          bool                `json:"maintained,omitempty"`
	Supplier            *Supplier           `json:"supplier,omitempty"`
	Category            *Category           `json:"category,omitempty"`
	CreatedBy           *User               `json:"created_by,omitempty"`
	CreatedAt           Datetime            `json:"created_at,omitempty"`
	UpdatedAt           Datetime            `json:"updated_at,omitempty"`
	DeletedAt           Datetime            `json:"deleted_at,omitempty"`
	UserCanCheckout     bool                `json:"user_can_checkout,omitempty"`
	AvailableActions    *AvailableActions   `json:"available_actions,omitempty"`
}
