package nullable

import (
	"bytes"
	"encoding/json"
)

// Int is a nullable int with support for JSON
// serialization.
type Int struct {
	Value int
	Null  bool
}

// NewInt creates a new nullable int.
func NewInt(value int) Int {
	return Int{
		Value: value,
		Null:  false,
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (t Int) MarshalJSON() ([]byte, error) {
	if t.Null {
		return bytes.Clone(null), nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (t Int) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, null) {
		t.Null = true
		return nil
	}

	t.Null = false
	return json.Unmarshal(v, &t.Value)
}

// Int8 is a nullable int8 with support for JSON
// serialization.
type Int8 struct {
	Value int8
	Null  bool
}

// NewInt8 creates a new nullable int8.
func NewInt8(value int8) Int8 {
	return Int8{
		Value: value,
		Null:  false,
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (t Int8) MarshalJSON() ([]byte, error) {
	if t.Null {
		return bytes.Clone(null), nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (t Int8) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, null) {
		t.Null = true
		return nil
	}

	t.Null = false
	return json.Unmarshal(v, &t.Value)
}

// Int16 is a nullable int16 with support for JSON
// serialization.
type Int16 struct {
	Value int16
	Null  bool
}

// NewInt16 creates a new nullable int16.
func NewInt16(value int16) Int16 {
	return Int16{
		Value: value,
		Null:  false,
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (t Int16) MarshalJSON() ([]byte, error) {
	if t.Null {
		return bytes.Clone(null), nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (t Int16) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, null) {
		t.Null = true
		return nil
	}

	t.Null = false
	return json.Unmarshal(v, &t.Value)
}

// Int32 is a nullable int32 with support for JSON
// serialization.
type Int32 struct {
	Value int32
	Null  bool
}

// NewInt32 creates a new nullable int32.
func NewInt32(value int32) Int32 {
	return Int32{
		Value: value,
		Null:  false,
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (t Int32) MarshalJSON() ([]byte, error) {
	if t.Null {
		return bytes.Clone(null), nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (t Int32) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, null) {
		t.Null = true
		return nil
	}

	t.Null = false
	return json.Unmarshal(v, &t.Value)
}

// Int64 is a nullable int64 with support for JSON
// serialization.
type Int64 struct {
	Value int64
	Null  bool
}

// NewInt64 creates a new nullable int64.
func NewInt64(value int64) Int64 {
	return Int64{
		Value: value,
		Null:  false,
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (t Int64) MarshalJSON() ([]byte, error) {
	if t.Null {
		return bytes.Clone(null), nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (t Int64) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, null) {
		t.Null = true
		return nil
	}

	t.Null = false
	return json.Unmarshal(v, &t.Value)
}
