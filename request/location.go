package request

type FetchLocations struct {
	Name     string        `url:"name,omitempty"`
	Limit    int32         `url:"limit,omitempty"`
	Offset   int32         `url:"offset,omitempty"`
	Search   string        `url:"search,omitempty"`
	Sort     string        `url:"sort,omitempty"`
	Order    SortOrderType `url:"order,omitempty"`
	Address  string        `url:"address,omitempty"`
	Address2 string        `url:"address2,omitempty"`
	City     string        `url:"city,omitempty"`
	Zip      string        `url:"zip,omitempty"`
	Country  string        `url:"country,omitempty"`
}

type CreateLocation struct {
	Name      string `json:"name"`
	Address   string `json:"address,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Country   string `json:"country,omitempty"`
	Zip       string `json:"zip,omitempty"`
	LDAPOU    string `json:"ldap_ou,omitempty"`
	ParentID  int32  `json:"parent_id,omitempty"`
	Currency  string `json:"currency,omitempty"`
	ManagerID int32  `json:"manager_id,omitempty"`
}

type UpdateLocation struct {
	Name      string `json:"name"`
	Address   string `json:"address,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Country   string `json:"country,omitempty"`
	Zip       string `json:"zip,omitempty"`
	LDAPOU    string `json:"ldap_ou,omitempty"`
	ParentID  int32  `json:"parent_id,omitempty"`
	Currency  string `json:"currency,omitempty"`
	ManagerID int32  `json:"manager_id,omitempty"`
}

type PartiallyUpdateLocation struct {
	Name      string `json:"name,omitempty"`
	Address   string `json:"address,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Country   string `json:"country,omitempty"`
	Zip       string `json:"zip,omitempty"`
	LDAPOU    string `json:"ldap_ou,omitempty"`
	ParentID  int32  `json:"parent_id,omitempty"`
	Currency  string `json:"currency,omitempty"`
	ManagerID int32  `json:"manager_id,omitempty"`
}
