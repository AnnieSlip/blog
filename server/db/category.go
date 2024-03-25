package db

import (
	"blogs/server/response"
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Category struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
type Categories []Category

func CaregoryList(c context.Context, db *sqlx.DB) (*Categories, error) {
	var cs Categories
	err := selectContext(c, db, &cs, `
	SELECT
	   id,
	   name,
	   created_at
	FROM users`)

	if err != nil {
		return nil, err
	}
	return &cs, nil
}

func (c Category) Response() response.Category {
	return response.Category{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
	}
}

func (cs Categories) Response() response.Categories {
	var resp response.Categories
	for _, c := range cs {
		resp = append(resp, c.Response())

	}
	return resp
}
