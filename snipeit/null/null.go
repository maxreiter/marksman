// Package null provides nullable primatives for use in JSON bodies.
package null

import "encoding/json"

// NullableStringData contains data for nullable strings.
type NullableStringData struct {
	Val  string
	Init bool
}

// NullableString represents a null string.
type NullableString = *NullableStringData

// NullString is a null string.
var NullString = &NullableStringData{}

// NewNullableString creates a new nullable string from the given string.
func NewNullableString(str string) NullableString {
	return &NullableStringData{
		Val:  str,
		Init: true,
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (s NullableStringData) MarshalJSON() ([]byte, error) {
	if !s.Init {
		return []byte("null"), nil
	}

	return json.Marshal(s.Val)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (s *NullableStringData) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*s = *NullString
		return nil
	}

	s.Init = true
	return json.Unmarshal(b, &s.Val)
}

// NullableIntData contains data for nullable integers.
type NullableIntData struct {
	Val  int
	Init bool
}

// NullableInt represents a null integer.
type NullableInt = *NullableIntData

// NullInt is a null integer.
var NullInt = &NullableIntData{}

// NewNullableInt creates a new nullable integer from the given integer.
func NewNullableInt(i int) NullableInt {
	return &NullableIntData{
		Val:  i,
		Init: true,
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i NullableIntData) MarshalJSON() ([]byte, error) {
	if !i.Init {
		return []byte("null"), nil
	}

	return json.Marshal(i.Val)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *NullableIntData) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*i = *NullInt
		return nil
	}

	i.Init = true
	return json.Unmarshal(b, &i.Val)
}

// NullableInt32Data contains data for nullable int32s.
type NullableInt32Data struct {
	Val  int32
	Init bool
}

// NullableInt32 represents a null int32.
type NullableInt32 = *NullableInt32Data

// NullInt32 is a null int32.
var NullInt32 = &NullableInt32Data{}

// NewNullableInt32 creates a new null int32 from the given int32.
func NewNullableInt32(i int32) NullableInt32 {
	return &NullableInt32Data{
		Val:  i,
		Init: true,
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i NullableInt32Data) MarshalJSON() ([]byte, error) {
	if !i.Init {
		return []byte("null"), nil
	}

	return json.Marshal(i.Val)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *NullableInt32Data) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*i = *NullInt32
		return nil
	}

	i.Init = true
	return json.Unmarshal(b, &i.Val)
}
