package snipeit

import (
	"encoding/json"
	"time"
)

type Datetime time.Time

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
