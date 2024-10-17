package marksman

import (
	"context"
	"net/http"

	"github.com/maxreiter/marksman/params/report"
	"github.com/maxreiter/marksman/snipeit"
)

// ReportActivities is the expected response from endpoints that list [snipeit.ActionLog].
type ReportActivities struct {
	Total int                  `json:"total"`
	Rows  []*snipeit.ActionLog `json:"rows"`
}

// ReportActivities fetches a list of [snipeit.ActionLog].
//
// The following query parameters are accepted:
//   - [report.Limit]: defaults to 2
//   - [report.Offset]: defaults to 0
//   - [report.Search]
//   - [report.TargetType]: required if passing target_id
//   - [report.TargetID]: required if passing target_type
//   - [report.ItemType]: required if passing item_id
//   - [report.ItemID]: required if passing item_type
//   - [report.ActionType]
//   - [report.Order]: defaults to "desc"
//   - [report.Sort]: defaults to "created_at"
func (c *Client) ReportActivities(ctx context.Context, opts ...report.RequestOption) (*ReportActivities, error) {
	ro := &report.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
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
