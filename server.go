package main

import (
  "fmt"
	"github.com/codegangsta/martini"
	"github.com/jeffchao/feeds/config/routes"
  "github.com/jeffchao/feeds/config/cassandra"
	"net/http"
)

func main() {
	server := martini.Classic()
	routes.Init(server)

  cql, err := cassandra.Init()
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(cql)

	http.ListenAndServe(":8080", server)
	server.Run()
}
