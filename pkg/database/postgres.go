package database

import (
	"blog/pkg/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type Postgres struct {
	Db *sqlx.DB
}

func NewPostgresDB(config *config.Config) *Postgres {
	db, err := sqlx.Connect("pgx", config.Database.Url)
	if err != nil {
		log.Fatalln(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}
	return &Postgres{Db: db}
}
