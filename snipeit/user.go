package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type User struct {
	ID                    UserID              `json:"id"`
	Avatar                null.NullableString `json:"avatar,omitempty"`
	Name                  string              `json:"name"`
	FirstName             string              `json:"first_name,omitempty"`
	LastName              string              `json:"last_name,omitempty"`
	Username              string              `json:"username,omitempty"`
	Remote                bool                `json:"remote,omitempty"`
	Locale                null.NullableString `json:"locale,omitempty"`
	EmployeeNum           null.NullableString `json:"employee_num,omitempty"`
	Manager               *User               `json:"manager,omitempty"`
	JobTitle              null.NullableString `json:"jobtitle,omitempty"`
	VIP                   bool                `json:"vip,omitempty"`
	Phone                 null.NullableString `json:"phone,omitempty"`
	Website               null.NullableString `json:"website,omitempty"`
	Address               null.NullableString `json:"address,omitempty"`
	City                  null.NullableString `json:"city,omitempty"`
	State                 null.NullableString `json:"state,omitempty"`
	Country               null.NullableString `json:"country,omitempty"`
	Zip                   null.NullableString `json:"zip,omitempty"`
	Email                 null.NullableString `json:"email,omitempty"`
	Department            *Department         `json:"department,omitempty"`
	Location              *Location           `json:"location,omitempty"`
	Notes                 null.NullableString `json:"notes,omitempty"`
	Permissions           *Permissions        `json:"permissions,omitempty"`
	Activated             bool                `json:"activated,omitempty"`
	AutoAssignLicenses    bool                `json:"autoassign_licenses,omitempty"`
	LDAPImport            bool                `json:"ldap_import,omitempty"`
	TwoFactorEnrolled     bool                `json:"two_factor_enrolled,omitempty"`
	TwoFactorOptin        bool                `json:"two_factor_optin,omitempty"`
	AssetsCount           int                 `json:"assets_count"`
	LicensesCount         int                 `json:"licenses_count,omitempty"`
	AccessoriesCount      int                 `json:"accessories_count,omitempty"`
	ConsumablesCount      int                 `json:"consumables_count,omitempty"`
	ManagesUsersCount     int                 `json:"manages_users_count,omitempty"`
	ManagesLocationsCount int                 `json:"manages_locations_count,omitempty"`
	Company               *Company            `json:"company,omitempty"`
	CreatedBy             *User               `json:"created_by,omitempty"`
	CreatedAt             Datetime            `json:"created_at,omitempty"`
	UpdatedAt             Datetime            `json:"updated_at,omitempty"`
	StartDate             Datetime            `json:"start_date,omitempty"`
	EndDate               Datetime            `json:"end_date,omitempty"`
	LastLogin             Datetime            `json:"last_login,omitempty"`
	DeletedAt             Datetime            `json:"deleted_at,omitempty"`
}
