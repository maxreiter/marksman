package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/accessory"
	"github.com/maxreiter/marksman/snipeit"
)

type Accessories struct {
	Total int                  `json:"total"`
	Rows  []*snipeit.Accessory `json:"rows"`
}

func (c *Client) Accessories(ctx context.Context, opts ...accessory.RequestOption) (*Accessories, error) {
	ro := &accessory.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
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

func (c *Client) CreateAccessory(ctx context.Context, opts ...accessory.RequestOption) error {
	ro := &accessory.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) UpdateAccessory(ctx context.Context, id snipeit.AccessoryID, opts ...accessory.RequestOption) error {
	ro := &accessory.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) DeleteAccessory(ctx context.Context, id snipeit.AccessoryID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/accessories/%d", id),
	}

	return c.do(ctx, req, nil)
}

func (c *Client) AccessoryCheckouts(ctx context.Context, id snipeit.AccessoryID, opts ...accessory.RequestOption) (*Accessories, error) {
	ro := &accessory.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
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

func (c *Client) CheckoutAccessory(ctx context.Context, id snipeit.AccessoryID, opts ...accessory.RequestOption) error {
	ro := &accessory.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return err
	}

	bod, err := ro.Marshal()
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

func (c *Client) CheckinAccessory(ctx context.Context, id int32) error {
	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/accessories/%d/checkin", id),
	}

	return c.do(ctx, req, nil)
}
