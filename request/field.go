package request

import "github.com/maxreiter/marksman/snipeit"

type CreateField struct {
	Name           string                   `json:"name"`
	Element        snipeit.FieldElementType `json:"element"`
	FieldValues    string                   `json:"field_values,omitempty"`
	ShowInEmail    bool                     `json:"show_in_email,omitempty"`
	Format         string                   `json:"format,omitempty"`
	FieldEncrypted bool                     `json:"field_encrypted,omitempty"`
	HelpText       string                   `json:"help_text,omitempty"`
}

type UpdateField struct {
	Name    string                   `json:"name"`
	Element snipeit.FieldElementType `json:"element"`
}

type PartiallyUpdateField struct {
	Name    string                   `json:"name,omitempty"`
	Element snipeit.FieldElementType `json:"element,omitempty"`
}

type AssociateField struct {
	FieldsetID int32 `json:"fieldset_id"`
}

type DisassociateField struct {
	FieldsetID int32 `json:"fieldset_id"`
}
