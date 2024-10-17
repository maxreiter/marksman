package params

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"

	"github.com/google/go-querystring/query"
)

type SortType string

const (
	SortID              SortType = "id"
	SortName            SortType = "name"
	SortAssetTag        SortType = "asset_tag"
	SortSerial          SortType = "serial"
	SortModel           SortType = "model"
	SortModelNumber     SortType = "model_number"
	SortLastCheckout    SortType = "last_checkout"
	SortCategory        SortType = "category"
	SortManufacturer    SortType = "manufacturer"
	SortNotes           SortType = "notes"
	SortExpectedCheckin SortType = "expected_checkin"
	SortOrderNumber     SortType = "order_number"
	SortCompanyName     SortType = "companyName"
	SortLocation        SortType = "location"
	SortImage           SortType = "image"
	SortStatusLabel     SortType = "status_label"
	SortAssignedTo      SortType = "assigned_to"
	SortCreatedAt       SortType = "created_at"
	SortPurchaseDate    SortType = "purchase_date"
	SortPurchaseCost    SortType = "purchase_cost"
)

type OrderType string

const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)

type Resolver interface {
	Values() (url.Values, error)
	Marshal() (io.Reader, error)
}

type BaseResolver struct{}

func (r *BaseResolver) Values() (url.Values, error) {
	return query.Values(r)
}

func (r *BaseResolver) Marshal() (io.Reader, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(r); err != nil {
		return nil, err
	}

	return &buf, nil
}
