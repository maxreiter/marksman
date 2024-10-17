package snipeit

import (
	"encoding/json"
	"time"
)

// Datetime wraps the [time.Time] type for JSON support.
type Datetime time.Time

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (d *Datetime) UnmarshalJSON(v []byte) error {
	if string(v) == "null" {
		*d = Datetime{}
		return nil
	}

	var val struct {
		Datetime string `json:"datetime"`
	}

	if err := json.Unmarshal(v, &val); err != nil {
		return err
	}

	r, err := time.Parse(time.DateTime, val.Datetime)
	if err != nil {
		return err
	}

	*d = Datetime(r)
	return nil
}
