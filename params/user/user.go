// Package user provides request configuration for methods of the marksman client.
package user

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

// UserGroups wraps a slice of [snipeit.GroupID] for proper JSON handling.
type UserGroups []snipeit.GroupID

// MarshalJSON implements the [json.Marshaler] interface.
func (g UserGroups) MarshalJSON() ([]byte, error) {
	if len(g) == 0 {
		return []byte(`null`), nil
	}

	if len(g) == 1 {
		return json.Marshal(g[0])
	}

	return json.Marshal(g)
}

// RequestOptions contains possible options for requests made to the /users endpoints.
type RequestOptions struct {
	*params.Resolver

	// Query params
	Search           string             `url:"search,omitempty" json:"-"`
	Limit            int32              `url:"limit,omitempty" json:"-"`
	Offset           int32              `url:"offset,omitempty" json:"-" `
	Sort             params.SortType    `url:"sort,omitempty" json:"-" `
	Order            params.OrderType   `url:"order,omitempty" json:"-"`
	State            string             `url:"state,omitempty" json:"-"`
	Zip              string             `url:"zip,omitempty" json:"-"`
	Country          string             `url:"country,omitempty" json:"-"`
	GroupID          snipeit.GroupID    `url:"group_id,omitempty" json:"-"`
	Deleted          bool               `url:"deleted,omitempty" json:"-"`
	All              bool               `url:"all,omitempty" json:"-"`
	LDAPImport       bool               `url:"ldap_import,omitempty" json:"-"`
	AssetsCount      int32              `url:"assets_count,omitempty" json:"-"`
	LicensesCount    int32              `url:"licenses_count,omitempty" json:"-"`
	AccessoriesCount int32              `url:"accessories_count,omitempty" json:"-"`
	ConsumablesCount int32              `url:"consumables_count,omitempty" json:"-"`
	CategoryID       snipeit.CategoryID `url:"category_id,omitempty" json:"-"`
	ModelID          snipeit.ModelID    `url:"model_id,omitempty" json:"-"`

	// Compound params
	FirstName    string               `url:"first_name,omitempty" json:"first_name,omitempty"`
	LastName     string               `url:"last_name,omitempty" json:"last_name,omitempty"`
	Username     string               `url:"username,omitempty" json:"username,omitempty"`
	Email        string               `url:"email,omitempty" json:"email,omitempty"`
	EmployeeNum  string               `url:"employee_num,omitempty" json:"employee_num,omitempty"`
	CompanyID    snipeit.CompanyID    `url:"company_id,omitempty" json:"company_id,omitempty"`
	DepartmentID snipeit.DepartmentID `url:"department_id,omitempty" json:"department_id,omitempty"`
	LocationID   snipeit.LocationID   `url:"location_id,omitempty" json:"location_id,omitempty"`
	Remote       bool                 `url:"remote,omitempty" json:"remote,omitempty"`
	StartDate    string               `url:"start_date,omitempty" json:"start_date,omitempty"`
	EndDate      string               `url:"end_date,omitempty" json:"end_date,omitempty"`

	// Body params
	Password             string         `json:"password,omitempty" url:"-"`
	PasswordConfirmation string         `json:"password_confirmation,omitempty" url:"-"`
	Permissions          string         `json:"permissions,omitempty" url:"-"`
	Activated            bool           `json:"activated,omitempty" url:"-"`
	Phone                string         `json:"phone,omitempty" url:"-"`
	JobTitle             string         `json:"jobtitle,omitempty" url:"-"`
	ManagerID            snipeit.UserID `json:"manager_id,omitempty" url:"-"`
	Notes                string         `json:"notes,omitempty" url:"-"`
	TwoFactorEnrolled    bool           `json:"two_factor_enrolled,omitempty" url:"-"`
	TwoFactorOptin       bool           `json:"two_factor_optin,omitempty" url:"-"`
	Groups               UserGroups     `json:"groups,omitempty" url:"-"`
	VIP                  bool           `json:"vip,omitempty" url:"-"`
}

// Query marshals the [RequestOptions] as a [url.Values].
func (ro *RequestOptions) Query() (url.Values, error) {
	return ro.Resolver.Query()
}

// JSON encodes the [RequestOptions] as JSON.
func (ro *RequestOptions) JSON() (io.Reader, error) {
	return ro.Resolver.JSON()
}

// RequestOption is used to configure a [RequestOptions].
type RequestOption func(*RequestOptions)

// Search sets the search string of a request.
func Search(search string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Search = search
	}
}

// Limit sets the return limit of a request.
func Limit(limit int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Limit = limit
	}
}

// Offset sets the pagination offset of a request.
func Offset(offset int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Offset = offset
	}
}

// Sort sets the return sort of a request.
func Sort(sort params.SortType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Sort = sort
	}
}

// Order sets the return order of a request.
func Order(order params.OrderType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Order = order
	}
}

// State sets the state of a [snipeit.User].
func State(state string) RequestOption {
	return func(ro *RequestOptions) {
		ro.State = state
	}
}

// Zip sets the zip code of a [snipeit.User].
func Zip(zip string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Zip = zip
	}
}

// Country sets the country of a [snipeit.User].
func Country(country string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Country = country
	}
}

// GroupID sets the group ID of a [snipeit.User].
func GroupID(id snipeit.GroupID) RequestOption {
	return func(ro *RequestOptions) {
		ro.GroupID = id
	}
}

// Deleted toggles the inclusion of deleted [snipeit.User] in a request.
func Deleted() RequestOption {
	return func(ro *RequestOptions) {
		ro.Deleted = true
	}
}

// All toggles the inclusion of both active and deleted [snipeit.User] in a request.
func All() RequestOption {
	return func(ro *RequestOptions) {
		ro.All = true
	}
}

// LDAPImport marks a [snipeit.User] as having been imported via LDAP.
func LDAPImport() RequestOption {
	return func(ro *RequestOptions) {
		ro.LDAPImport = true
	}
}

// AssetsCount sets the number of [snipeit.Asset] checked out to a [snipeit.User].
func AssetsCount(count int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssetsCount = count
	}
}

// LicensesCount sets the number of [snipeit.License] checked out to a [snipeit.User].
func LicensesCount(count int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.LicensesCount = count
	}
}

// AccessoriesCount sets the number of [snipeit.Accessory] checked out to a [snipeit.User].
func AccessoriesCount(count int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.AccessoriesCount = count
	}
}

// ConsumablesCount sets the number of [snipeit.Consumable] checked out to a [snipeit.User].
func ConsumablesCount(count int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.ConsumablesCount = count
	}
}

// CategoryID sets the [snipeit.CategoryID] of a [snipeit.User].
func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

// ModelID sets the [snipeit.ModelID] of a [snipeit.User].
func ModelID(id snipeit.ModelID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ModelID = id
	}
}

// FirstName sets the first name of a [snipeit.User].
func FirstName(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.FirstName = name
	}
}

// LastName sets the last name of a [snipeit.User].
func LastName(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.LastName = name
	}
}

// Username sets the username of a [snipeit.User].
func Username(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Username = name
	}
}

// Email sets the email address of a [snipeit.User].
func Email(email string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Email = email
	}
}

// EmployeeNum sets the employee number of a [snipeit.User].
func EmployeeNum(num string) RequestOption {
	return func(ro *RequestOptions) {
		ro.EmployeeNum = num
	}
}

// CompanyID sets the [snipeit.CompanyID] of a [snipeit.User].
func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

// DepartmentID sets the [snipeit.DepartmentID] of a [snipeit.User].
func DepartmentID(id snipeit.DepartmentID) RequestOption {
	return func(ro *RequestOptions) {
		ro.DepartmentID = id
	}
}

// LocationID sets the [snipeit.LocationID] of a [snipeit.User].
func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

// Remote marks a [snipeit.User] as remote.
func Remote() RequestOption {
	return func(ro *RequestOptions) {
		ro.Remote = true
	}
}

// StartDate sets the start date of a [snipeit.User].
func StartDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.StartDate = date
	}
}

// EndDate sets the end date of a [snipeit.User].
func EndDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.EndDate = date
	}
}

// Password sets the password of a [snipeit.User].
func Password(password string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Password = password
		ro.PasswordConfirmation = password
	}
}

// Permissions sets the permissions of a [snipeit.User].
func Permissions(perms string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Permissions = perms
	}
}

// Activated marks a [snipeit.User] as activated.
func Activated() RequestOption {
	return func(ro *RequestOptions) {
		ro.Activated = true
	}
}

// Phone sets the phone number of a [snipeit.User].
func Phone(phone string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Phone = phone
	}
}

// JobTitle sets the job title of a [snipeit.User].
func JobTitle(title string) RequestOption {
	return func(ro *RequestOptions) {
		ro.JobTitle = title
	}
}

// ManagerID sets the manager ID of a [snipeit.User].
func ManagerID(id snipeit.UserID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManagerID = id
	}
}

// Notes sets the notes of a [snipeit.User].
func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

// TwoFactorEnrolled marks a [snipeit.User] as enrolled in two factor authentication.
func TwoFactorEnrolled() RequestOption {
	return func(ro *RequestOptions) {
		ro.TwoFactorEnrolled = true
	}
}

// TwoFactorOptin denotes that a [snipeit.User] may opt in to two factor authentication.
func TwoFactorOptin() RequestOption {
	return func(ro *RequestOptions) {
		ro.TwoFactorOptin = true
	}
}

// Groups sets the groups a [snipeit.User] is part of.
func Groups(groups ...snipeit.GroupID) RequestOption {
	return func(ro *RequestOptions) {
		if len(ro.Groups) > 0 {
			ro.Groups = append(ro.Groups, groups...)
		} else {
			ro.Groups = groups
		}
	}
}

// VIP marks a [snipeit.User] as a VIP.
func VIP() RequestOption {
	return func(ro *RequestOptions) {
		ro.VIP = true
	}
}
