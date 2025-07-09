package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Supplier struct {
	ID               int32           `json:"id"`
	Name             string          `json:"name"`
	Image            nullable.String `json:"image"`
	URL              nullable.String `json:"url"`
	Address          nullable.String `json:"address"`
	Address2         nullable.String `json:"address2"`
	City             nullable.String `json:"city"`
	State            nullable.String `json:"state"`
	Zip              nullable.String `json:"zip"`
	Fax              nullable.String `json:"fax"`
	Phone            nullable.String `json:"phone"`
	Email            nullable.String `json:"email"`
	Contact          nullable.String `json:"contact"`
	AssetsCount      int             `json:"assets_count"`
	AccessoriesCount int             `json:"accessories_count"`
	LicensesCount    int             `json:"licenses_count"`
	ConsumablesCount int             `json:"consumables_count"`
	ComponentsCount  int             `json:"components_count"`
	Notes            nullable.String `json:"notes"`
	CreatedAt        Time            `json:"created_at"`
	CreatedBy        *User           `json:"created_by"`
	UpdatedAt        Time            `json:"updated_at"`
	// TODO: AvailableActions
}
