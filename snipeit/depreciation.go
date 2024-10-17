package snipeit

type Depreciation struct {
	ID               DepreciationID    `json:"id"`
	Name             string            `json:"name"`
	Months           string            `json:"months,omitempty"`
	DepreciationMin  string            `json:"depreciation_min,omitempty"`
	AssetsCount      int               `json:"assets_count,omitempty"`
	LicensesCount    int               `json:"licenses_count,omitempty"`
	CreatedBy        *User             `json:"created_by,omitempty"`
	CreatedAt        Datetime          `json:"created_at,omitempty"`
	UpdatedAt        Datetime          `json:"updated_at,omitempty"`
	AvailableActions *AvailableActions `json:"available_actions,omitempty"`
}
