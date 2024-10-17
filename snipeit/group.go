package snipeit

type Group struct {
	ID               GroupID           `json:"id"`
	Name             string            `json:"name"`
	Permissions      *Permissions      `json:"permissions"`
	UsersCount       int               `json:"users_count"`
	CreatedBy        *User             `json:"created_by"`
	CreatedAt        Datetime          `json:"created_at"`
	UpdatedAt        Datetime          `json:"updated_at"`
	AvailableActions *AvailableActions `json:"available_actions"`
}
