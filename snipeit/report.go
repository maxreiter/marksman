package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type ActionType string

const (
	ActionAddSeats    ActionType = "add seats"
	ActionCheckinFrom ActionType = "checkin from"
	ActionCheckout    ActionType = "checkout"
	ActionUpdate      ActionType = "update"
)

type ActionLogFile struct {
	URL        string `json:"url"`
	Filename   string `json:"filename"`
	Inlineable bool   `json:"inlineable"`
}

type ActionLogItem struct {
	ID     ID                  `json:"id"`
	Name   string              `json:"name"`
	Type   string              `json:"type"`
	Serial null.NullableString `json:"serial"`
}

type ActionLogTarget struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ActionLog struct {
	ID              ActionLogID         `json:"id"`
	Icon            null.NullableString `json:"icon"`
	File            *ActionLogFile      `json:"file"`
	Item            *ActionLogItem      `json:"item"`
	Location        *Location           `json:"location"`
	CreatedAt       Datetime            `json:"created_at"`
	UpdatedAt       Datetime            `json:"updated_at"`
	NextAuditDate   Datetime            `json:"next_audit_date"`
	DaysToNextAudit int                 `json:"days_to_next_audit"`
	ActionType      ActionType          `json:"action_type"`
	CreatedBy       *User               `json:"created_by"`
	Target          *ActionLogTarget    `json:"target"`
	Note            null.NullableString `json:"note"`
	SignatureFile   null.NullableString `json:"signature_file"`
	RemoteIP        null.NullableString `json:"remote_ip"`
	UserAgent       null.NullableString `json:"user_agent"`
	ActionSource    null.NullableString `json:"action_source"`
	ActionDate      Datetime            `json:"action_date"`
}
