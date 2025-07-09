package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Group struct {
	ID          int32           `json:"id"`
	Name        string          `json:"name"`
	Permissions *Permission     `json:"permissions"`
	UsersCount  int             `json:"users_count"`
	Notes       nullable.String `json:"notes"`
	CreatedBy   *User           `json:"created_by"`
	CreatedAt   Time            `json:"created_at"`
	UpdatedAt   Time            `json:"updated_at"`
}
