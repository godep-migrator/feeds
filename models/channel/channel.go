package channel

import (
	"github.com/gocql/gocql"
)

const columnFamily = "channels"

type Record struct {
	Id   gocql.UUID
	Name string
}

func FindAllChannels(cql *gocql.Session) ([]Record, error) {
	channels := []Record{}
	channel := Record{}

	iter := cql.Query("SELECT * FROM " + columnFamily).Consistency(gocql.One).Iter()

	for iter.Scan(&channel.Id, &channel.Name) {
		channels = append(channels, channel)
	}

	return channels, nil
}
