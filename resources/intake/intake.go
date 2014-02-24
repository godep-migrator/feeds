package intake

import (
	"github.com/codegangsta/martini"
	"github.com/gocql/gocql"
	"log"
	"net/http"
)

func Index(cql *gocql.Session, req *http.Request, res http.ResponseWriter, params martini.Params, logger *log.Logger) string {
	logger.Println(req.URL.Query())
	logger.Println(params)

	return "using intake.Index()"
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
