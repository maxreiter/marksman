package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Manufacturer struct {
	ID                int32           `json:"id"`
	Name              string          `json:"name"`
	URL               string          `json:"url"`
	Image             nullable.String `json:"image"`
	SupportURL        string          `json:"support_url"`
	WarrantyLookupURL string          `json:"warranty_lookup_url"`
	SupportPhone      string          `json:"support_phone"`
	SupportEmail      string          `json:"support_email"`
	AssetsCount       int             `json:"assets_count"`
	LicensesCount     int             `json:"licenses_count"`
	ConsumablesCount  int             `json:"consumables_count"`
	AccessoriesCount  int             `json:"accessories_count"`
	ComponentsCount   int             `json:"components_count"`
	Notes             nullable.String `json:"notes"`
	CreatedBy         *User           `json:"created_by"`
	CreatedAt         Time            `json:"created_at"`
	UpdatedAt         Time            `json:"updated_at"`
	DeletedAt         Time            `json:"deleted_at"`
	// AvailableActions
}
