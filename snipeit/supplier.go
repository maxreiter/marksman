package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

// Supplier represents a supplier in the SnipeIT interface.
type Supplier struct {
	ID               SupplierID          `json:"id"`
	Name             string              `json:"name"`
	Image            null.NullableString `json:"image,omitempty"`
	URL              string              `json:"url,omitempty"`
	Address          string              `json:"address,omitempty"`
	Address2         string              `json:"address2,omitempty"`
	City             string              `json:"city,omitempty"`
	State            string              `json:"state,omitempty"`
	Country          string              `json:"country,omitempty"`
	Zip              string              `json:"zip,omitempty"`
	Fax              string              `json:"fax,omitempty"`
	Phone            string              `json:"phone,omitempty"`
	Contact          string              `json:"contact,omitempty"`
	AssetsCount      int                 `json:"assets_count,omitempty"`
	AccessoriesCount int                 `json:"accessories_count,omitempty"`
	LicensesCount    int                 `json:"licenses_count,omitempty"`
	ConsumablesCount int                 `json:"consumables_count,omitempty"`
	ComponentsCount  int                 `json:"components_count,omitempty"`
	Notes            null.NullableString `json:"notes,omitempty"`
	CreatedAt        Datetime            `json:"created_at,omitempty"`
	UpdatedAt        Datetime            `json:"updated_at,omitempty"`
	AvailableActions *AvailableActions   `json:"available_actions,omitempty"`
}
