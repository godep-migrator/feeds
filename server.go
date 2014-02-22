package main

import (
	"github.com/codegangsta/martini"
	"github.com/jeffchao/feeds/config/routes"
	"github.com/jeffchao/feeds/lib/testpkg"
	"net/http"
)

func main() {
	testpkg.TestMessage()

	server := martini.Classic()
	routes.Init(server)

	http.ListenAndServe(":8080", server)
	server.Run()
}
