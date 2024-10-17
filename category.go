package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/category"
	"github.com/maxreiter/marksman/snipeit"
)

// Categories is the expected response from endpoints that list categories.
type Categories struct {
	Total int                 `json:"total"`
	Rows  []*snipeit.Category `json:"rows"`
}

// Categories fetches a list of [snipeit.Category].
//
// The following query parameters are accepted:
//
//   - [category.Name]
//   - [category.Limit]: defaults to 50
//   - [category.Offset]: defaults to 0
//   - [category.Search]
//   - [category.Sort]: defaults to "created_at"
//   - [category.Order]: defaults to "desc"
//   - [category.CategoryID]
//   - [category.CategoryType]
//   - [category.UseDefaultEULA]
//   - [category.RequireAcceptance]
//   - [category.CheckinEmail]
func (c *Client) Categories(ctx context.Context, opts ...category.RequestOption) (*Categories, error) {
	ro := &category.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    "/categories",
		query:  ro,
	}

	var response *Categories
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateCategory creates a new [snipeit.Category].
//
// The following body parameters are accepted:
//   - [category.Name]: required
//   - [category.CategoryType]: required
//   - [category.UseDefaultEULA]: defaults to false
//   - [category.RequireAcceptance]: defaults to false
//   - [category.CheckinEmail]: defaults to false
func (c *Client) CreateCategory(ctx context.Context, opts ...category.RequestOption) (*snipeit.Category, error) {
	ro := &category.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    "/categories",
		body:   ro,
	}

	var response *snipeit.Category
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// Category fetches a single [snipeit.Category].
func (c *Client) Category(ctx context.Context, id snipeit.CategoryID) (*snipeit.Category, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/categories/%d", id),
	}

	var response *snipeit.Category
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateCategory updates a [snipeit.Category].
//
// The following body parameters are accepted:
//   - [category.Name]: required
//   - [category.CategoryType]: required
//   - [category.UseDefaultEULA]: defaults to false
//   - [category.RequireAcceptance]: defaults to false
//   - [category.CheckinEmail]: defaults to false
func (c *Client) UpdateCategory(ctx context.Context, id snipeit.CategoryID, opts ...category.RequestOption) (*snipeit.Category, error) {
	ro := &category.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/categories/%d", id),
		body:   ro,
	}

	var response *snipeit.Category
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteCategory deletes a [snipeit.Category].
func (c *Client) DeleteCategory(ctx context.Context, id snipeit.CategoryID) (*snipeit.Category, error) {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/categories/%d", id),
	}

	var response *snipeit.Category
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
