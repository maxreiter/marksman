package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/consumable"
	"github.com/maxreiter/marksman/snipeit"
)

// Consumables is the expected response from endpoints that list [snipeit.Consumable].
type Consumables struct {
	Total int                   `json:"total"`
	Rows  []*snipeit.Consumable `json:"rows"`
}

// Consumables fetches a list of [snipeit.Consumable].
//
// The following query parameters are accepted:
//   - [consumable.Name]
//   - [consumable.Limit]: defaults to 50
//   - [consumable.Offset]: defaults to 0
//   - [consumable.Search]
//   - [consumable.OrderNumber]: defaults to null
//   - [consumable.Sort]: defaults to "created_at"
//   - [consumable.Order]: defaults to "desc"
//   - [consumable.Expand]: defaults to false
//   - [consumable.CategoryID]
//   - [consumable.CompanyID]
//   - [consumable.ManufacturerID]
func (c *Client) Consumables(ctx context.Context, opts ...consumable.RequestOption) (*Consumables, error) {
	ro := &consumable.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    "/comsumables",
		query:  ro,
	}

	var response *Consumables
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateConsumable creates a new [snipeit.Consumable].
//
// The following body parameters are accepted:
//   - [consumable.Name]: required
//   - [consumable.Qty]: required
//   - [consumable.CategoryID]: required
//   - [consumable.OrderNumber]
//   - [consumable.ManufacturerID]
//   - [consumable.LocationID]
//   - [consumable.Requestable]
//   - [consumable.PurchaseDate]
//   - [consumable.MinAmt]
//   - [consumable.ModelNumber]
//   - [consumable.ItemNo]
func (c *Client) CreateConsumable(ctx context.Context, opts ...consumable.RequestOption) error {
	ro := &consumable.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    "/consumables",
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// Consumable fetches a single [snipeit.Consumable].
func (c *Client) Consumable(ctx context.Context, id snipeit.ConsumableID) (*snipeit.Consumable, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/consumables/%d", id),
	}

	var response *snipeit.Consumable
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateConsumable updates a [snipeit.Consumable].
//
// The following body parameters are accepted:
//   - [consumable.Name]: required
//   - [consumable.Qty]: required
//   - [consumable.CategoryID]: required
//   - [consumable.CompanyID]
//   - [consumable.OrderNumber]
//   - [consumable.ManufacturerID]
//   - [consumable.LocationID]
//   - [consumable.Requestable]
//   - [consumable.PurchaseDate]
//   - [consumable.MinAmt]
//   - [consumable.ModelNumber]
//   - [consumable.ItemNo]
func (c *Client) UpdateConsumable(ctx context.Context, id snipeit.ConsumableID, opts ...consumable.RequestOption) error {
	ro := &consumable.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/consumables/%d", id),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// DeleteConsumable deletes a [snipeit.Consumable].
func (c *Client) DeleteConsumable(ctx context.Context, id snipeit.ConsumableID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/consumables/%d", id),
	}

	return c.do(ctx, req, nil)
}

// CheckoutConsumable checks a [snipeit.Consumable] out to a [snipeit.User].
//
// The following body parameters are accepted:
//   - [consumable.AssignedTo]
func (c *Client) CheckoutConsumable(ctx context.Context, id snipeit.ConsumableID) error {
	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/consumables/%d/checkout", id),
	}

	return c.do(ctx, req, nil)
}
