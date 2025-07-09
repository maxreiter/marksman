package request

type FetchDepartments struct {
	Name       string `url:"name,omitempty"`
	CompanyID  int32  `url:"company_id,omitempty"`
	ManagerID  int32  `url:"manager_id,omitempty"`
	LocationID int32  `url:"location_id,omitempty"`
}

type CreateDepartment struct {
	Name       string `json:"name"`
	CompanyID  int32  `json:"company_id"`
	ManagerID  int32  `json:"manager_id"`
	LocationID int32  `json:"location_id"`
}

type UpdateDepartment struct {
	Name       string `json:"name"`
	CompanyID  int32  `json:"company_id"`
	ManagerID  int32  `json:"manager_id"`
	LocationID int32  `json:"location_id"`
}

type PartiallyUpdateDepartment struct {
	Name       string `json:"name,omitempty"`
	CompanyID  int32  `json:"company_id,omitempty"`
	ManagerID  int32  `json:"manager_id,omitempty"`
	LocationID int32  `json:"location_id,omitempty"`
}
