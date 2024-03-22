package domain

import (
	"blogs/server/db"
	"blogs/server/request"
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

func UserCreate(c context.Context, d *sqlx.DB, ucr request.UserCreate) (*int, error) {
	pass, err := HashPassword(ucr.Password)
	if err != nil {
		return nil, err
	}

	ucr.Password = pass
	req := ucr.DB()
	req.CreatedAt = time.Now()

	err = db.UserCreate(context.Background(), d, req)
	if err != nil {
		return nil, err
	}

	return nil, nil

}
