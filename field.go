package marksman

import (
	"context"
	"net/http"

	"github.com/maxreiter/marksman/snipeit"
)

type Fields struct {
	Total int              `json:"total"`
	Rows  []*snipeit.Field `json:"rows"`
}

func (c *Client) Fields(ctx context.Context) (*Fields, error) {
	req := request{
		method: http.MethodGet,
		url:    "/fields",
	}

	var response *Fields
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
