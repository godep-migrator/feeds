package cassandra

import (
  "fmt"
  "encoding/json"
  "github.com/gocql/gocql"
  "os"
  "path/filepath"
)

type CassandraConfig struct {
  Keyspace string "keyspace"
  Hosts []interface{} "hosts"
}

type Config map[string]CassandraConfig

func Init() (*gocql.Session, error) {
  jsonFile, err := filepath.Abs("config/cassandra/cassandra.json")
  if err != nil {
    fmt.Println(err)
    return nil, err
  }

  file, err := os.Open(jsonFile)
  if err != nil {
    fmt.Println(err)
    return nil, err
  }

  decoder := json.NewDecoder(file)
  config := &Config{}

  err = decoder.Decode(&config)
  if err != nil {
    fmt.Println(err)
    return nil, err
  }

  env := os.Getenv("GO_ENV")
  if env == "" {
    env = "development"
  }

  fmt.Printf("Using %s config: %+v\n", env, (*config)[env])

  cluster := gocql.NewCluster((*config)[env].Hosts[0].(string))
  cluster.Keyspace = "feeds_" + env
  session, _ := cluster.CreateSession()
  defer session.Close()

  return session, err
}
