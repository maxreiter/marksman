package nullable

import (
	"bytes"
	"encoding/json"
)

type Bool struct {
	Value bool
	Null  bool
}

var (
	True  = Bool{true, false}
	False = Bool{false, false}
)

func (b *Bool) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, null) {
		b.Null = true
		return nil
	}

	b.Null = false
	return json.Unmarshal(v, &b.Value)
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b.Null {
		return bytes.Clone(null), nil
	}

	return json.Marshal(b.Value)
}
