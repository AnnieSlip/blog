package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func selectContext(c context.Context, db *sqlx.DB, dest interface{}, query string, args ...interface{}) error {
	return db.SelectContext(c, dest, query, args...)
}

func getContext(ctx context.Context, db *sqlx.DB, dest interface{}, query string, args ...interface{}) error {
	return db.GetContext(ctx, dest, query, args...)
}

func namedExecContext(ctx context.Context, db *sqlx.DB, query string, args ...interface{}) (sql.Result, error) {
	return db.NamedExecContext(ctx, query, args)
}
