package snipeit

import "github.com/maxreiter/marksman/snipeit/null"

type AssignmentType string

const (
	AssignmentTypeUser     AssignmentType = "user"
	AssignmentTypeAsset    AssignmentType = "asset"
	AssignmentTypeLocation AssignmentType = "location"
)

type Assignment struct {
	ID          UserID              `json:"id"`
	Username    string              `json:"username,omitempty"`
	Name        string              `json:"name,omitempty"`
	FirstName   string              `json:"first_name,omitempty"`
	LastName    null.NullableString `json:"last_name,omitempty"`
	Email       null.NullableString `json:"email,omitempty"`
	EmployeeNum null.NullableString `json:"employee_num,omitempty"`
	Type        AssignmentType      `json:"type,omitempty"`
}

type Asset struct {
	ID               AssetID             `json:"id"`
	Name             string              `json:"name"`
	AssetTag         string              `json:"asset_tag,omitempty"`
	Serial           string              `json:"serial,omitempty"`
	Model            *Model              `json:"model,omitempty"`
	BYOD             bool                `json:"byod,omitempty"`
	Requestable      bool                `json:"requestable,omitempty"`
	ModelNumber      null.NullableString `json:"model_number,omitempty"`
	EOL              null.NullableString `json:"eol,omitempty"`
	AssetEOLDate     Datetime            `json:"asset_eol_date,omitempty"`
	StatusLabel      *StatusLabel        `json:"status_label,omitempty"`
	Category         *Category           `json:"category,omitempty"`
	Manufacturer     *Manufacturer       `json:"manufacturer,omitempty"`
	Supplier         *Supplier           `json:"supplier,omitempty"`
	Notes            null.NullableString `json:"notes,omitempty"`
	OrderNumber      null.NullableString `json:"order_number,omitempty"`
	Company          *Company            `json:"company,omitempty"`
	Location         *Location           `json:"location,omitempty"`
	RTDLocation      *Location           `json:"rtd_location,omitempty"`
	Image            null.NullableString `json:"image,omitempty"`
	QR               null.NullableString `json:"qr,omitempty"`
	AltBarcode       null.NullableString `json:"alt_barcode,omitempty"`
	AssignedTo       *Assignment         `json:"assigned_to,omitempty"`
	WarrantyMonths   null.NullableString `json:"warranty_months,omitempty"`
	WarrantyExpires  Datetime            `json:"warranty_expires,omitempty"`
	CreatedBy        *User               `json:"created_by,omitempty"`
	CreatedAt        Datetime            `json:"created_at,omitempty"`
	UpdatedAt        Datetime            `json:"updated_at,omitempty"`
	LastAuditDate    Datetime            `json:"last_audit_date,omitempty"`
	NextAuditDate    Datetime            `json:"next_audit_date,omitempty"`
	DeletedAt        Datetime            `json:"deleted_at,omitempty"`
	PurchaseDate     Datetime            `json:"purchase_date,omitempty"`
	Age              string              `json:"age,omitempty"`
	LastCheckout     Datetime            `json:"last_checkout,omitempty"`
	LastCheckin      Datetime            `json:"last_checkin,omitempty"`
	ExpectedCheckin  Datetime            `json:"expected_checkin,omitempty"`
	PurchaseCost     string              `json:"purchase_cost,omitempty"`
	CheckinCounter   int                 `json:"checkin_counter,omitempty"`
	CheckoutCounter  int                 `json:"checkout_counter,omitempty"`
	RequestsCounter  int                 `json:"requests_counter,omitempty"`
	UserCanCheckout  bool                `json:"user_can_checkout,omitempty"`
	BookValue        string              `json:"book_value,omitempty"`
	CustomFields     []*Field            `json:"custom_fields,omitempty"`
	AvailableActions *AvailableActions   `json:"available_actions,omitempty"`
}
