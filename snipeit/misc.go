package snipeit

type AvailableActions struct {
	Update         bool `json:"update,omitempty"`
	Delete         bool `json:"delete,omitempty"`
	Clone          bool `json:"clone,omitempty"`
	Restore        bool `json:"restore,omitempty"`
	Checkout       bool `json:"checkout,omitempty"`
	Checkin        bool `json:"checkin,omitempty"`
	BulkSelectable struct {
		Delete bool `json:"delete,omitempty"`
	} `json:"bulk_selectable,omitempty"`
}
