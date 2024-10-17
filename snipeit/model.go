package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type Model struct {
	ID                    ModelID             `json:"id"`
	Name                  string              `json:"name"`
	Manufacturer          *Manufacturer       `json:"manufacturer,omitempty"`
	Image                 null.NullableString `json:"image,omitempty"`
	ModelNumber           string              `json:"model_number,omitempty"`
	MinAmt                int                 `json:"min_amt,omitempty"`
	Depreciation          *Depreciation       `json:"depreciation,omitempty"`
	AssetsCount           int                 `json:"assets_count,omitempty"`
	Category              *Category           `json:"category,omitempty"`
	Fieldset              *Fieldset           `json:"fieldset,omitempty"`
	DefaultFieldsetValues []*Field            `json:"default_fieldset_values,omitempty"`
	EOL                   string              `json:"eol,omitempty"`
	Requestable           bool                `json:"requestable,omitempty"`
	Notes                 null.NullableString `json:"notes,omitempty"`
	CreatedAt             Datetime            `json:"created_at,omitempty"`
	UpdatedAt             Datetime            `json:"updated_at,omitempty"`
	DeletedAt             Datetime            `json:"deleted_at,omitempty"`
	AvailableActions      *AvailableActions   `json:"available_actions,omitempty"`
}
