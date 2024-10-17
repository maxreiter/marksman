package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type StatusLabelType string

const (
	StatusLabelTypePending      StatusLabelType = "pending"
	StatusLabelTypeArchived     StatusLabelType = "archived"
	StatusLabelTypeUndeployable StatusLabelType = "undeployable"
)

type StatusLabel struct {
	ID               StatusLabelID       `json:"id"`
	Name             string              `json:"name"`
	Type             StatusLabelType     `json:"type,omitempty"`
	Color            null.NullableString `json:"color,omitempty"`
	ShowInNav        bool                `json:"show_in_nav,omitempty"`
	DefaultLabel     bool                `json:"default_label,omitempty"`
	AssetsCount      int                 `json:"assets_count,omitempty"`
	Notes            null.NullableString `json:"notes,omitempty"`
	CreatedBy        *User               `json:"created_by,omitempty"`
	CreatedAt        Datetime            `json:"created_at,omitempty"`
	UpdatedAt        Datetime            `json:"updated_at,omitempty"`
	AvailableActions *AvailableActions   `json:"available_actions,omitempty"`
}
