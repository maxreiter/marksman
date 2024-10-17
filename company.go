package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/company"
	"github.com/maxreiter/marksman/snipeit"
)

// Companies is the expected response from endpoints that list companies.
type Companies struct {
	Total int                `json:"total"`
	Rows  []*snipeit.Company `json:"rows"`
}

// Companies fetches a list of [snipeit.Companies].
//
// The following query parameters are accepted:
//   - [company.Name]
func (c *Client) Companies(ctx context.Context, opts ...company.RequestOption) (*Companies, error) {
	ro := &company.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    "/companies",
		query:  ro,
	}

	var response *Companies
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateCompany creates a new [snipeit.Company].
//
// The following body parameters are accepted:
//   - [company.Name]: defaults to "Google, inc."
func (c *Client) CreateCompany(ctx context.Context, opts ...company.RequestOption) error {
	ro := &company.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    "/companies",
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// Company fetches a single [snipeit.Company].
func (c *Client) Company(ctx context.Context, id snipeit.CompanyID) (*snipeit.Company, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/companies/%d", id),
	}

	var response *snipeit.Company
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateCompany updates a [snipeit.Company].
func (c Client) UpdateCompany(ctx context.Context, id snipeit.CompanyID, opts ...company.RequestOption) (*snipeit.Company, error) {
	ro := &company.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/companies/%d", id),
		body:   ro,
	}

	var response *snipeit.Company
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteCompany deletes a [snipeit.Company].
func (c *Client) DeleteCompany(ctx context.Context, id snipeit.CompanyID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/companies/%d", id),
	}

	return c.do(ctx, req, nil)
}
