package nullable

import (
	"bytes"
	"encoding/json"
)

// Float32 is a nullable float32 with support for JSON
// serialization.
type Float32 struct {
	Value float32
	Null  bool
}

// NewFloat32 creates a new nullable float32.
func NewFloat32(value float32) Float32 {
	return Float32{
		Value: value,
		Null:  false,
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (t Float32) MarshalJSON() ([]byte, error) {
	if t.Null {
		return bytes.Clone(null), nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (t Float32) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, null) {
		t.Null = true
		return nil
	}

	t.Null = false
	return json.Unmarshal(v, &t.Value)
}

// Float64 is a nullable float64 with support for JSON
// serialization.
type Float64 struct {
	Value float64
	Null  bool
}

// NewFloat64 creates a new nullable float64.
func NewFloat64(value float64) Float64 {
	return Float64{
		Value: value,
		Null:  false,
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (t Float64) MarshalJSON() ([]byte, error) {
	if t.Null {
		return bytes.Clone(null), nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (t Float64) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, null) {
		t.Null = true
		return nil
	}

	t.Null = false
	return json.Unmarshal(v, &t.Value)
}
