package marksman

import "github.com/maxreiter/marksman/snipeit"

type Licenses struct {
	Total int                `json:"total"`
	Rows  []*snipeit.License `json:"rows"`
}
