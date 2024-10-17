package marksman

import (
	"context"
	"net/http"

	"github.com/maxreiter/marksman/params/report"
	"github.com/maxreiter/marksman/snipeit"
)

type ReportActivities struct {
	Total int                  `json:"total"`
	Rows  []*snipeit.ActionLog `json:"rows"`
}

func (c *Client) ReportActivities(ctx context.Context, opts ...report.RequestOption) (*ReportActivities, error) {
	ro := &report.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/reports/activity",
		query:  values,
	}

	var response *ReportActivities
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
