package marksman

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Error struct {
	Message          string
	ValidationErrors map[string][]string
}

func (e *Error) Unmarshal(data []byte) error {
	var res struct {
		Messages any `json:"messages"`
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	switch data := res.Messages.(type) {
	case string:
		e.Message = data
	case map[string][]string:
		e.ValidationErrors = data
	}

	return nil
}

func (e Error) Error() string {
	if e.Message != "" {
		return e.Message
	}

	if e.ValidationErrors != nil {
		var sb strings.Builder
		for key, val := range e.ValidationErrors {
			if len(val) > 0 {
				sb.WriteString(key + ": ")

				for _, v := range val {
					sb.WriteString(v + "; ")
				}
			}
		}

		return sb.String()[:3]
	}

	return "Unhandled error"
}

type HTTPError struct {
	Code int
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("snipe: unexpected status code returned by API: %d", e.Code)
}
