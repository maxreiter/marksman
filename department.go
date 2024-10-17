package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/department"
	"github.com/maxreiter/marksman/snipeit"
)

// Departments is the expected response from endpoints that list [snipeit.Department].
type Departments struct {
	Total int                   `json:"total"`
	Rows  []*snipeit.Department `json:"rows"`
}

// Departments fetches a list of [snipeit.Department].
//
// The following query parameters are accepted:
//   - [department.Name]
//   - [department.CompanyID]
//   - [department.ManagerID]
//   - [department.LocationID]
func (c *Client) Departments(ctx context.Context, opts ...department.RequestOption) (*Departments, error) {
	ro := &department.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    "/departments",
		query:  ro,
	}

	var response *Departments
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateDepartment creates a new [snipeit.Department].
//
// The following body parameters are accepted:
//   - [department.Name]: required
func (c *Client) CreateDepartment(ctx context.Context, opts ...department.RequestOption) error {
	ro := &department.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    "/departments",
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// Department fetches a single [snipeit.Department].
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

// UpdateDepartment updates a [snipeit.Department].
//
// The following body parameters are accepted:
//   - [department.Name]: required
func (c *Client) UpdateDepartment(ctx context.Context, id snipeit.DepartmentID, opts ...department.RequestOption) error {
	ro := &department.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/departments/%d", id),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// DeleteDepartment deletes a [snipeit.Department].
func (c *Client) DeleteDepartment(ctx context.Context, id snipeit.DepartmentID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/departments/%d", id),
	}

	return c.do(ctx, req, nil)
}
