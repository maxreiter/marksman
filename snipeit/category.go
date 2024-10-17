package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type CategoryType string

const (
	CategoryTypeAccessory  CategoryType = "accessory"
	CategoryTypeAsset      CategoryType = "asset"
	CategoryTypeConsumable CategoryType = "consumable"
	CategoryTypeComponent  CategoryType = "component"
	CategoryTypeLicense    CategoryType = "license"
)

type Category struct {
	ID                CategoryID          `json:"id"`
	Name              string              `json:"name"`
	Image             null.NullableString `json:"image,omitempty"`
	CategoryType      CategoryType        `json:"category_type,omitempty"`
	HasEULA           bool                `json:"has_eula,omitempty"`
	UseDefaultEULA    bool                `json:"use_default_eula,omitempty"`
	EULA              null.NullableString `json:"eula,omitempty"`
	CheckinEmail      bool                `json:"checkin_email,omitempty"`
	RequireAcceptance bool                `json:"require_acceptance,omitempty"`
	ItemCount         int                 `json:"item_count,omitempty"`
	AssetsCount       int                 `json:"assets_count,omitempty"`
	AccessoriesCount  int                 `json:"accessories_count,omitempty"`
	ConsumablesCount  int                 `json:"consumables_count,omitempty"`
	ComponentsCount   int                 `json:"components_count,omitempty"`
	LicensesCount     int                 `json:"licenses_count,omitempty"`
	CreatedBy         *User               `json:"created_by,omitempty"`
	CreatedAt         Datetime            `json:"created_at,omitempty"`
	UpdatedAt         Datetime            `json:"updated_at,omitempty"`
	AvailableActions  *AvailableActions   `json:"available_actions,omitempty"`
}
