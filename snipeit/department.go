package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Department struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	Phone      nullable.String `json:"phone"`
	Fax        nullable.String `json:"fax"`
	Image      nullable.String `json:"image"`
	Company    *Company        `json:"company"`
	Manager    *User           `json:"user"`
	Location   *Location       `json:"location"`
	UsersCount int             `json:"users_count,string"`
	Notes      nullable.String `json:"notes"`
	CreatedAt  Time            `json:"created_at"`
	UpdatedAt  Time            `json:"updated_at"`
	// TODO: AvailableActions
}
