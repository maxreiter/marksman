package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

// Department represents an institutional department in the SnipeIT interface.
type Department struct {
	ID               DepartmentID        `json:"id"`
	Name             string              `json:"name"`
	Phone            null.NullableString `json:"phone,omitempty"`
	Fax              null.NullableString `json:"fax,omitempty"`
	Image            null.NullableString `json:"image,omitempty"`
	Company          *Company            `json:"company,omitempty"`
	Manager          *User               `json:"manager,omitempty"`
	Location         *Location           `json:"location,omitempty"`
	UsersCount       int                 `json:"users_count,omitempty"`
	CreatedAt        Datetime            `json:"created_at,omitempty"`
	UpdatedAt        Datetime            `json:"updated_at,omitempty"`
	AvailableActions *AvailableActions   `json:"available_actions,omitempty"`
}
