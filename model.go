package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/model"
	"github.com/maxreiter/marksman/snipeit"
)

// Models is the expected response from endpoints that list [snipeit.Model].
type Models struct {
	Total int              `json:"total"`
	Rows  []*snipeit.Model `json:"rows"`
}

// Models fetches a list of [snipeit.Model].
//
// The following query parameters are accepted:
//   - [model.Limit]: defaults to 50
//   - [model.Offset]: defaults to 0
//   - [model.Search]
//   - [model.Sort]: defaults to "created_at"
//   - [model.Order]: defaults to "asc"
func (c *Client) Models(ctx context.Context, opts ...model.RequestOption) (*Models, error) {
	ro := &model.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    "/models",
		query:  ro,
	}

	var response *Models
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateModel creates a new [snipeit.Model].
//
// The following body parameters are accepted:
//   - [model.Name]: required
//   - [model.CategoryID]: required
//   - [model.ModelNumber]
//   - [model.ManufacturerID]
//   - [model.EOL]
//   - [model.FieldsetID]
func (c *Client) CreateModel(ctx context.Context, opts ...model.RequestOption) error {
	ro := &model.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    "/models",
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) Model(ctx context.Context, id snipeit.ModelID, opts ...model.RequestOption) (*snipeit.Model, error) {
	ro := &model.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/models/%d", id),
		query:  ro,
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

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/models/%d", id),
		body:   ro,
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
