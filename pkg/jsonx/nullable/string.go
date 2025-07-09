package nullable

import (
	"bytes"
	"encoding/json"
)

type String struct {
	Value string
	Null  bool
}

func NewString(value string) String {
	return String{
		Value: value,
		Null:  false,
	}
}

func (s *String) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, null) {
		s.Null = true
		return nil
	}

	s.Null = false
	return json.Unmarshal(v, &s.Value)
}

func (s String) MarshalJSON() ([]byte, error) {
	if s.Null {
		return null, nil
	}

	return json.Marshal(s.Value)
}
