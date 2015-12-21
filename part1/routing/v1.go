package routing

import (
	"github.com/zenazn/goji/web"

	"github.com/eure/example-blog-golang/controller/apiv1"
)

// SetV1 sets api routing ver1
func SetV1(r *web.Mux) {
	r.Get("/author/:name", apiv1.GetAuthor)
}
