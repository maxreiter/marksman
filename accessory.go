package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/accessory"
	"github.com/maxreiter/marksman/snipeit"
)

// Accessories is the expected response from endpoints that list accessories.
type Accessories struct {
	Total int                  `json:"total"`
	Rows  []*snipeit.Accessory `json:"rows"`
}

// Accessories fetches a list of [snipeit.Accessory].
//
// The following query parameters are accepted:
//   - [accessory.Limit]: defaults to 50
//   - [accessory.Offset]: defaults to 0
//   - [accessory.Search]
//   - [accessory.OrderNumber]: defaults to null
//   - [accessory.Sort]: defaults to "created_at"
//   - [accessory.Order]: defaults to "desc"
//   - [accessory.Expand]: defaults to false
func (c *Client) Accessories(ctx context.Context, opts ...accessory.RequestOption) (*Accessories, error) {
	ro := &accessory.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/accessories",
		query:  values,
	}

	var response *Accessories
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateAccessory creates a new [snipeit.Accessory].
//
// The following body parameters are accepted:
//   - [accessory.Name]: required
//   - [accessory.Qty]: required
//   - [accessory.CategoryID]: required
//   - [accessory.OrderNumber]
//   - [accessory.PurchaseCost]
//   - [accessory.PurchaseDate]
//   - [accessory.ModelNumber]
//   - [accessory.CompanyID]
//   - [accessory.LocationID]
//   - [accessory.ManufacturerID]
//   - [accessory.SupplierID]
func (c *Client) CreateAccessory(ctx context.Context, opts ...accessory.RequestOption) error {
	ro := &accessory.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/accessories",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

// Accessory fetches a single [snipeit.Accessory].
func (c *Client) Accessory(ctx context.Context, id snipeit.AccessoryID) (*snipeit.Accessory, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/accessories/%d", id),
	}

	var response *snipeit.Accessory
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateAccessory updates an [snipeit.Accessory].
//
// The following body parameters are accepted:
//   - [accessory.Name]: required
//   - [accessory.Qty]: required
//   - [accessory.CategoryID]: required
//   - [accessory.OrderNumber]
//   - [accessory.PurchaseCost]
//   - [accessory.PurchaseDate]
//   - [accessory.ModelNumber]
//   - [accessory.CompanyID]
//   - [accessory.LocationID]
//   - [accessory.ManufacturerID]
//   - [accessory.SupplierID]
func (c *Client) UpdateAccessory(ctx context.Context, id snipeit.AccessoryID, opts ...accessory.RequestOption) error {
	ro := &accessory.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/accessories/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

// DeleteAccessory deletes an [snipe.Accessory].
func (c *Client) DeleteAccessory(ctx context.Context, id snipeit.AccessoryID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/accessories/%d", id),
	}

	return c.do(ctx, req, nil)
}

// AccessoryCheckouts fetches the list of checkouts associated with an [snipeit.Accessory].
//
// The following query parameters are accepted:
//   - [accessory.Limit]
//   - [accessory.Offset]
func (c *Client) AccessoryCheckouts(ctx context.Context, id snipeit.AccessoryID, opts ...accessory.RequestOption) (*Accessories, error) {
	ro := &accessory.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/accessories/%d/checkedout", id),
		query:  values,
	}

	var response *Accessories
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CheckoutAccessory checks an [snipeit.Accessory] out to a [snipeit.User].
//
// The following query parameters are accepted:
//   - [accessory.Limit]
//   - [accessory.Offset]
//
// The following body parameters are accepted:
//   - [accessory.AssignedUser]: required
//   - [accessory.Note]
//   - [accessory.CheckoutQty]: defaults to 1
func (c *Client) CheckoutAccessory(ctx context.Context, id snipeit.AccessoryID, opts ...accessory.RequestOption) error {
	ro := &accessory.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
	if err != nil {
		return err
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/accessories/%d/checkout", id),
		body:   bod,
		query:  values,
	}

	return c.do(ctx, req, nil)
}

// CheckinAccessory checks an [snipeit.Accessory] in from a [snipeit.User].
func (c *Client) CheckinAccessory(ctx context.Context, id int32) error {
	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/accessories/%d/checkin", id),
	}

	return c.do(ctx, req, nil)
}
