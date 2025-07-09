package request

type FetchModels struct {
	Limit          int32         `url:"limit,omitempty"`
	Offset         int32         `url:"offset,omitempty"`
	Search         string        `url:"search,omitempty"`
	Sort           string        `url:"sort,omitempty"`
	Order          SortOrderType `url:"order,omitempty"`
	Name           string        `url:"name,omitempty"`
	Notes          string        `url:"notes,omitempty"`
	ModelNumber    string        `url:"model_number,omitempty"`
	Requestable    bool          `url:"requestable,omitempty"`
	CategoryID     int32         `url:"category_id,omitempty"`
	DepreciationID int32         `url:"depreciation_id,omitempty"`
}

type CreateModel struct {
	Name           string `json:"name"`
	ModelNumber    string `json:"model_number,omitempty"`
	CategoryID     int32  `json:"category_id"`
	ManufacturerID int32  `json:"manufacturer_id,omitempty"`
	EOL            int32  `json:"eol,omitempty"`
	DepreciationID int32  `json:"depreciation_id,omitempty"`
	FieldsetID     int32  `json:"fieldset_id,omitempty"`
	Requestable    bool   `json:"requestable,omitempty"`
	Notes          string `json:"notes,omitempty"`
}

type UpdateModel struct {
	Name           string `json:"name"`
	ModelNumber    string `json:"model_number,omitempty"`
	CategoryID     int32  `json:"category_id"`
	ManufacturerID int32  `json:"manufacturer_id,omitempty"`
	EOL            int32  `json:"eol,omitempty"`
	DepreciationID int32  `json:"depreciation_id,omitempty"`
	FieldsetID     int32  `json:"fieldset_id,omitempty"`
	Requestable    bool   `json:"requestable,omitempty"`
	Notes          string `json:"notes,omitempty"`
}

type PartiallyUpdateModel struct {
	Name           string `json:"name,omitempty"`
	ModelNumber    string `json:"model_number,omitempty"`
	CategoryID     int32  `json:"category_id,omitempty"`
	ManufacturerID int32  `json:"manufacturer_id,omitempty"`
	EOL            int32  `json:"eol,omitempty"`
	DepreciationID int32  `json:"depreciation_id,omitempty"`
	FieldsetID     int32  `json:"fieldset_id,omitempty"`
	Requestable    bool   `json:"requestable,omitempty"`
	Notes          string `json:"notes,omitempty"`
}
