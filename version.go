package marksman

import (
	"context"
	"net/http"

	"github.com/maxreiter/marksman/snipeit"
)

func (c *Client) FetchVersion(ctx context.Context) (*snipeit.Version, error) {
	var payload *snipeit.Version
	return payload, c.do(ctx, nil, &payload, http.MethodGet, "version")
}
