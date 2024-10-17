package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/manufacturer"
	"github.com/maxreiter/marksman/snipeit"
)

type Manufacturers struct {
	Total int                     `json:"total"`
	Rows  []*snipeit.Manufacturer `json:"rows"`
}

func (c *Client) Manufacturers(ctx context.Context, opts ...manufacturer.RequestOption) (*Manufacturers, error) {
	ro := &manufacturer.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/manufacturers",
		query:  values,
	}

	var response *Manufacturers
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateManufacturer(ctx context.Context, id snipeit.ManufacturerID, opts ...manufacturer.RequestOption) error {
	ro := &manufacturer.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/manufacturers/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) Manufacturer(ctx context.Context, id snipeit.ManufacturerID) (*snipeit.Manufacturer, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/manufacturer/%d", id),
	}

	var response *snipeit.Manufacturer
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateManufacturer(ctx context.Context, id snipeit.ManufacturerID, opts ...manufacturer.RequestOption) error {
	ro := &manufacturer.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/manufacturers/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteManufacturer(ctx context.Context, id snipeit.ManufacturerID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/manufacturers/%d", id),
	}

	return c.do(ctx, req, nil)
}
