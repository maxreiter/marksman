package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/consumable"
	"github.com/maxreiter/marksman/snipeit"
)

type Consumables struct {
	Total int                   `json:"total"`
	Rows  []*snipeit.Consumable `json:"rows"`
}

func (c *Client) Consumables(ctx context.Context, opts ...consumable.RequestOption) (*Consumables, error) {
	ro := &consumable.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/comsumables",
		query:  values,
	}

	var response *Consumables
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateConsumable(ctx context.Context, opts ...consumable.RequestOption) error {
	ro := &consumable.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return nil
	}

	req := request{
		method: http.MethodPost,
		url:    "/consumables",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

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

func (c *Client) UpdateConsumable(ctx context.Context, id snipeit.ConsumableID, opts ...consumable.RequestOption) error {
	ro := &consumable.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/consumables/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteConsumable(ctx context.Context, id snipeit.ConsumableID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/consumables/%d", id),
	}

	return c.do(ctx, req, nil)
}

func (c *Client) CheckoutConsumable(ctx context.Context, id snipeit.ConsumableID) error {
	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/consumables/%d/checkout", id),
	}

	return c.do(ctx, req, nil)
}
