package message

import (
	"github.com/gocql/gocql"
)

type Schema struct {
	Id        gocql.UUID `json:"id"`
	Type      string     `json:"type"`
	CreatedAt gocql.UUID `json:"created_at"`
}

func (self Schema) Save() error {
	return nil
}
