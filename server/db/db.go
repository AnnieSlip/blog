package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func selectContext(c context.Context, db *sqlx.DB, dest interface{}, query string, args ...interface{}) error {
	return db.SelectContext(c, dest, query, args...)
}

func getContext(ctx context.Context, db *sqlx.DB, dest interface{}, query string, args ...interface{}) error {
	return db.GetContext(ctx, dest, query, args...)
}
