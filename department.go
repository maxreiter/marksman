package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/department"
	"github.com/maxreiter/marksman/snipeit"
)

type Departments struct {
	Total int                   `json:"total"`
	Rows  []*snipeit.Department `json:"rows"`
}

func (c *Client) Departments(ctx context.Context, opts ...department.RequestOption) (*Departments, error) {
	ro := &department.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/departments",
		query:  values,
	}

	var response *Departments
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateDepartment(ctx context.Context, opts ...department.RequestOption) error {
	ro := &department.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/departments",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) Department(ctx context.Context, id snipeit.DepartmentID) (*snipeit.Department, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/departments/%d", id),
	}

	var response *snipeit.Department
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateDepartment(ctx context.Context, id snipeit.DepartmentID, opts ...department.RequestOption) error {
	ro := &department.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/departments/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteDepartment(ctx context.Context, id snipeit.DepartmentID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/departments/%d", id),
	}

	return c.do(ctx, req, nil)
}
