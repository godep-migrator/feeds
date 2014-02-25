package routes

import (
	"github.com/codegangsta/martini"
	"github.com/jeffchao/feeds/resources/intake"
)

func Route(server *martini.ClassicMartini) {
	server.Get("/", intake.Index)
	server.Post("/", intake.Create)
}
