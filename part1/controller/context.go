package controller

import (
	"github.com/zenazn/goji/web"
	"golang.org/x/net/context"

	"github.com/eure/example-blog-golang/framework/middleware"
)

// GetContext returns context
func GetContext(c web.C) context.Context {
	if ctx, ok := c.Env[middleware.ContextKey]; ok {
		return ctx.(context.Context)
	}
	panic("context missing!!")
}
