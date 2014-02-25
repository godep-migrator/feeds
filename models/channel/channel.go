package channel

import (
	"github.com/gocql/gocql"
)

const columnFamily = "channels"

type Record struct {
	Id   gocql.UUID `json:"id"`
	Name string `json:"name"`
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

func Stringify(channels []Record) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(channels))
	channel := make(map[string]interface{})

	for _, c := range channels {
		channel["id"] = c.Id.String()
		channel["name"] = c.Name
		result = append(result, channel)
	}

	return result
}
