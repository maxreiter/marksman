package httpx

import (
	"net/http"
	"time"
)

const (
	ContentTypeJSON = "application/json"
)

var DefaultClient = &http.Client{
	Timeout: time.Second * 15,
}

func IsOK(status int) bool {
	return status >= http.StatusOK && status < http.StatusMultipleChoices
}
