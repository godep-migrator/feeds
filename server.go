package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jeffchao/feeds/config/cassandra"
	"github.com/jeffchao/feeds/config/routes"
	"net/http"
)

func main() {
	server := martini.Classic()
	routes.Route(server)

	server.Use(cassandra.CQL())
	server.Use(render.Renderer(render.Options{IndentJSON: true}))

	http.ListenAndServe(":8080", server)
	server.Run()
}
