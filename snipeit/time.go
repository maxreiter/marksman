package snipeit

import (
	"bytes"
	"encoding/json"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, []byte("null")) {
		return nil
	}

	var datetime struct {
		Datetime string `json:"datetime"`
	}

	if err := json.Unmarshal(v, &datetime); err != nil {
		return err
	}

	timestamp, err := time.Parse(time.DateTime, datetime.Datetime)
	if err != nil {
		return err
	}

	*t = Time{timestamp}
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte(t.Format(time.DateTime)), nil
}
