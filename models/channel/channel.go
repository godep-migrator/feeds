package channel

import (
	"github.com/gocql/gocql"
	"log"
)

func FindAllChannels(cql *gocql.Session) (bool, error) {
	log.Println(cql)
	return true, nil
}
