package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/field"
	"github.com/maxreiter/marksman/snipeit"
)

// Fields is the expected response from endpoints listing [snipeit.Field].
type Fields struct {
	Total int              `json:"total"`
	Rows  []*snipeit.Field `json:"rows"`
}

// Fields fetches a list of [snipeit.Field].
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

// CreateField creates a new [snipeit.Field].
//
// The following body parameters are accepted:
//   - [field.Name]: required
//   - [field.Element]: required
//   - [field.FieldValues]
//   - [field.ShowInEmail]: defaults to false
//   - [field.Format]
//   - [field.FieldEncrypted]: defaults to false
//   - [field.HelpText]
func (c *Client) CreateField(ctx context.Context, opts ...field.RequestOption) error {
	ro := &field.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    "/fields",
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// Field fetches a single [snipeit.Field].
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

// UpdateField updates a [snipeit.Field].
//
// The following body parameters are accepted:
//   - [field.Name]: required
//   - [field.Element]: required
func (c *Client) UpdateField(ctx context.Context, id snipeit.FieldID, opts ...field.RequestOption) error {
	ro := &field.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/fields/%d", id),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// DeleteField deletes a [snipeit.Field].
func (c *Client) DeleteField(ctx context.Context, id snipeit.FieldID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/fields/%d", id),
	}

	return c.do(ctx, req, nil)
}

// AssociateField associates a [snipeit.Field] with a [snipeit.Fieldset].
//
// The following body parameters are accepted:
//   - [field.FieldsetID]: required
func (c *Client) AssociateField(ctx context.Context, id snipeit.FieldID, opts ...field.RequestOption) error {
	ro := &field.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/fields/%d/associate", id),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// DisassociateField disassociates a [snipeit.Field] from a [snipeit.Fieldset].
//
// The following body parameters are accepted:
//   - [field.FieldsetID]: required
func (c *Client) DisassociateField(ctx context.Context, id snipeit.FieldID, opts ...field.RequestOption) error {
	ro := &field.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/fields/%d/disassociate", id),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}
