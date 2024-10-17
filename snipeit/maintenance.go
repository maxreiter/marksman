package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type AssetMaintenanceType string

const (
	MaintenanceGeneral         AssetMaintenanceType = "Maintenance"
	MaintenanceRepair          AssetMaintenanceType = "Repair"
	MaintenancePATTest         AssetMaintenanceType = "PAT test"
	MaintenanceUpgrade         AssetMaintenanceType = "Upgrade"
	MaintenanceHardwareSupport AssetMaintenanceType = "Hardware Support"
	MaintenanceSoftwareSupport AssetMaintenanceType = "Software Support"
)

type Maintenance struct {
	ID                   MaintenanceID        `json:"id"`
	Asset                *Asset               `json:"asset"`
	Model                *Model               `json:"model"`
	StatusLabel          *StatusLabel         `json:"status_label"`
	Company              *Company             `json:"company"`
	Title                null.NullableString  `json:"title"`
	Location             *Location            `json:"location"`
	RTDLocation          *Location            `json:"rtd_location"`
	Notes                null.NullableString  `json:"notes"`
	Supplier             *Supplier            `json:"supplier"`
	Cost                 string               `json:"cost"`
	AssetMaintenanceType AssetMaintenanceType `json:"asset_maintenance_type"`
	StartDate            Datetime             `json:"start_date"`
	AssetMaintenanceTime int                  `json:"asset_maintenance_time"`
	CompletionDate       Datetime             `json:"completion_date"`
	CreatedBy            *User                `json:"created_by"`
	CreatedAt            Datetime             `json:"created_at"`
	UpdatedAt            Datetime             `json:"updated_at"`
	IsWarranty           bool                 `json:"is_warranty"`
	AvailableActions     *AvailableActions    `json:"available_actions"`
}
