package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type FieldFormatType string

const (
	FieldFormatTypeAny          FieldFormatType = "ANY"
	FieldFormatTypeCustomRegex  FieldFormatType = "CUSTOM REGEX"
	FieldFormatTypeAlpha        FieldFormatType = "ALPHA"
	FieldFormatTypeAlphaDash    FieldFormatType = "ALPHA-DASH"
	FieldFormatTypeNumeric      FieldFormatType = "NUMERIC"
	FieldFormatTypeAlphaNumeric FieldFormatType = "ALPHA-NUMERIC"
	FieldFormatTypeEmail        FieldFormatType = "EMAIL"
	FieldFormatTypeDate         FieldFormatType = "DATE"
	FieldFormatTypeURL          FieldFormatType = "URL"
	FieldFormatTypeIP           FieldFormatType = "IP"
	FieldFormatTypeIPv4         FieldFormatType = "IPV4"
	FieldFormatTypeIPv6         FieldFormatType = "IPV6"
	FieldFormatTypeMAC          FieldFormatType = "MAC"
	FieldFormatTypeBoolean      FieldFormatType = "BOOLEAN"
)

type FieldElementType string

const (
	FieldElementTypeText     FieldElementType = "text"
	FieldElementTypeList     FieldElementType = "list"
	FieldElementTypeTextarea FieldElementType = "textarea"
	FieldElementTypeCheckbox FieldElementType = "checkbox"
	FieldElementTypeRadio    FieldElementType = "radio"
)

type Field struct {
	ID                 int32            `json:"id"`
	Name               string           `json:"name"`
	DatabaseColumnName string           `json:"db_column_name"`
	Format             FieldFormatType  `json:"format"`
	FieldValues        nullable.String  `json:"field_values"`
	FieldValuesArray   []string         `json:"field_values_array"`
	Type               FieldElementType `json:"type"`
	Required           bool             `json:"required"`
	DisplayInUserView  bool             `json:"display_in_user_view"`
	AutoAddToFieldsets bool             `json:"auto_add_to_fieldsets"`
	ShowInListView     bool             `json:"show_in_listview"`
	DisplayCheckin     bool             `json:"display_checkin"`
	DisplayCheckout    bool             `json:"display_checkout"`
	DisplayAudit       bool             `json:"display_audit"`
	CreatedAt          Time             `json:"created_at"`
	UpdatedAt          Time             `json:"updated_at"`
}
