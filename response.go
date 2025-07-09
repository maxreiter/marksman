package marksman

import "encoding/json"

type payloadResponse struct {
	Status   string           `json:"status"`
	Messages responseMessages `json:"messages"`
	Payload  json.RawMessage  `json:"payload"`
}

type responseMessages struct {
	Message          string
	ValidationErrors map[string][]string
}

func (m *responseMessages) UnmarshalJSON(v []byte) error {
	var value any
	if err := json.Unmarshal(v, &value); err != nil {
		return err
	}

	switch asserted := value.(type) {
	case string:
		m.Message = asserted
	case map[string]any:
		validationErrors := make(map[string][]string)

		for key, unknownValues := range asserted {
			stringValues, ok := unknownValues.([]string)
			if !ok {
				continue
			}

			validationErrors[key] = make([]string, 0, len(stringValues))
			for _, stringValue := range stringValues {
				validationErrors[key] = append(validationErrors[key], stringValue)
			}
		}

		m.ValidationErrors = validationErrors
	}

	return nil
}

type rowResponse struct {
	Total int             `json:"total"`
	Rows  json.RawMessage `json:"rows"`
}
