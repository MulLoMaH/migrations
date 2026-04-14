package postgres_connect

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Sql_conn(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:2906@localhost:5432/postgres")
}
