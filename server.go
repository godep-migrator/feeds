package main

import (
	"github.com/codegangsta/martini"
	"github.com/jeffchao/feeds/config/cassandra"
	"github.com/jeffchao/feeds/config/routes"
	"net/http"
)

func main() {
	server := martini.Classic()
	routes.Init(server)

	server.Use(cassandra.CQL())

	http.ListenAndServe(":8080", server)
	server.Run()
}
