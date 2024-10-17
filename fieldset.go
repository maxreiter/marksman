package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/fieldset"
	"github.com/maxreiter/marksman/snipeit"
)

type Fieldsets struct {
	Total int                 `json:"total"`
	Rows  []*snipeit.Fieldset `json:"rows"`
}

func (c *Client) Fieldsets(ctx context.Context) (*Fieldsets, error) {
	req := request{
		method: http.MethodGet,
		url:    "/fieldsets",
	}

	var response *Fieldsets
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateFieldset(ctx context.Context, opts ...fieldset.RequestOption) error {
	ro := &fieldset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/fieldsets",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) Fieldset(ctx context.Context, id snipeit.FieldsetID) (*snipeit.Fieldset, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/fieldsets/%d", id),
	}

	var response *snipeit.Fieldset
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateFieldset(ctx context.Context, id snipeit.FieldsetID, opts ...fieldset.RequestOption) error {
	ro := &fieldset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/fieldsets/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteFieldset(ctx context.Context, id snipeit.FieldsetID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/fieldsets/%d", id),
	}

	return c.do(ctx, req, nil)
}

func (c *Client) FieldsetFields(ctx context.Context, id snipeit.FieldsetID) (*Fields, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/fieldsets/%d/fields", id),
	}

	var response *Fields
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
