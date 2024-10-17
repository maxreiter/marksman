package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/category"
	"github.com/maxreiter/marksman/snipeit"
)

type Categories struct {
	Total int                 `json:"total"`
	Rows  []*snipeit.Category `json:"rows"`
}

func (c *Client) Categories(ctx context.Context, opts ...category.RequestOption) (*Categories, error) {
	ro := &category.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/categories",
		query:  values,
	}

	var response *Categories
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateCategory(ctx context.Context, opts ...category.RequestOption) (*snipeit.Category, error) {
	ro := &category.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodPost,
		url:    "/categories",
		body:   bod,
	}

	var response *snipeit.Category
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

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

func (c *Client) UpdateCategory(ctx context.Context, id snipeit.CategoryID, opts ...category.RequestOption) (*snipeit.Category, error) {
	ro := &category.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/categories/%d", id),
		body:   bod,
	}

	var response *snipeit.Category
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

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
