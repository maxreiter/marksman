package snipeit

import "github.com/maxreiter/marksman/pkg/jsonx/nullable"

type Asset struct {
	ID               int32            `json:"id"`
	Name             string           `json:"name"`
	AssetTag         string           `json:"asset_tag"`
	Serial           nullable.String  `json:"serial"`
	Model            *Model           `json:"model"`
	BYOD             bool             `json:"byod"`
	Requestable      bool             `json:"requestable"`
	ModelNumber      nullable.String  `json:"model_number"`
	EOL              nullable.String  `json:"eol"`
	AssetEOLDate     Date             `json:"asset_eol_date"`
	StatusLabel      *StatusLabel     `json:"status_label"`
	Category         *Category        `json:"category"`
	Manufacturer     *Manufacturer    `json:"manufacturer"`
	Supplier         *Supplier        `json:"supplier"`
	Notes            nullable.String  `json:"notes"`
	OrderNumber      nullable.String  `json:"order_number"`
	Company          *Company         `json:"company"`
	Location         *Location        `json:"location"`
	RTDLocation      *Location        `json:"rtd_location"`
	Image            nullable.String  `json:"image"`
	QR               nullable.String  `json:"qr"`
	AlternateBarcode nullable.String  `json:"alt_barcode"`
	AssignedTo       *User            `json:"assigned_to"`
	WarrantyMonths   nullable.String  `json:"warranty_months"`
	WarrantyExpres   Time             `json:"warranty_expires"`
	CreatedBy        *User            `json:"created_by"`
	CreatedAt        Time             `json:"created_at"`
	UpdatedAt        Time             `json:"updated_at"`
	LastAuditDate    Date             `json:"last_audit_date"`
	NextAuditDate    Date             `json:"next_audit_date"`
	DeletedAt        Time             `json:"deleted_at"`
	PurchaseDate     Date             `json:"purchase_date"`
	Age              nullable.String  `json:"age"`
	LastCheckin      Time             `json:"last_checkin"`
	LastCheckout     Time             `json:"last_checkout"`
	ExpectedCheckin  Time             `json:"expected_checkin"`
	PurchaseCost     nullable.String  `json:"purchase_cost"`
	CheckinCounter   int              `json:"checkin_counter"`
	CheckoutCounter  int              `json:"checkout_counter"`
	RequestsCounter  int              `json:"requests_counter"`
	BooKValue        nullable.String  `json:"book_value"`
	CustomFields     map[string]Field `json:"custom_fields"`
	// AvailableActions
}
