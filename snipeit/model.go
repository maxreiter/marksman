package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Model struct {
	ID                    int32           `json:"id"`
	Name                  string          `json:"name"`
	Manufacturer          *Manufacturer   `json:"manufacturer"`
	Image                 nullable.String `json:"image"`
	ModelNumber           nullable.String `json:"model_number"`
	MinimumAmount         nullable.Int    `json:"min_amt"`
	Remaining             nullable.Int    `json:"remaining"`
	Depreciation          *Depreciation   `json:"depreciation"`
	AssetsCount           int             `json:"assets_count"`
	Category              *Category       `json:"category"`
	Fieldset              *Fieldset       `json:"fieldset"`
	DefaultFieldsetValues []Field         `json:"default_fieldset_values"`
	EOL                   nullable.String `json:"eol"`
	Requestable           bool            `json:"requestable"`
	Notes                 nullable.String `json:"notes"`
	CreatedBy             *User           `json:"created_by"`
	CreatedAt             Time            `json:"created_at"`
	UpdatedAt             Time            `json:"updated_at"`
	DeletedAt             Time            `json:"deleted_at"`
	// AvailableActions
}
