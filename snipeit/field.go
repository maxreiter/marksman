package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

// FieldFormat represents the format of a [Field].
type FieldFormat string

const (
	FieldFormatAny          FieldFormat = "ANY"
	FieldFormatAlpha        FieldFormat = "ALPHA"
	FieldFormatAlphaDash    FieldFormat = "ALPHA-DASH"
	FieldFormatNumeric      FieldFormat = "NUMERIC"
	FieldFormatAlphaNumeric FieldFormat = "ALPHA-NUMERIC"
	FieldFormatEmail        FieldFormat = "EMAIL"
	FieldFormatDate         FieldFormat = "DATE"
	FieldFormatURL          FieldFormat = "URL"
	FieldFormatIP           FieldFormat = "IP"
	FieldFormatIPV4         FieldFormat = "IPV4"
	FieldFormatIPV6         FieldFormat = "IPV6"
	FieldFormatMAC          FieldFormat = "MAC"
	FieldFormatBoolean      FieldFormat = "BOOLEAN"
	FieldFormatCustom       FieldFormat = "CUSTOM"
)

// Field represents a custom field seen on the SnipeIT interface.
type Field struct {
	ID                 FieldID             `json:"id"`
	Name               string              `json:"name"`
	DBColumnName       string              `json:"db_column_name,omitempty"`
	Format             FieldFormat         `json:"format,omitempty"`
	FieldValues        null.NullableString `json:"field_values,omitempty"`
	FieldValuesArray   []string            `json:"field_values_array,omitempty"`
	Type               string              `json:"type,omitempty"`
	Required           bool                `json:"required,omitempty"`
	DisplayInUserView  bool                `json:"display_in_user_view,omitempty"`
	AutoAddToFieldsets bool                `json:"auto_add_to_fieldsets,omitempty"`
	ShowInListview     bool                `json:"show_in_listview,omitempty"`
	CreatedAt          Datetime            `json:"created_at,omitempty"`
	UpdatedAt          Datetime            `json:"updated_at,omitempty"`
}

// Fieldset represents a set of [Field] seen on the SnipeIT interface.
type Fieldset struct {
	ID        FieldsetID `json:"id"`
	Name      string     `json:"name"`
	Fields    []*Field   `json:"fields,omitempty"`
	Models    []*Model   `json:"models,omitempty"`
	CreatedAt Datetime   `json:"created_at,omitempty"`
	UpdatedAt Datetime   `json:"updated_at,omitempty"`
}
