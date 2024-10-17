package user

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/maxreiter/marksman/params"
	"github.com/maxreiter/marksman/snipeit"
)

type UserGroups []snipeit.GroupID

func (g UserGroups) MarshalJSON() ([]byte, error) {
	if len(g) == 0 {
		return []byte(`null`), nil
	}

	if len(g) == 1 {
		return json.Marshal(g[0])
	}

	return json.Marshal(g)
}

type RequestOptions struct {
	*params.BaseResolver

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
	Password             string     `json:"password,omitempty" url:"-"`
	PasswordConfirmation string     `json:"password_confirmation,omitempty" url:"-"`
	Permissions          string     `json:"permissions,omitempty" url:"-"`
	Activated            bool       `json:"activated,omitempty" url:"-"`
	Phone                string     `json:"phone,omitempty" url:"-"`
	JobTitle             string     `json:"jobtitle,omitempty" url:"-"`
	ManagerID            int32      `json:"manager_id,omitempty" url:"-"`
	Notes                string     `json:"notes,omitempty" url:"-"`
	TwoFactorEnrolled    bool       `json:"two_factor_enrolled,omitempty" url:"-"`
	TwoFactorOptin       bool       `json:"two_factor_optin,omitempty" url:"-"`
	Groups               UserGroups `json:"groups,omitempty" url:"-"`
	VIP                  bool       `json:"vip,omitempty" url:"-"`
}

func (ro *RequestOptions) Values() (url.Values, error) {
	return ro.BaseResolver.Values()
}

func (ro *RequestOptions) Marshal() (io.Reader, error) {
	return ro.BaseResolver.Marshal()
}

type RequestOption func(*RequestOptions)

func Search(search string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Search = search
	}
}

func Limit(limit int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Limit = limit
	}
}

func Offset(offset int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.Offset = offset
	}
}

func Sort(sort params.SortType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Sort = sort
	}
}

func Order(order params.OrderType) RequestOption {
	return func(ro *RequestOptions) {
		ro.Order = order
	}
}

func State(state string) RequestOption {
	return func(ro *RequestOptions) {
		ro.State = state
	}
}

func Zip(zip string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Zip = zip
	}
}

func Country(country string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Country = country
	}
}

func GroupID(id snipeit.GroupID) RequestOption {
	return func(ro *RequestOptions) {
		ro.GroupID = id
	}
}

func Deleted() RequestOption {
	return func(ro *RequestOptions) {
		ro.Deleted = true
	}
}

func All() RequestOption {
	return func(ro *RequestOptions) {
		ro.All = true
	}
}

func LDAPImport() RequestOption {
	return func(ro *RequestOptions) {
		ro.LDAPImport = true
	}
}

func AssetsCount(count int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.AssetsCount = count
	}
}

func LicensesCount(count int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.LicensesCount = count
	}
}

func AccessoriesCount(count int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.AccessoriesCount = count
	}
}

func ConsumablesCount(count int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.ConsumablesCount = count
	}
}

func CategoryID(id snipeit.CategoryID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CategoryID = id
	}
}

func ModelID(id snipeit.ModelID) RequestOption {
	return func(ro *RequestOptions) {
		ro.ModelID = id
	}
}

func FirstName(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.FirstName = name
	}
}

func LastName(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.LastName = name
	}
}

func Username(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Username = name
	}
}

func Email(email string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Email = email
	}
}

func EmployeeNum(num string) RequestOption {
	return func(ro *RequestOptions) {
		ro.EmployeeNum = num
	}
}

func CompanyID(id snipeit.CompanyID) RequestOption {
	return func(ro *RequestOptions) {
		ro.CompanyID = id
	}
}

func DepartmentID(id snipeit.DepartmentID) RequestOption {
	return func(ro *RequestOptions) {
		ro.DepartmentID = id
	}
}

func LocationID(id snipeit.LocationID) RequestOption {
	return func(ro *RequestOptions) {
		ro.LocationID = id
	}
}

func Remote() RequestOption {
	return func(ro *RequestOptions) {
		ro.Remote = true
	}
}

func StartDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.StartDate = date
	}
}

func EndDate(date string) RequestOption {
	return func(ro *RequestOptions) {
		ro.EndDate = date
	}
}

func Password(password string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Password = password
		ro.PasswordConfirmation = password
	}
}

func Permissions(perms string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Permissions = perms
	}
}

func Activated() RequestOption {
	return func(ro *RequestOptions) {
		ro.Activated = true
	}
}

func Phone(phone string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Phone = phone
	}
}

func JobTitle(title string) RequestOption {
	return func(ro *RequestOptions) {
		ro.JobTitle = title
	}
}

func ManagerID(id int32) RequestOption {
	return func(ro *RequestOptions) {
		ro.ManagerID = id
	}
}

func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

func TwoFactorEnrolled() RequestOption {
	return func(ro *RequestOptions) {
		ro.TwoFactorEnrolled = true
	}
}

func TwoFactorOptin() RequestOption {
	return func(ro *RequestOptions) {
		ro.TwoFactorOptin = true
	}
}

func Groups(groups ...snipeit.GroupID) RequestOption {
	return func(ro *RequestOptions) {
		if len(ro.Groups) > 0 {
			ro.Groups = append(ro.Groups, groups...)
		} else {
			ro.Groups = groups
		}
	}
}

func VIP() RequestOption {
	return func(ro *RequestOptions) {
		ro.VIP = true
	}
}
