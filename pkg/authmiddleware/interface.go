package authmiddleware

import (
	"context"
	"net/http"
	"net/url"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type PromAPIClient interface {
	URL(ep string, args map[string]string) *url.URL
	Do(context.Context, *http.Request) (*http.Response, []byte, v1.Warnings, error)
	DoGetFallback(ctx context.Context, u *url.URL, args url.Values) (*http.Response, []byte, v1.Warnings, error)
}
