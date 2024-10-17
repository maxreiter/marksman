package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/group"
	"github.com/maxreiter/marksman/snipeit"
)

type Groups struct {
	Total int              `json:"total"`
	Rows  []*snipeit.Group `json:"rows"`
}

func (c *Client) Groups(ctx context.Context, opts ...group.RequestOption) (*Groups, error) {
	ro := &group.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/groups",
		query:  values,
	}

	var response *Groups
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Group(ctx context.Context, id snipeit.GroupID) (*snipeit.Group, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/groups/%d", id),
	}

	var response *snipeit.Group
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateGroup(ctx context.Context, opts ...group.RequestOption) error {
	ro := &group.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/groups",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) UpdateGroup(ctx context.Context, id snipeit.GroupID, opts ...group.RequestOption) error {
	ro := &group.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/groups/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteGroup(ctx context.Context, id snipeit.GroupID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/groups/%d", id),
	}

	return c.do(ctx, req, nil)
}
