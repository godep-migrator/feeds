package channel

import (
	"encoding/json"
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

func Stringify(channels []Record) string {
	result := make([]map[string]interface{}, 0, len(channels))
	channel := make(map[string]interface{})

	for _, c := range channels {
		channel["id"] = c.Id.String()
		channel["name"] = c.Name
		result = append(result, channel)
	}

	resultJSON, _ := json.Marshal(&result)
	return string(resultJSON)
}
