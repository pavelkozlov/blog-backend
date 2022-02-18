//go:generate mockgen -source mocks.go --destination mock/postgres_mock.go
package database

import (
	"blog/pkg/config"
	"blog/pkg/logger"
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type PostgresDatabase struct {
	db                *pgxpool.Pool
	ActiveConnections *int64
}

func (p PostgresDatabase) Acquire() (PgxPoolConn, error) {
	return p.db.Acquire(context.Background())
}

func NewPostgresDB(config *config.Config) Postgres {

	lll := logger.NewLogger(config)

	cfg, err := pgxpool.ParseConfig(config.Database.Url)
	if err != nil {
		log.Fatalln(err)
	}
	cfg.MinConns = 0
	cfg.MaxConns = 12

	cfg.ConnConfig.Logger = zapadapter.NewLogger(lll.Zap())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}
	if err = conn.Ping(ctx); err != nil {
		log.Fatalln(err)
	}

	return &PostgresDatabase{
		db:                conn,
		ActiveConnections: new(int64),
	}
}
