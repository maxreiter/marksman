package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type StatusLabelType string

const (
	StatusLabelTypeDeployable   = "deployable"
	StatusLabelTypePending      = "pending"
	StatusLabelTypeArchived     = "archived"
	StatusLabelTypeUndeployable = "undeployable"
)

type StatusLabel struct {
	ID           int32           `json:"id"`
	Name         string          `json:"name"`
	Type         StatusLabelType `json:"type"`
	Color        nullable.String `json:"color"`
	ShowInNav    bool            `json:"show_in_nav"`
	DefaultLabel bool            `json:"default_label"`
	Notes        nullable.String `json:"notes"`
	CreatedBy    *User           `json:"created_by"`
	CreatedAt    Time            `json:"created_at"`
	UpdatedAt    Time            `json:"updated_at"`
	// AvailableActions
}
