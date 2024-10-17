package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/component"
	"github.com/maxreiter/marksman/snipeit"
)

type Components struct {
	Total int                  `json:"total"`
	Rows  []*snipeit.Component `json:"rows"`
}

func (c *Client) Components(ctx context.Context, opts ...component.RequestOption) (*Components, error) {
	ro := &component.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
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

func (c *Client) CreateComponent(ctx context.Context, opts ...component.RequestOption) error {
	ro := &component.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) UpdateComponent(ctx context.Context, id snipeit.ComponentID, opts ...component.RequestOption) error {
	ro := &component.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
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

func (c *Client) DeleteComponent(ctx context.Context, id snipeit.ComponentID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/components/%d", id),
	}

	return c.do(ctx, req, nil)
}

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

func (c *Client) CheckoutComponent(ctx context.Context, id snipeit.ComponentID, opts ...component.RequestOption) error {
	ro := &component.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) CheckinComponent(ctx context.Context, id snipeit.ComponentID, opts ...component.RequestOption) error {
	ro := &component.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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
