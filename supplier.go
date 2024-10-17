package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/supplier"
	"github.com/maxreiter/marksman/snipeit"
)

// Suppliers is the expected response from endpoints that list [snipeit.Supplier].
type Suppliers struct {
	Total int                 `json:"total"`
	Rows  []*snipeit.Supplier `json:"rows"`
}

// Suppliers fetches a list of [snipeit.Supplier].
//
// The following query parameters are accepted:
//   - [supplier.Name]
//   - [supplier.Address]
//   - [supplier.Address2]
//   - [supplier.City]
//   - [supplier.Zip]
//   - [supplier.Country]
//   - [supplier.Fax]
//   - [supplier.Email]
//   - [supplier.URL]
//   - [supplier.Notes]
func (c *Client) Suppliers(ctx context.Context, opts ...supplier.RequestOption) (*Suppliers, error) {
	ro := &supplier.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    "/suppliers",
	}

	var response *Suppliers
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// Supplier fetches a single [snipeit.Supplier].
func (c *Client) Supplier(ctx context.Context, id snipeit.SupplierID) (*snipeit.Supplier, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/suppliers/%d", id),
	}

	var response *snipeit.Supplier
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
