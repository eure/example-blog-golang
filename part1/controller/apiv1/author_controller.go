package apiv1

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"github.com/eure/example-blog-golang/controller"
)

// GetAuthor shows author data
func GetAuthor(c web.C, w http.ResponseWriter, r *http.Request) {
	data := controller.NewResponse()
	data.Add("object", "author")
	data.Add("name", c.URLParams["name"])

	ctx := controller.GetContext(c)
	if ctx != nil {
		r2 := ctx.Value("request").(*http.Request)
		data.Add("is_same_request", r == r2)
	}
	controller.RenderJSON(w, data)
}
