package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Depreciation struct {
	ID                  int32           `json:"id"`
	Name                string          `json:"name"`
	Months              nullable.String `json:"months"`
	DepreciationMinimum nullable.String `json:"depreciation_min"`
	AssetsCount         int             `json:"assets_count"`
	ModelsCount         int             `json:"models_count"`
	LicensesCount       int             `json:"licenses_count"`
	CreatedBy           *User           `json:"created_by"`
	CreatedAt           Time            `json:"created_at"`
	UpdatedAt           Time            `json:"updated_at"`
	// AvailableActions
}
