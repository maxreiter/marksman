package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type Company struct {
	ID               CompanyID           `json:"company"`
	Name             string              `json:"name"`
	Phone            null.NullableString `json:"phone,omitempty"`
	Fax              null.NullableString `json:"fax,omitempty"`
	Email            null.NullableString `json:"email,omitempty"`
	Image            null.NullableString `json:"image,omitempty"`
	AssetsCount      int                 `json:"assets_count,omitempty"`
	LicensesCount    int                 `json:"licenses_count,omitempty"`
	ConsumablesCount int                 `json:"consumables_count,omitempty"`
	ComponentsCount  int                 `json:"components_count,omitempty"`
	UsersCount       int                 `json:"users_count,omitempty"`
	CreatedBy        *User               `json:"created_by,omitempty"`
	CreatedAt        Datetime            `json:"created_at,omitempty"`
	UpdatedAt        Datetime            `json:"updated_at,omitempty"`
	AvailableActions *AvailableActions   `json:"available_actions,omitempty"`
}
