package marksman

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Error is an error returned by the SnipeIT API.
type Error struct {
	Message          string              `json:"message"`
	ValidationErrors map[string][]string `json:"validation_errors"`
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (e *Error) UnmarshalJSON(b []byte) error {
	var res struct {
		Messages any `json:"messages"`
	}

	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}

	switch data := res.Messages.(type) {
	case string:
		e.Message = data
	case map[string]interface{}:
		e.ValidationErrors = make(map[string][]string)

		for key, val := range data {
			iv, ok := val.([]interface{})
			if !ok {
				return errors.New("unknown type passed in validation errors")
			}

			e.ValidationErrors[key] = make([]string, 0)

			for _, v := range iv {
				switch kind := v.(type) {
				case string:
					e.ValidationErrors[key] = append(e.ValidationErrors[key], kind)
				case int:
					i := strconv.FormatInt(int64(kind), 10)
					e.ValidationErrors[key] = append(e.ValidationErrors[key], i)
				case bool:
					b := strconv.FormatBool(kind)
					e.ValidationErrors[key] = append(e.ValidationErrors[key], b)
				default:
					return errors.New("unknown type in validation errors slice")
				}
			}
		}
	}

	return nil
}

// Error implements the error interface.
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
					v = strings.ToLower(v)
					if strings.HasSuffix(v, ".") {
						v = v[:len(v)-1]
					}

					sb.WriteString(v + "; ")
				}
			}
		}

		return sb.String()[:sb.Len()-2]
	}

	return "Unhandled error"
}

// HTTPError is an error containing information pertaining to the HTTP request.
type HTTPError struct {
	Code int
}

// Error implements the error interface.
func (e HTTPError) Error() string {
	return fmt.Sprintf("snipe: unexpected status code returned by API: %d", e.Code)
}
