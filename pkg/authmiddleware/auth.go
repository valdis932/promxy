package authmiddleware

import (
	"context"
	"net/http"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
type AuthHeaderKey string

const (
	AuthHeaderCtxKey AuthHeaderKey = "AuthHeader"
)

// TODO: Wrap request
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hdrs := make(http.Header)

		// Copy all headers we want into the context
		hdrs.Add("Authorization", r.Header.Get("Authorization"))

		ctx := context.WithValue(r.Context(), AuthHeaderCtxKey, hdrs)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Wrap client
type AuthClient struct {
	PromAPIClient
}

func (a *AuthClient) Do(ctx context.Context, r *http.Request) (*http.Response, []byte, v1.Warnings, error) {
	hdrsToAdd, _ := ctx.Value(AuthHeaderCtxKey).(http.Header)
	for k, _ := range hdrsToAdd {
		r.Header.Add(k, hdrsToAdd.Get(k))
	}

	return a.PromAPIClient.Do(ctx, r)
}
