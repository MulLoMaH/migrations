package table_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Create_table(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	full_name VARCHAR(100) NOT NULL,
	age INTEGER,

	UNIQUE(full_name)
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
