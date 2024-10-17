package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/maintenance"
	"github.com/maxreiter/marksman/snipeit"
)

// Maintenances is the expected response from endpoints that list [snipeit.Maintenance].
type Maintenances struct {
	Total int                    `json:"total"`
	Rows  []*snipeit.Maintenance `json:"rows"`
}

// Maintenances fetches a list of [snipeit.Maintenance].
//
// The following query parameters are accepted:
//   - [maintenance.Limit]: defaults to 50
//   - [maintenance.Offset]: defaults to 0
//   - [maintenance.Search]
//   - [maintenance.Sort]: defaults to "created_at"
//   - [maintenance.Order]
//   - [maintenance.AssetID]
func (c *Client) Maintenances(ctx context.Context, opts ...maintenance.RequestOption) (*Maintenances, error) {
	ro := &maintenance.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    "/maintenances",
		query:  ro,
	}

	var response *Maintenances
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateMaintenance creates a new [snipeit.Maintenance].
//
// The following body parameters are accepted:
//   - [maintenance.Title]: required
//   - [maintenance.AssetID]: required
//   - [maintenance.SupplierID]: required
//   - [maintenance.AssetMaintenanceType]: required
//   - [maintenance.StartDate]: required
//   - [maintenance.IsWarranty]
//   - [maintenance.Cost]
//   - [maintenance.Notes]
//   - [maintenance.CompletionDate]
func (c *Client) CreateMaintenance(ctx context.Context, opts ...maintenance.RequestOption) error {
	ro := &maintenance.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    "/maintenances",
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// UpdateMaintenance updates a [snipeit.Maintenance].
//
// The following body parameters are accepted:
//   - [maintenance.Title]: required
//   - [maintenance.AssetID]: required
//   - [maintenance.SupplierID]: required
//   - [maintenance.AssetMaintenanceType]: required
//   - [maintenance.StartDate]: required
//   - [maintenance.IsWarranty]
//   - [maintenance.Cost]
//   - [maintenance.Notes]
//   - [maintenance.CompletionDate]
func (c *Client) UpdateMaintenance(ctx context.Context, id snipeit.MaintenanceID, opts ...maintenance.RequestOption) error {
	ro := &maintenance.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/maintenances/%d", id),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// DeleteMaintenance deletes a [snipeit.Maintenance].
func (c *Client) DeleteMaintenance(ctx context.Context, id snipeit.MaintenanceID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/maintenances/%d", id),
	}

	return c.do(ctx, req, nil)
}
