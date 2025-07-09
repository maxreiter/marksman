package request

import "github.com/maxreiter/marksman/snipeit"

type FetchGroups struct {
	Name string `url:"name,omitempty"`
}

type CreateGroup struct {
	Name        string              `json:"name"`
	Permissions *snipeit.Permission `json:"permissions,omitempty"`
}

type UpdateGroup struct {
	Name        string              `json:"name"`
	Permissions *snipeit.Permission `json:"permissions,omitempty"`
}

type PartiallyUpdateGroup struct {
	Name        string              `json:"name,omitempty"`
	Permissions *snipeit.Permission `json:"permissions,omitempty"`
}
