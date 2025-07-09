package request

import "github.com/maxreiter/marksman/snipeit"

type FetchStatusLabels struct {
	Name       string        `url:"name,omitempty"`
	Limit      int32         `url:"limit,omitempty"`
	Offset     int32         `url:"offset,omitempty"`
	Search     string        `url:"search,omitempty"`
	Sort       string        `url:"sort,omitempty"`
	Order      SortOrderType `url:"order,omitempty"`
	StatusType string        `url:"status_type,omitempty"`
}

type CreateStatusLabel struct {
	Name         string                  `json:"name"`
	Type         snipeit.StatusLabelType `json:"type"`
	Notes        string                  `json:"notes,omitempty"`
	Color        string                  `json:"color,omitempty"`
	ShowInNav    bool                    `json:"show_in_nav,omitempty"`
	DefaultLabel bool                    `json:"default_label,omitempty"`
}

type UpdateStatusLabel struct {
	Name         string                  `json:"name"`
	Type         snipeit.StatusLabelType `json:"type"`
	Notes        string                  `json:"notes,omitempty"`
	Color        string                  `json:"color,omitempty"`
	ShowInNav    bool                    `json:"show_in_nav,omitempty"`
	DefaultLabel bool                    `json:"default_label,omitempty"`
}

type PartiallyUpdateStatusLabel struct {
	Name         string                  `json:"name,omitempty"`
	Type         snipeit.StatusLabelType `json:"type,omitempty"`
	Notes        string                  `json:"notes,omitempty"`
	Color        string                  `json:"color,omitempty"`
	ShowInNav    bool                    `json:"show_in_nav,omitempty"`
	DefaultLabel bool                    `json:"default_label,omitempty"`
}
