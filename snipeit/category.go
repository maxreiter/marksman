package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type CategoryType string

const (
	CategoryTypeAccessory  CategoryType = "accessory"
	CategoryTypeAsset      CategoryType = "asset"
	CategoryTypeConsumable CategoryType = "consumable"
	CategoryTypeComponent  CategoryType = "component"
	CategoryTypeLicense    CategoryType = "license"
)

type Category struct {
	ID                int32           `json:"id"`
	Name              string          `json:"name"`
	Image             nullable.String `json:"image"`
	CategoryType      CategoryType    `json:"category_type"`
	HasEULA           bool            `json:"has_eula"`
	EULA              nullable.String `json:"eula"`
	CheckinEmail      bool            `json:"checkin_email"`
	RequireAcceptance bool            `json:"require_acceptance"`
	ItemCount         int             `json:"item_count"`
	AssetsCount       int             `json:"assets_count"`
	AccessoriesCount  int             `json:"accessories_count"`
	ConsumablesCount  int             `json:"consumables_count"`
	ComponentsCount   int             `json:"components_count"`
	LicensesCount     int             `json:"licenses_count"`
	CreatedBy         *User           `json:"created_by"`
	Notes             nullable.String `json:"notes"`
	CreatedAt         Time            `json:"created_at"`
	UpdatedAt         Time            `json:"updated_at"`
	// AvailableActions
}
