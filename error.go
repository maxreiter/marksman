package marksman

import (
	"fmt"
	"strings"
)

// Error is an error sent by the SnipeIT API.
type Error struct {
	Response
}

// Error implements the error interface.
func (e Error) Error() string {
	if e.Message != "" {
		response := e.Message

		if strings.HasSuffix(response, ".") {
			response = response[:len(response)-1]
		}

		return strings.ToLower(response)
	}

	if len(e.ValidationErrors) > 0 {
		var sb strings.Builder

		for key, value := range e.ValidationErrors {
			if len(value) > 0 {
				sb.WriteString(key + ": ")

				for _, message := range value {
					if message == "" {
						continue
					}

					response := message
					if strings.HasSuffix(response, ".") {
						response = response[:len(response)-1]
					}

					sb.WriteString(strings.ToLower(response) + "; ")
				}
			}
		}

		return sb.String()[:sb.Len()-2]
	}

	return "unhandled error"
}

// HTTPError is an error containing information pertaining to an HTTP request.
type HTTPError struct {
	Code int
}

// Error implements the error interface.
func (e HTTPError) Error() string {
	return fmt.Sprintf("snipe: unexpected status code returned by API: %d", e.Code)
}
