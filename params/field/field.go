package field

import (
	"io"
	"net/url"

	params "github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type ElementType string

const (
	ElementText     ElementType = "text"
	ElementTextarea ElementType = "textarea"
	ElementCheckbox ElementType = "checkbox"
	ElementRadio    ElementType = "radio"
	ElementListbox  ElementType = "listbox"
)

type RequestOptions struct {
	*params.BaseResolver

	Name           string             `json:"name,omitempty" url:"-"`
	Element        ElementType        `json:"element,omitempty" url:"-"`
	FieldValues    string             `json:"field_values,omitempty" url:"-"`
	ShowInEmail    bool               `json:"show_in_email,omitempty" url:"-"`
	Format         string             `json:"format,omitempty" url:"-"`
	FieldEncrypted bool               `json:"field_encrypted,omitempty" url:"-"`
	HelpText       string             `json:"help_text,omitempty" url:"-"`
	FieldsetID     snipeit.FieldsetID `json:"fieldset_id,omitempty" url:"-"`
}

func (ro *RequestOptions) Values() (url.Values, error) {
	return ro.BaseResolver.Values()
}

func (ro *RequestOptions) Marshal() (io.Reader, error) {
	return ro.BaseResolver.Marshal()
}

type RequestOption func(*RequestOptions)

func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

func Element(element ElementType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Element = element
	}
}

func FieldValues(values string) RequestOption {
	return func(ro *RequestOptions) {
		ro.FieldValues = values
	}
}

func ShowInEmail() RequestOption {
	return func(ro *RequestOptions) {
		ro.ShowInEmail = true
	}
}

func Format(format string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Format = format
	}
}

func FieldEncrypted() RequestOption {
	return func(ro *RequestOptions) {
		ro.FieldEncrypted = true
	}
}

func HelpText(text string) RequestOption {
	return func(ro *RequestOptions) {
		ro.HelpText = text
	}
}

func FieldsetID(id snipeit.FieldsetID) RequestOption {
	return func(ro *RequestOptions) {
		ro.FieldsetID = id
	}
}
