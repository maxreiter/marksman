package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/manufacturer"
	"github.com/maxreiter/marksman/snipeit"
)

// Manufacturers is the expected response from endpoints that list [snipeit.Manufacturer].
type Manufacturers struct {
	Total int                     `json:"total"`
	Rows  []*snipeit.Manufacturer `json:"rows"`
}

// Manufacturers fetches a list of [snipeit.Manufacturer].
//
// The following query parameters are accepted:
//   - [manufacturer.Name]
//   - [manufacturer.URL]
//   - [manufacturer.SupportURL]
//   - [manufacturer.SupportPhone]
//   - [manufacturer.SupportEmail]
func (c *Client) Manufacturers(ctx context.Context, opts ...manufacturer.RequestOption) (*Manufacturers, error) {
	ro := &manufacturer.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
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

// CreateManufacturer creates a new [manufacturer.Manufacturer].
//
// The following body parameters are accepted:
//   - [manufacturer.Name]: required
func (c *Client) CreateManufacturer(ctx context.Context, id snipeit.ManufacturerID, opts ...manufacturer.RequestOption) error {
	ro := &manufacturer.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
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

// Manufacturer fetches a single [snipeit.Manufacturer].
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

// UpdateManufacturer updates a [snipeit.Manufacturer].
//
// The following body parameters are accepted:
//   - [manufacturer.Name]: required
func (c *Client) UpdateManufacturer(ctx context.Context, id snipeit.ManufacturerID, opts ...manufacturer.RequestOption) error {
	ro := &manufacturer.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
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

// DeleteManufacturer deletes a [snipeit.Manufacturer].
func (c *Client) DeleteManufacturer(ctx context.Context, id snipeit.ManufacturerID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/manufacturers/%d", id),
	}

	return c.do(ctx, req, nil)
}
