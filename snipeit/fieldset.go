package snipeit

type Fieldset struct {
	ID        int32          `json:"id"`
	Name      string         `json:"name"`
	Fields    FieldsetFields `json:"fields"`
	Models    FieldsetModels `json:"models"`
	CreatedAt Time           `json:"created_at"`
	UpdatedAt Time           `json:"updated_at"`
}

type FieldsetFields struct {
	Total int     `json:"total"`
	Rows  []Field `json:"rows"`
}

type FieldsetModels struct {
	Total int     `json:"total"`
	Rows  []Model `json:"rows"`
}
