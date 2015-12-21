// main.go
package main

import (
	"flag"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	mw "github.com/eure/example-blog-golang/framework/middleware"
	"github.com/eure/example-blog-golang/routing"
)

func main() {
	flag.Set("bind", ":1234")

	// api v1 routing
	routeV1 := web.New()
	routeV1.Use(mw.Context)
	routeV1.Use(middleware.SubRouter)
	goji.Handle("/api/v1/*", routeV1)
	routing.SetV1(routeV1)

	// run http server
	goji.Serve()
}
