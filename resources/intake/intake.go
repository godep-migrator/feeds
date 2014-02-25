package intake

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/gocql/gocql"
	"github.com/jeffchao/feeds/models/channel"
	"log"
	"net/http"
)

func Index(cql *gocql.Session, req *http.Request, res http.ResponseWriter, params martini.Params, logger *log.Logger, r render.Render) {
	logger.Println(req.URL.Query())
	logger.Println(params)

	channels, _ := channel.FindAllChannels(cql)

	log.Printf("%+v", channels)

	r.JSON(200, channel.Stringify(channels))
}

func Create(req *http.Request, res http.ResponseWriter, params martini.Params, logger *log.Logger) string {
	logger.Println(params)

	err := req.ParseForm()
	if err != nil {
		return "error"
	}

	logger.Println(req.Form)

	return "created"
}
