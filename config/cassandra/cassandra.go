package cassandra

import (
	"encoding/json"
	"fmt"
  "github.com/codegangsta/martini"
	"github.com/gocql/gocql"
	"os"
	"path/filepath"
)

type CassandraConfig struct {
	Keyspace string        "keyspace"
	Hosts    []interface{} "hosts"
}

type Config map[string]CassandraConfig

func CQL() martini.Handler {
	jsonFile, err := filepath.Abs("config/cassandra/cassandra.json")
	if err != nil {
    panic(err)
	}

	file, err := os.Open(jsonFile)
	if err != nil {
    panic(err)
	}

	decoder := json.NewDecoder(file)
	config := &Config{}

	err = decoder.Decode(&config)
	if err != nil {
    panic(err)
	}

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	fmt.Printf("Using %s config: %+v\n", (*config)[env].Keyspace, (*config)[env])

	cluster := gocql.NewCluster((*config)[env].Hosts[0].(string))
	cluster.Keyspace = (*config)[env].Keyspace

	return func(context martini.Context) {
    session, _ := cluster.CreateSession()
    context.Map(session)
    defer session.Close()
    context.Next()
  }
}
