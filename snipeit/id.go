package snipeit

import (
	"strconv"
)

// ID is the base ID type of any SnipeIT model.
type ID int32

// NullID represents a null ID.
const NullID = ^ID(0)

// IsNull returns whether or not an ID is null.
func (id ID) IsNull() bool {
	return id == NullID
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (id *ID) UnmarshalJSON(b []byte) error {
	val := string(b)
	if val == "null" {
		*id = NullID
	}

	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return err
	}

	*id = ID(i)
	return nil
}

// MarshalJSON implements the [json.Marshaler] interface.
func (id ID) MarshalJSON() ([]byte, error) {
	if id.IsNull() {
		return []byte("null"), nil
	}

	return []byte(strconv.FormatInt(int64(id), 32)), nil
}

// AccessoryID represents the ID of an [Accessory].
type AccessoryID ID

// ActionLogID represents the ID of an [ActionLog].
type ActionLogID ID

// ConsumableID represents the ID of a [Consumable].
type ConsumableID ID

// UserID represents the ID of a [User].
type UserID ID

// AssetID represents the ID of an [Asset].
type AssetID ID

// CategoryID represents the ID of a [Category].
type CategoryID ID

// CompanyID represents the ID of a [Company].
type CompanyID ID

// Component represents the ID of a [Component].
type ComponentID ID

// DepartmentID represents the ID of a [Department].
type DepartmentID ID

// DepreciationID represents the ID of a [Depreciation].
type DepreciationID ID

// FieldID represents the ID of a [Field].
type FieldID ID

// FieldsetID represents the ID of a [Fieldset].
type FieldsetID ID

// GroupID represents the ID of a [Group].
type GroupID ID

// LicenseID represents the ID of a [License].
type LicenseID ID

// LicenseSeatID represents the ID of a [LicenseSeat].
type LicenseSeatID ID

// LocationID represents the ID of a [Location].
type LocationID ID

// ManufacturerID represents the ID of a [Manufacturer].
type ManufacturerID ID

// ModelID represents the ID of a [Model].
type ModelID ID

// Maintenance represents the ID of a [Maintenance].
type MaintenanceID ID

// StatusLabelID represents the ID of a [StatusLabel].
type StatusLabelID ID

// SupplierID represents the ID of a [Supplier].
type SupplierID ID
