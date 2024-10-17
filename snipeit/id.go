package snipeit

import "encoding/json"

// ID is the base ID type of any SnipeIT model.
type ID int32

// AccessoryID represents the ID of an [Accessory].
type AccessoryID ID

// ActionLogID represents the ID of an [ActionLog].
type ActionLogID ID

// ConsumableID represents the ID of a [Consumable].
type ConsumableID ID

// UserID represents the ID of a [User].
type UserID ID

// NullUserID is a null [UserID].
const NullUserID = ^UserID(-1)

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (id *UserID) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*id = NullUserID
		return nil
	}

	return json.Unmarshal(b, &id)
}

// MarshalJSON implements the [json.Marshaler] interface.
func (id UserID) MarshalJSON() ([]byte, error) {
	if id == -1 {
		return []byte("null"), nil
	}

	return json.Marshal(id)
}

// AssetID represents the ID of an [Asset].
type AssetID ID

// NullAssetID is a null asset ID.
const NullAssetID = ^AssetID(-1)

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (id *AssetID) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*id = NullAssetID
	}

	return json.Unmarshal(b, &id)
}

// MarshalJSON implements the [json.Marshaler] interface.
func (id AssetID) MarshalJSON() ([]byte, error) {
	if id == -1 {
		return []byte("null"), nil
	}

	return json.Marshal(id)
}

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
