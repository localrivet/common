package mysql

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DoConnect(c MysqlConfig) *sqlx.DB {
	db, err := sqlx.Connect("mysql", c.AsDSN())
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	return db
}
