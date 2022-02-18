package database

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Postgres interface {
	Acquire() (PgxPoolConn, error)
}

type Rows interface {
	pgx.Rows
}

type PgxPoolConn interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Release()
}
