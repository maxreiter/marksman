package request

import "github.com/maxreiter/marksman/snipeit"

type FetchCategories struct {
	Name              string               `url:"name,omitempty"`
	Limit             int32                `url:"limit,omitempty"`
	Offset            int32                `url:"offset,omitempty"`
	Search            string               `url:"search,omitempty"`
	Sort              string               `url:"sort,omitempty"`
	Order             SortOrderType        `url:"order,omitempty"`
	CategoryID        int32                `url:"category_id,omitempty"`
	CategoryType      snipeit.CategoryType `url:"category_type,omitempty"`
	UseDefaultEULA    bool                 `url:"use_default_eula,omitempty"`
	RequireAcceptance bool                 `url:"require_acceptance,omitempty"`
	CheckinEmail      bool                 `url:"checkin_email,omitempty"`
}

type CreateCategory struct {
	Name              string               `json:"name"`
	CategoryType      snipeit.CategoryType `json:"category_type"`
	UseDefaultEULA    bool                 `json:"use_default_eula,omitempty"`
	RequireAcceptance bool                 `json:"require_acceptance,omitempty"`
	CheckinEmail      bool                 `json:"checkin_email,omitempty"`
}

type UpdateCategory struct {
	Name              string               `json:"name"`
	CategoryType      snipeit.CategoryType `json:"category_type"`
	UseDefaultEULA    bool                 `json:"use_default_eula,omitempty"`
	RequireAcceptance bool                 `json:"require_acceptance,omitempty"`
	CheckinEmail      bool                 `json:"checkin_email,omitempty"`
}

type PartiallyUpdateCategory struct {
	Name              string               `json:"name,omitempty"`
	CategoryType      snipeit.CategoryType `json:"category_type,omitempty"`
	UseDefaultEULA    bool                 `json:"use_default_eula,omitempty"`
	RequireAcceptance bool                 `json:"require_acceptance,omitempty"`
	CheckinEmail      bool                 `json:"checkin_email,omitempty"`
}
