package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type Manufacturer struct {
	ID                ManufacturerID      `json:"id"`
	Name              string              `json:"name"`
	URL               string              `json:"url,omitempty"`
	Image             null.NullableString `json:"image,omitempty"`
	SupportURL        string              `json:"support_url,omitempty"`
	WarrantyLookupURL string              `json:"warranty_lookup_url,omitempty"`
	SupportPhone      string              `json:"support_phone,omitempty"`
	SupportEmail      string              `json:"support_email,omitempty"`
	AssetsCount       int                 `json:"assets_count,omitempty"`
	LicensesCount     int                 `json:"licenses_count,omitempty"`
	ConsumablesCount  int                 `json:"consumables_count,omitempty"`
	AccessoriesCount  int                 `json:"accessories_count,omitempty"`
	CreatedBy         *User               `json:"created_by,omitempty"`
	CreatedAt         Datetime            `json:"created_at,omitempty"`
	UpdatedAt         Datetime            `json:"updated_at,omitempty"`
	DeletedAt         Datetime            `json:"deleted_at,omitempty"`
	AvailableActions  *AvailableActions   `json:"available_actions,omitempty"`
}
