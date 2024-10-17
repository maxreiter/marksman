package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/company"
	"github.com/maxreiter/marksman/snipeit"
)

type Companies struct {
	Total int                `json:"total"`
	Rows  []*snipeit.Company `json:"rows"`
}

func (c *Client) Companies(ctx context.Context, opts ...company.RequestOption) (*Companies, error) {
	ro := &company.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/companies",
		query:  values,
	}

	var response *Companies
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateCompany(ctx context.Context, opts ...company.RequestOption) error {
	ro := &company.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/companies",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

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

func (c Client) UpdateCompany(ctx context.Context, id snipeit.CompanyID, opts ...company.RequestOption) (*snipeit.Company, error) {
	ro := &company.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/companies/%d", id),
		body:   bod,
	}

	var response *snipeit.Company
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteCompany(ctx context.Context, id snipeit.CompanyID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/companies/%d", id),
	}

	return c.do(ctx, req, nil)
}
