package marksman

import (
	"context"
	"net/http"

	"github.com/maxreiter/marksman/params/supplier"
	"github.com/maxreiter/marksman/snipeit"
)

type Suppliers struct {
	Total int                 `json:"total"`
	Rows  []*snipeit.Supplier `json:"rows"`
}

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
