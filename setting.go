package marksman

import (
	"context"
	"net/http"

	"github.com/maxreiter/marksman/snipeit"
)

type Backups struct {
	Total int               `json:"total"`
	Rows  []*snipeit.Backup `json:"rows"`
}

func (c *Client) Backups(ctx context.Context) (*Backups, error) {
	req := request{
		method: http.MethodGet,
		url:    "/settings/backups",
	}

	var response *Backups
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
