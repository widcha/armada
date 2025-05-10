package database

import (
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
)

var store sync.Map

type Configuration struct {
	Key     string `json:"key" yaml:"key"` // Key is connection key identification
	Dsn     string `json:"dsn" yaml:"dsn"` // Dsn configuration using uri
	SqlxKey string
}

func NewConfiguration(dsn, sqlxKey string) *Configuration {
	return &Configuration{
		Dsn:     dsn,
		SqlxKey: sqlxKey,
	}
}

func GetSqlxClient(key string) *sqlx.DB {
	client, ok := store.Load(key)
	if !ok {
		log.Fatal("please open db first", key)
	}

	return client.(*sqlx.DB)
}
