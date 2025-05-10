package pkg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/widcha/armada/configs/database"
)

type PostgresConfig struct {
	Username string
	Password string
	DBName   string
	Host     string
	Port     string
}

func NewPostgres(config PostgresConfig, sqlxKey string) *sqlx.DB {
	client := database.NewConfiguration(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Jakarta",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	), sqlxKey)

	dbConn, err := sqlx.Open("postgres", client.Dsn)
	if err != nil {
		panic(err)
	}

	return dbConn
}
