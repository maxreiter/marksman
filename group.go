package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/group"
	"github.com/maxreiter/marksman/snipeit"
)

// Groups is the expected response from endpoints that list [snipeit.Group].
type Groups struct {
	Total int              `json:"total"`
	Rows  []*snipeit.Group `json:"rows"`
}

// Groups fetches a list of [snipeit.Group].
//
// The following query parameters are accepted:
//   - [group.Name]
func (c *Client) Groups(ctx context.Context, opts ...group.RequestOption) (*Groups, error) {
	ro := &group.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    "/groups",
		query:  ro,
	}

	var response *Groups
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// Group fetches a single [snipeit.Group].
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

// CreateGroup creates a new [snipeit.Group].
//
// The following body parameters are accepted:
//   - [group.Name]: required
//   - [group.Permissions]
func (c *Client) CreateGroup(ctx context.Context, opts ...group.RequestOption) error {
	ro := &group.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    "/groups",
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// UpdateGroup updates a [snipeit.Group].
//
// The following body parameters are accepted:
//   - [group.Name]: required
//   - [group.Permissions]
func (c *Client) UpdateGroup(ctx context.Context, id snipeit.GroupID, opts ...group.RequestOption) error {
	ro := &group.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/groups/%d", id),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// DeleteGroup deletes a [snipeit.Group].
func (c *Client) DeleteGroup(ctx context.Context, id snipeit.GroupID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/groups/%d", id),
	}

	return c.do(ctx, req, nil)
}
