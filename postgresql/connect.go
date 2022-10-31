package postgresql

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func DoConnectPG(c PostgreSQLConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx", c.DataSourceName)
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	db.SetMaxOpenConns(c.MaxOpenConnections)
	db.SetMaxIdleConns(c.MaxIdleConnections)
	db.SetConnMaxLifetime(time.Minute * time.Duration(c.MaxConnectionLifetimeMinutes))
	return db
}
