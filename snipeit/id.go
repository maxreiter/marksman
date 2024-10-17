package snipeit

import "encoding/json"

type ID int32

type AccessoryID ID

type ConsumerID ID

type ConsumableID ID

type UserID ID

const NullUserID = ^UserID(-1)

func (id *UserID) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*id = NullUserID
		return nil
	}

	return json.Unmarshal(b, &id)
}

func (id UserID) MarshalJSON() ([]byte, error) {
	if id == -1 {
		return []byte("null"), nil
	}

	return json.Marshal(id)
}

type AssetID ID

const NullAssetID = ^AssetID(-1)

func (id *AssetID) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*id = NullAssetID
	}

	return json.Unmarshal(b, &id)
}

func (id AssetID) MarshalJSON() ([]byte, error) {
	if id == -1 {
		return []byte("null"), nil
	}

	return json.Marshal(id)
}

type CategoryID ID

type CompanyID ID

type ComponentID ID

type DepartmentID ID

type DepreciationID ID

type FieldID ID

type FieldsetID ID

type GroupID ID

type LicenseID ID

type LicenseSeatID ID

type LocationID ID

type ManufacturerID ID

type ModelID ID

type MaintenanceID ID

type StatusLabelID ID

type SupplierID ID
