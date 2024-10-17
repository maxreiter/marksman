package marksman

import (
	"context"
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
