package marksman

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/maxreiter/marksman/pkg/httpx"
)

type Error struct {
	Message          string
	ValidationErrors map[string][]string
	HTTPStatus       int
}

func createError(response *http.Response) error {
	var payload *payloadResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return err
	}

	err := Error{
		HTTPStatus: response.StatusCode,
	}

	if payload.Messages.Message != "" {
		err.Message = payload.Messages.Message
	}

	if payload.Messages.ValidationErrors != nil {
		err.ValidationErrors = payload.Messages.ValidationErrors
	}

	return err
}

func (e Error) Error() string {
	var builder strings.Builder
	builder.WriteString("marksman: ")

	if !httpx.IsOK(e.HTTPStatus) {
		builder.WriteString("HTTP ")
		builder.WriteString(strconv.Itoa(e.HTTPStatus))
		builder.WriteString(": ")
	}

	if e.Message != "" {
		message := strings.ToLower(e.Message)
		message = strings.TrimSuffix(message, ".")
		builder.WriteString(message)

		return builder.String()
	}

	if e.ValidationErrors != nil {
		for key, values := range e.ValidationErrors {
			key = strings.ToLower(key)

			builder.WriteString(key)
			builder.WriteString(": ")

			for index, message := range values {
				message = strings.ToLower(message)
				message = strings.TrimSuffix(message, ".")
				builder.WriteString(message)

				if index < len(values)-1 {
					builder.WriteString(";")
				}
			}

			builder.WriteString(", ")
		}

		return strings.TrimSuffix(builder.String(), ", ")
	}

	builder.WriteString("unknown error")
	return builder.String()
}
