package snipeit

import (
	"bytes"
	"encoding/json"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(v []byte) error {
	if bytes.EqualFold(v, []byte("null")) {
		return nil
	}

	var date struct {
		Date string `json:"date"`
	}

	if err := json.Unmarshal(v, &date); err != nil {
		return err
	}

	timestamp, err := time.Parse(time.DateOnly, date.Date)
	if err != nil {
		return err
	}

	*d = Date{timestamp}
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}

	return []byte(d.Format(time.DateOnly)), nil
}
