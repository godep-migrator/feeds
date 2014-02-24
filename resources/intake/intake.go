package intake

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/gocql/gocql"
	"github.com/jeffchao/feeds/models/channel"
	"log"
	"net/http"
)

func Index(cql *gocql.Session, req *http.Request, res http.ResponseWriter, params martini.Params, logger *log.Logger) string {
	logger.Println(req.URL.Query())
	logger.Println(params)

	channels, _ := channel.FindAllChannels(cql)

	log.Printf("%+v", channels)

	response, _ := json.Marshal(channels)

	return string(response)
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
