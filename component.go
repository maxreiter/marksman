package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/component"
	"github.com/maxreiter/marksman/snipeit"
)

// Components is the expected response from endpoints that list [snipeit.Component].
type Components struct {
	Total int                  `json:"total"`
	Rows  []*snipeit.Component `json:"rows"`
}

// Components fetches a list of [snipeit.Component].
//
// The following query parameters are accepted:
//   - [component.Name]
//   - [component.Limit]: defaults to 50
//   - [component.Offset]: defaults to 0
//   - [component.Search]
//   - [component.OrderNumber]: defaults to null
//   - [component.Sort]: defaults to "created_at"
//   - [component.Order]: defaults to "desc"
//   - [component.Expand]: defaults to false
func (c *Client) Components(ctx context.Context, opts ...component.RequestOption) (*Components, error) {
	ro := &component.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/components",
		query:  values,
	}

	var response *Components
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateComponent creates a new [snipeit.Component].
//
// The following body parameters are accepted:
//   - [component.Name]: required
//   - [component.Qty]: required
//   - [component.CategoryID]: required
//   - [component.LocationID]
//   - [component.CompanyID]
//   - [component.OrderNumber]
//   - [component.PurchaseDate]
//   - [component.PurchaseCost]
//   - [component.MinAmt]
//   - [component.Serial]
func (c *Client) CreateComponent(ctx context.Context, opts ...component.RequestOption) error {
	ro := &component.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/components",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

// Component fetches a single [snipeit.Component].
func (c *Client) Component(ctx context.Context, id snipeit.ComponentID) (*snipeit.Component, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/components/%d", id),
	}

	var response *snipeit.Component
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateComponent updates a [snipeit.Component].
//
// The following body parameters are accepted:
//   - [component.Name]: required
//   - [component.Qty]: required
//   - [component.CategoryID]: required
//   - [component.LocationID]
//   - [component.CompanyID]
//   - [component.OrderNumber]
//   - [component.PurchaseDate]
//   - [component.PurchaseOrder]
//   - [component.MinAmt]
//   - [component.Serial]
func (c *Client) UpdateComponent(ctx context.Context, id snipeit.ComponentID, opts ...component.RequestOption) error {
	ro := &component.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/components/%d", id),
		query:  values,
	}

	return c.do(ctx, req, nil)
}

// DeleteComponent deletes a [snipeit.Component].
func (c *Client) DeleteComponent(ctx context.Context, id snipeit.ComponentID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/components/%d", id),
	}

	return c.do(ctx, req, nil)
}

// ComponentAssets fetches a list of [snipeit.Asset] a [snipeit.Component] has been checked out to.
func (c *Client) ComponentAssets(ctx context.Context, id snipeit.ComponentID) (*Assets, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/components/%d/assets", id),
	}

	var response *Assets
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CheckoutComponent checks a [snipeit.Component] out to an [snipeit.Asset].
//
// The following body parameters are accepted:
//   - [component.AssignedTo]: required
//   - [component.AssignedQty]: required
func (c *Client) CheckoutComponent(ctx context.Context, id snipeit.ComponentID, opts ...component.RequestOption) error {
	ro := &component.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/components/%d/checkout", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

// CheckinComponent checks a [snipeit.Component] in from an [snipeit.Asset].
//
// The following body parameters are accepted:
//   - [component.CheckinQty]: required
func (c *Client) CheckinComponent(ctx context.Context, id snipeit.ComponentID, opts ...component.RequestOption) error {
	ro := &component.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/components/%d/checkin", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}
