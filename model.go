package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/model"
	"github.com/maxreiter/marksman/snipeit"
)

type Models struct {
	Total int              `json:"total"`
	Rows  []*snipeit.Model `json:"rows"`
}

func (c *Client) Models(ctx context.Context, opts ...model.RequestOption) (*Models, error) {
	ro := &model.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/models",
		query:  values,
	}

	var response *Models
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateModel(ctx context.Context, opts ...model.RequestOption) error {
	ro := &model.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/models",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) Model(ctx context.Context, id snipeit.ModelID, opts ...model.RequestOption) (*snipeit.Model, error) {
	ro := &model.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/models/%d", id),
		query:  values,
	}

	var response *snipeit.Model
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateModel(ctx context.Context, id snipeit.ModelID, opts ...model.RequestOption) error {
	ro := &model.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/models/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteModel(ctx context.Context, id snipeit.ModelID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/models/%d", id),
	}

	return c.do(ctx, req, nil)
}
