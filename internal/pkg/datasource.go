package pkg

import (
	"github.com/jmoiron/sqlx"
	"github.com/widcha/armada/configs"
)

const (
	projectXSqlxKey = "projectx_sqlx"
)

type DataSource struct {
	Postgre *sqlx.DB
}

func NewDataSource() *DataSource {
	postgresClient := NewPostgres(PostgresConfig{
		Username: configs.Get().DBUsername,
		Password: configs.Get().DBPassword,
		DBName:   configs.Get().DBName,
		Host:     configs.Get().DBHost,
		Port:     configs.Get().DBPort,
	}, projectXSqlxKey)

	return &DataSource{
		Postgre: postgresClient,
	}
}
