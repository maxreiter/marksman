package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type User struct {
	ID                    int32           `json:"id"`
	Avatar                nullable.String `json:"avatar"`
	Name                  nullable.String `json:"name"`
	FirstName             nullable.String `json:"first_name"`
	LastName              nullable.String `json:"last_name"`
	Username              nullable.String `json:"username"`
	Remote                bool            `json:"remote"`
	Locale                string          `json:"locale"`
	EmployeeNumber        nullable.String `json:"employee_num"`
	Manager               *User           `json:"user"`
	JobTitle              nullable.String `json:"jobtitle"`
	VIP                   bool            `json:"vip"`
	Phone                 nullable.String `json:"phone"`
	Website               nullable.String `json:"website"`
	Address               nullable.String `json:"address"`
	City                  nullable.String `json:"city"`
	State                 nullable.String `json:"state"`
	Country               nullable.String `json:"country"`
	Zip                   nullable.String `json:"zip"`
	Email                 nullable.String `json:"email"`
	Department            *Department     `json:"department"`
	DepartmentManager     *User           `json:"department_manager"`
	Location              *Location       `json:"location"`
	Notes                 nullable.String `json:"notes"`
	Permissions           *Permission     `json:"permissions"`
	Activated             bool            `json:"activated"`
	AutoassignLicenses    bool            `json:"autoassign_licenses"`
	LDAPImport            bool            `json:"ldap_import"`
	TwoFactorEnrolled     bool            `json:"two_factor_enrolled"`
	TwoFactorOptin        bool            `json:"two_factor_optin"`
	AssetsCount           int             `json:"assets_count"`
	LicensesCount         int             `json:"licenses_count"`
	AccessoriesCount      int             `json:"accessories_count"`
	ConsumablesCount      int             `json:"consumables_count"`
	ManagesUsersCount     int             `json:"manages_users_count"`
	ManagesLocationsCount int             `json:"manages_locations_count"`
	Company               *Company        `json:"company"`
	CreatedBy             *User           `json:"created_by"`
	CreatedAt             Time            `json:"created_at"`
	UpdatedAt             Time            `json:"updated_at"`
	StartDate             Date            `json:"start_date"`
	EndDate               Date            `json:"end_date"`
	LastLogin             Time            `json:"last_login"`
	// TODO: AvailableActions
	Groups *UserGroups `json:"groups"`
}

type UserGroups struct {
	Total int     `json:"total"`
	Rows  []Group `json:"rows"`
}
