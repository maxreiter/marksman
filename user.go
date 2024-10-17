package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/user"
	"github.com/maxreiter/marksman/snipeit"
)

type Users struct {
	Total int             `json:"total"`
	Rows  []*snipeit.User `json:"rows"`
}

func (c *Client) Me(ctx context.Context) (*snipeit.User, error) {
	req := request{
		method: http.MethodGet,
		url:    "/users/me",
	}

	var response *snipeit.User
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Users(ctx context.Context, opts ...user.RequestOption) (*Users, error) {
	ro := &user.RequestOptions{
		Limit:  50,
		Offset: 0,
		Sort:   "created_at",
		Order:  "desc",
	}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		query:  values,
		method: http.MethodGet,
		url:    "/users",
	}

	var response *Users
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) User(ctx context.Context, id int32) (*snipeit.User, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/users/%d", id),
	}

	var response *snipeit.User
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateUser(ctx context.Context, id int32, opts ...user.RequestOption) error {
	req := request{
		method: http.MethodPatch,
		url:    fmt.Sprintf("/users/%d", id),
	}

	reqOpts := &user.RequestOptions{}

	for _, o := range opts {
		o(reqOpts)
	}

	bod, err := reqOpts.Marshal()
	if err != nil {
		return err
	}

	req.body = bod

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteUser(ctx context.Context, id int32) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/users/%d", id),
	}

	return c.do(ctx, req, nil)
}

func (c *Client) RestoreUser(ctx context.Context, id int32) error {
	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/users/%d/restore", id),
	}

	return c.do(ctx, req, nil)
}

func (c *Client) UserAssets(ctx context.Context, id int32, opts ...user.RequestOption) (*Assets, error) {
	ro := &user.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/users/%d/assets", id),
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req.query = values

	var response *Assets
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UserAccessories(ctx context.Context, id int32) (*Accessories, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/users/%d/accessories", id),
	}

	var response *Accessories
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
