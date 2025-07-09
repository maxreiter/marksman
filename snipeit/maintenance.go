package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type AssetMaintenanceType string

const (
	AssetMaintenanceTypeMaintenance     AssetMaintenanceType = "maintenance"
	AssetMaintenanceTypeRepair          AssetMaintenanceType = "repair"
	AssetMaintenanceTypePATTest         AssetMaintenanceType = "pat test"
	AssetMaintenanceTypeUpgrade         AssetMaintenanceType = "upgrade"
	AssetMaintenanceTypeHardwareSupport AssetMaintenanceType = "hardware support"
	AssetMaintenanceTypeSoftwareSupport AssetMaintenanceType = "software support"
)

type Maintenance struct {
	ID                   int32                `json:"id"`
	Asset                *Asset               `json:"asset"`
	Model                *Model               `json:"model"`
	StatusLabel          *StatusLabel         `json:"status_label"`
	Company              *Company             `json:"company"`
	Title                nullable.String      `json:"title"`
	Location             *Location            `json:"location"`
	RTDLocation          *Location            `json:"rtd_location"`
	Notes                nullable.String      `json:"notes"`
	Supplier             *Supplier            `json:"supplier"`
	Cost                 nullable.String      `json:"cost"`
	AssetMaintenanceType AssetMaintenanceType `json:"asset_maintenance_type"`
	StartDate            Date                 `json:"start_date"`
	CompletionDate       Date                 `json:"completion_date"`
	CreatedBy            *User                `json:"created_by"`
	CreatedAt            Time                 `json:"created_at"`
	UpdatedAt            Time                 `json:"updated_at"`
	IsWarranty           bool                 `json:"is_warranty"`
}
