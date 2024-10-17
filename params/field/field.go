// Package field provides request configuration for methods of the marksman client.
package field

import (
	"github.com/maxreiter/marksman/snipeit"
)

// ElementType represents the different types of HTML elements a [snipeit.Field] may be.
type ElementType string

// The different types of HTML elements a [snipeit.Field] may be.
const (
	ElementText     ElementType = "text"
	ElementTextarea ElementType = "textarea"
	ElementCheckbox ElementType = "checkbox"
	ElementRadio    ElementType = "radio"
	ElementListbox  ElementType = "listbox"
)

// RequestOptions contains possible options for requests made to the /fields endpoints.
type RequestOptions struct {
	Name           string             `json:"name,omitempty" url:"-"`
	Element        ElementType        `json:"element,omitempty" url:"-"`
	FieldValues    string             `json:"field_values,omitempty" url:"-"`
	ShowInEmail    bool               `json:"show_in_email,omitempty" url:"-"`
	Format         string             `json:"format,omitempty" url:"-"`
	FieldEncrypted bool               `json:"field_encrypted,omitempty" url:"-"`
	HelpText       string             `json:"help_text,omitempty" url:"-"`
	FieldsetID     snipeit.FieldsetID `json:"fieldset_id,omitempty" url:"-"`
}

// RequestOption is used to configure [RequestOptions].
type RequestOption func(*RequestOptions)

// Name sets the name of a [snipeit.Field].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// Element sets the element type of a [snipeit.Field].
func Element(element ElementType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Element = element
	}
}

// FieldValues sets the values of a [snipeit.Field].
func FieldValues(values string) RequestOption {
	return func(ro *RequestOptions) {
		ro.FieldValues = values
	}
}

// ShowInEmail toggles the display of a [snipeit.Field] in emails on a request.
func ShowInEmail() RequestOption {
	return func(ro *RequestOptions) {
		ro.ShowInEmail = true
	}
}

// Formata sets the format of a [snipeit.Field].
func Format(format string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Format = format
	}
}

// FieldEncrypted toggles encryption on a [snipeit.Field].
func FieldEncrypted() RequestOption {
	return func(ro *RequestOptions) {
		ro.FieldEncrypted = true
	}
}

// HelpText sets the help text of a [snipeit.Field].
func HelpText(text string) RequestOption {
	return func(ro *RequestOptions) {
		ro.HelpText = text
	}
}

// FieldsetID sets the [snipeit.FieldsetID] of a [snipeit.Field].
func FieldsetID(id snipeit.FieldsetID) RequestOption {
	return func(ro *RequestOptions) {
		ro.FieldsetID = id
	}
}
