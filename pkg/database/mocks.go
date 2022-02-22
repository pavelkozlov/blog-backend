package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Postgres interface {
	Acquire() (PgxPoolConn, error)
}

type Row interface {
	pgx.Rows
}

type Rows interface {
	pgx.Rows
}

type PgxPoolConn interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Release()
}
