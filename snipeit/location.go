package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type Location struct {
	ID                  LocationID          `json:"id"`
	Name                string              `json:"name"`
	Image               null.NullableString `json:"image,omitempty"`
	Address             null.NullableString `json:"address,omitempty"`
	Address2            null.NullableString `json:"address2,omitempty"`
	City                null.NullableString `json:"city,omitempty"`
	State               null.NullableString `json:"state,omitempty"`
	Country             null.NullableString `json:"country,omitempty"`
	Zip                 null.NullableString `json:"zip,omitempty"`
	Phone               null.NullableString `json:"phone,omitempty"`
	Fax                 null.NullableString `json:"fax,omitempty"`
	AssignedAssetsCount int                 `json:"assigned_assets_count,omitempty"`
	AssetsCount         int                 `json:"assets_count,omitempty"`
	RTDAssetsCount      int                 `json:"rtd_assets_count,omitempty"`
	UsersCount          int                 `json:"users_count,omitempty"`
	Currency            null.NullableString `json:"currency,omitempty"`
	LDAPOu              null.NullableString `json:"ldap_ou,omitempty"`
	CreatedAt           Datetime            `json:"created_at,omitempty"`
	UpdatedAt           Datetime            `json:"updated_at,omitempty"`
	Parent              *Location           `json:"parent,omitempty"`
	Manager             *User               `json:"manager,omitempty"`
	Children            []*Location         `json:"children,omitempty"`
}
