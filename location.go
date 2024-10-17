package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/location"
	"github.com/maxreiter/marksman/snipeit"
)

type Locations struct {
	Total int                 `json:"total"`
	Rows  []*snipeit.Location `json:"rows"`
}

func (c *Client) Locations(ctx context.Context, opts ...location.RequestOption) (*Locations, error) {
	ro := &location.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/locations",
		query:  values,
	}

	var response *Locations
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateLocation(ctx context.Context, id snipeit.LocationID, opts ...location.RequestOption) error {
	ro := &location.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return nil
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/locations/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) Location(ctx context.Context, id snipeit.LocationID) (*snipeit.Location, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/locations/%d", id),
	}

	var response *snipeit.Location
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateLocation(ctx context.Context, id snipeit.LocationID, opts ...location.RequestOption) error {
	ro := &location.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/locations/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteLocation(ctx context.Context, id snipeit.LocationID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/locations/%d", id),
	}

	return c.do(ctx, req, nil)
}
