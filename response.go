package marksman

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

var (
	// ErrUnknownResponse is returned when an unknown response type is given in a response.
	ErrUnknownResponse = errors.New("marksman: unknown response type")
	// ErrUnknownMessageType is returned when a response contains a message with an unknown type.
	ErrUnknownMessageType = errors.New("marksman: response contains unknown message type")
)

// ResponseType denotes the different kinds of responses the SnipeIT API may send.
type ResponseType int

// The response types sent by the SnipeIT API.
const (
	ResponseSuccess ResponseType = iota + 1
	ResponseError
)

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (r *ResponseType) UnmarshalJSON(v []byte) error {
	val := strings.Trim(strings.ToLower(string(v)), `"`)

	switch val {
	case "success":
		*r = ResponseSuccess
	case "error":
		*r = ResponseError
	default:
		return ErrUnknownResponse
	}

	return nil
}

// Response is a response from a SnipeIT API endpoint.
type Response struct {
	Status  ResponseType    `json:"status"`
	Payload json.RawMessage `json:"payload"`

	Message          string              `json:"-"`
	ValidationErrors map[string][]string `json:"-"`
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (r *Response) UnmarshalJSON(v []byte) error {
	var res struct {
		Status   ResponseType    `json:"status"`
		Messages any             `json:"messages"`
		Payload  json.RawMessage `json:"payload"`
	}

	if err := json.Unmarshal(v, &res); err != nil {
		return err
	}

	r.Status = res.Status
	r.Payload = res.Payload

	switch messages := res.Messages.(type) {
	case string:
		r.Message = messages
	case map[string]interface{}:
		r.ValidationErrors = make(map[string][]string)

		for key, itf := range messages {
			slice, ok := itf.([]interface{})
			if !ok {
				panic("marksman: validation error is wrong type")
			}

			r.ValidationErrors[key] = make([]string, len(slice))

			for _, value := range slice {
				var result string

				switch newValue := value.(type) {
				case string:
					result = newValue
				case int:
					result = strconv.FormatInt(int64(newValue), 10)
				case bool:
					result = strconv.FormatBool(newValue)
				default:
					return ErrUnknownMessageType
				}

				r.ValidationErrors[key] = append(r.ValidationErrors[key], result)
			}
		}
	}

	return nil
}
