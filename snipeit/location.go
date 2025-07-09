package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Location struct {
	ID                       int32           `json:"id"`
	Name                     string          `json:"name"`
	Image                    nullable.String `json:"image"`
	Address                  nullable.String `json:"address"`
	Address2                 nullable.String `json:"address2"`
	City                     nullable.String `json:"city"`
	State                    nullable.String `json:"state"`
	Country                  nullable.String `json:"country"`
	Zip                      nullable.String `json:"zip"`
	Phone                    nullable.String `json:"phone"`
	Fax                      nullable.String `json:"fax"`
	AccessoriesCount         int             `json:"accessories_count"`
	AssignedAccessoriesCount int             `json:"assigned_accessories_count"`
	AssignedAssetsCount      int             `json:"assigned_assets_count"`
	AssetsCount              int             `json:"assets_count"`
	RTDAssetsCount           int             `json:"rtd_assets_count"`
	UsersCount               int             `json:"users_count"`
	Currency                 nullable.String `json:"currency"`
	LDAPOU                   nullable.String `json:"ldap_ou"`
	Notes                    nullable.String `json:"notes"`
	CreatedAt                Time            `json:"created_at"`
	UpdatedAt                Time            `json:"updated_at"`
	Parent                   *Location       `json:"parent"`
	Manager                  *User           `json:"manager"`
	Company                  *Company        `json:"company"`
	Children                 []Location      `json:"children"`
}
