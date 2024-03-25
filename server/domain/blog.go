package domain

import (
	"blogs/server/db"
	"blogs/server/request"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func BlogUpdate(c context.Context, d *sqlx.DB, req request.BlogUpdate, id int) error {
	blog, err := db.BlogByID(c, d, id)
	if err != nil {
		return err
	}

	if blog == nil {
		return fmt.Errorf("bad blog id")

	}

	blog.Author = req.Author
	blog.Title = req.Title
	blog.Categories = req.Categories
	blog.Content = req.Content

	err = db.BlogUpdate(c, d, *blog)
	if err != nil {
		return err
	}

	return nil
}
