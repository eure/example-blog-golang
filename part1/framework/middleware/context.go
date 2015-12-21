package middleware

import (
	"net/http"

	"github.com/zenazn/goji/web"
	"golang.org/x/net/context"
)

// ContextKey is key name for stored context
const ContextKey = "context"

// Context creates new Context for new request
func Context(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "request", r)
		c.Env[ContextKey] = ctx
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
