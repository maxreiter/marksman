package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Company struct {
	ID               int32           `json:"id"`
	Name             string          `json:"name"`
	Phone            nullable.String `json:"phone"`
	Fax              nullable.String `json:"fax"`
	Email            nullable.String `json:"email"`
	Image            nullable.String `json:"image"`
	AssetsCount      int             `json:"assets_count"`
	LicensesCount    int             `json:"licenses_count"`
	AccessoriesCount int             `json:"accessories_count"`
	ConsumablesCount int             `json:"consumables_count"`
	ComponentsCount  int             `json:"components_count"`
	UsersCount       int             `json:"users_count"`
	CreatedBy        *User           `json:"created_by"`
	Notes            nullable.String `json:"notes"`
	CreatedAt        Time            `json:"created_at"`
	UpdatedAt        Time            `json:"updated_at"`
	// TODO: AvailableActions
}
