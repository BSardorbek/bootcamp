package db

import (
	"log"
	"postgresql/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func RUN() *sqlx.DB {
	db, err := sqlx.Connect("postgres", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
