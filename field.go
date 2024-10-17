package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/field"
	"github.com/maxreiter/marksman/snipeit"
)

type Fields struct {
	Total int              `json:"total"`
	Rows  []*snipeit.Field `json:"rows"`
}

func (c *Client) Fields(ctx context.Context) (*Fields, error) {
	req := request{
		method: http.MethodGet,
		url:    "/fields",
	}

	var response *Fields
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateField(ctx context.Context, opts ...field.RequestOption) error {
	ro := &field.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/fields",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) Field(ctx context.Context, id snipeit.FieldID) (*snipeit.Field, error) {
	req := request{
		method: http.MethodGet,
		url:    "/fields",
	}

	var response *snipeit.Field
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateField(ctx context.Context, id snipeit.FieldID, opts ...field.RequestOption) error {
	ro := &field.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/fields/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteField(ctx context.Context, id snipeit.FieldID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/fields/%d", id),
	}

	return c.do(ctx, req, nil)
}

func (c *Client) AssociateField(ctx context.Context, id snipeit.FieldID, opts ...field.RequestOption) error {
	ro := &field.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/fields/%d/associate", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DisassociateField(ctx context.Context, id snipeit.FieldID, opts ...field.RequestOption) error {
	ro := &field.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/fields/%d/disassociate", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}
