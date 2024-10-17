package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/maintenance"
	"github.com/maxreiter/marksman/snipeit"
)

type Maintenances struct {
	Total int                    `json:"total"`
	Rows  []*snipeit.Maintenance `json:"rows"`
}

func (c *Client) Maintenances(ctx context.Context, opts ...maintenance.RequestOption) (*Maintenances, error) {
	ro := &maintenance.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/maintenances",
		query:  values,
	}

	var response *Maintenances
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateMaintenance(ctx context.Context, opts ...maintenance.RequestOption) error {
	ro := &maintenance.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/maintenances",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) UpdateMaintenance(ctx context.Context, id snipeit.MaintenanceID, opts ...maintenance.RequestOption) error {
	ro := &maintenance.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/maintenances/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteMaintenance(ctx context.Context, id snipeit.MaintenanceID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/maintenances/%d", id),
	}

	return c.do(ctx, req, nil)
}
