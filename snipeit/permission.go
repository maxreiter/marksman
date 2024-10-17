package snipeit

type PermissionBool bool

func (b *PermissionBool) UnmarshalJSON(v []byte) error {
	if string(v) == "1" {
		*b = true
		return nil
	}

	*b = false
	return nil
}

type Permissions struct {
	Superuser PermissionBool `json:"superuser,omitempty"`
	Admin     PermissionBool `json:"admin,omitempty"`
	Import    PermissionBool `json:"import,omitempty"`

	ViewReports PermissionBool `json:"reports.view,omitempty"`

	ViewAssets                     PermissionBool `json:"assets.view,omitempty"`
	CreateAssets                   PermissionBool `json:"assets.create,omitempty"`
	EditAssets                     PermissionBool `json:"assets.edit,omitempty"`
	DeleteAssets                   PermissionBool `json:"assets.delete,omitempty"`
	CheckoutAssets                 PermissionBool `json:"assets.checkout,omitempty"`
	CheckinAssets                  PermissionBool `json:"assets.checkin,omitempty"`
	AuditAssets                    PermissionBool `json:"assets.audit,omitempty"`
	ViewRequestableAssets          PermissionBool `json:"assets.view.requestable,omitempty"`
	ViewAssetEncryptedCustomFields PermissionBool `json:"assets.view.encrypted_custom_fields,omitempty"`

	ViewAccessories      PermissionBool `json:"accessories.view,omitempty"`
	CreateAccessories    PermissionBool `json:"accessories.create,omitempty"`
	EditAccessories      PermissionBool `json:"accessories.edit,omitempty"`
	DeleteAccessories    PermissionBool `json:"accessories.delete,omitempty"`
	CheckoutAccessories  PermissionBool `json:"accessories.checkout,omitempty"`
	CheckinAccessories   PermissionBool `json:"accessories.checkin,omitempty"`
	ViewAccessoriesFiles PermissionBool `json:"accessories.files,omitempty"`

	ViewConsumables      PermissionBool `json:"consumables.view,omitempty"`
	CreateConsumables    PermissionBool `json:"consumables.create,omitempty"`
	EditConsumables      PermissionBool `json:"consumables.edit,omitempty"`
	DeleteConsumables    PermissionBool `json:"consumables.delete,omitempty"`
	CheckoutConsumables  PermissionBool `json:"consumables.checkout,omitempty"`
	ViewConsumablesFiles PermissionBool `json:"consumables.files,omitempty"`

	ViewLicenses      PermissionBool `json:"licenses.view,omitempty"`
	CreateLicenses    PermissionBool `json:"licenses.create,omitempty"`
	EditLicenses      PermissionBool `json:"licenses.edit,omitempty"`
	DeleteLicenses    PermissionBool `json:"licenses.delete,omitempty"`
	CheckoutLicenses  PermissionBool `json:"licenses.checkout,omitempty"`
	ViewLicensesKeys  PermissionBool `json:"licenses.keys,omitempty"`
	ViewLicensesFiles PermissionBool `json:"licenses.files,omitempty"`

	ViewComponents      PermissionBool `json:"components.view,omitempty"`
	CreateComponents    PermissionBool `json:"components.create,omitempty"`
	EditComponents      PermissionBool `json:"components.edit,omitempty"`
	DeleteComponents    PermissionBool `json:"components.delete,omitempty"`
	CheckoutComponents  PermissionBool `json:"components.checkout,omitempty"`
	CheckinComponents   PermissionBool `json:"components.checkin,omitempty"`
	ViewComponentsFiles PermissionBool `json:"components.files,omitempty"`

	ViewKits   PermissionBool `json:"kits.view,omitempty"`
	CreateKits PermissionBool `json:"kits.create,omitempty"`
	EditKits   PermissionBool `json:"kits.edit,omitempty"`
	DeleteKits PermissionBool `json:"kits.delete,omitempty"`

	ViewUsers   PermissionBool `json:"users.view,omitempty"`
	CreateUsers PermissionBool `json:"users.create,omitempty"`
	EditUsers   PermissionBool `json:"users.edit,omitempty"`
	DeleteUsers PermissionBool `json:"users.delete,omitempty"`

	ViewModels   PermissionBool `json:"models.view,omitempty"`
	CreateModels PermissionBool `json:"models.create,omitempty"`
	EditModels   PermissionBool `json:"models.edit,omitempty"`
	DeleteModels PermissionBool `json:"models.delete,omitempty"`

	ViewCategories   PermissionBool `json:"categories.view,omitempty"`
	CreateCategories PermissionBool `json:"categories.create,omitempty"`
	EditCategories   PermissionBool `json:"categories.edit,omitempty"`
	DeleteCategories PermissionBool `json:"categories.delete,omitempty"`

	ViewDepartments   PermissionBool `json:"departments.view,omitempty"`
	CreateDepartments PermissionBool `json:"departments.create,omitempty"`
	EditDepartments   PermissionBool `json:"departments.edit,omitempty"`
	DeleteDepartments PermissionBool `json:"departments.delete,omitempty"`

	ViewStatusLabels   PermissionBool `json:"statuslabels.view,omitempty"`
	CreateStatusLabels PermissionBool `json:"statuslabels.create,omitempty"`
	EditStatusLabels   PermissionBool `json:"statuslabels.edit,omitempty"`
	DeleteStatusLabels PermissionBool `json:"statuslabels.delete,omitempty"`

	ViewCustomFields   PermissionBool `json:"customfields.view,omitempty"`
	CreateCustomFields PermissionBool `json:"customfields.create,omitempty"`
	EditCustomFields   PermissionBool `json:"customfields.edit,omitempty"`
	DeleteCustomFields PermissionBool `json:"customfields.delete,omitempty"`

	ViewSuppliers   PermissionBool `json:"suppliers.view,omitempty"`
	CreateSuppliers PermissionBool `json:"suppliers.create,omitempty"`
	EditSuppliers   PermissionBool `json:"suppliers.edit,omitempty"`
	DeleteSuppliers PermissionBool `json:"suppliers.delete,omitempty"`

	ViewManufacturers   PermissionBool `json:"manufacturers.view,omitempty"`
	CreateManufacturers PermissionBool `json:"manufacturers.create,omitempty"`
	EditManufacturers   PermissionBool `json:"manufacturers.edit,omitempty"`
	DeleteManufacturers PermissionBool `json:"manufacturers.delete,omitempty"`

	ViewDepreciations   PermissionBool `json:"depriciations.view,omitempty"`
	CreateDepreciations PermissionBool `json:"depriciations.create,omitempty"`
	EditDepreciations   PermissionBool `json:"depriciations.edit,omitempty"`
	DeleteDepreciations PermissionBool `json:"depriciations.delete,omitempty"`

	ViewLocations   PermissionBool `json:"locations.view,omitempty"`
	CreateLocations PermissionBool `json:"locations.create,omitempty"`
	EditLocations   PermissionBool `json:"locations.edit,omitempty"`
	DeleteLocations PermissionBool `json:"locations.delete,omitempty"`

	ViewCompanies   PermissionBool `json:"companies.view,omitempty"`
	CreateCompanies PermissionBool `json:"companies.create,omitempty"`
	EditCompanies   PermissionBool `json:"companies.edit,omitempty"`
	DeleteCompanies PermissionBool `json:"companies.delete,omitempty"`

	ToggleTwoFactor   PermissionBool `json:"self.two_factor,omitempty"`
	CreateAPIKeys     PermissionBool `json:"self.api,omitempty"`
	EditLocation      PermissionBool `json:"self.edit_location,omitempty"`
	CanCheckoutAssets PermissionBool `json:"self.checkout_assets,omitempty"`
	ViewPurchaseCost  PermissionBool `json:"self.view_purchase_cost,omitempty"`
}
