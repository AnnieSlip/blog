package db

import (
	"blogs/server/common"
	"context"
	"database/sql"
	"time"

	"blogs/server/response"

	"github.com/jmoiron/sqlx"
)

type Blog struct {
	ID          int            `db:"id"`
	Author      string         `db:"author"`
	AuthorEmail string         `db:"author_email"`
	Title       string         `db:"title"`
	Content     string         `db:"content"`
	Categories  common.Strings `db:"categories"`
	CreatedAt   *time.Time     `db:"created_at"`
}

type BlogDraft struct {
	Author      string         `db:"author"`
	AuthorEmail string         `db:"author_email"`
	Title       string         `db:"title"`
	Content     string         `db:"content"`
	Categories  common.Strings `db:"categories"`
	CreatedAt   *time.Time     `db:"created_at"`
}

type Blogs []Blog

func BlogList(c context.Context, db *sqlx.DB) (*Blogs, error) {
	var bs Blogs
	err := selectContext(c, db, &bs, `
	SELECT 
	   id,
	   author,
	   title,
	   content,
	   categories,
	   author_email,
	   created_at
	FROM public.blogs`)

	if err != nil {
		return nil, err
	}
	return &bs, nil
}

func BlogByID(c context.Context, db *sqlx.DB, id int) (*Blog, error) {
	var b Blog
	err := getContext(c, db, &b, `
	SELECT
	    id,
	    author,
	    title,
	    content,
	    categories,
	    author_email,
	    created_at
	FROM public.blogs
	WHERE id=$1`, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &b, nil
}

func BlogUpdate(c context.Context, db *sqlx.DB, blog Blog) error {
	_, err := namedExecContext(c, db, `
	UPDATE  blogs
	SET
	    author = :author
		title = :title,
		content = :content,
		categories = :categories,
		author_email = :author_email
	WHERE
	  	id = :id`, blog)
	if err != nil {
		return err
	}
	return nil
}

func BlogAdd(c context.Context, db *sqlx.DB, req BlogDraft) error {
	_, err := db.NamedExecContext(c, `
        INSERT INTO blogs
            (
			 author,
             title,
             content,
             categories,
             author_email,
             created_at)
        VALUES
            (:author, :title, :content, :categories, :author_email, :created_at)
    `, req)
	if err != nil {
		return err
	}
	return nil
}

func (b Blog) Response() response.Blog {
	return response.Blog{
		ID:         b.ID,
		Author:     b.Author,
		Title:      b.Title,
		Content:    b.Content,
		Categories: b.Categories,
		CreatedAt:  b.CreatedAt,
	}
}

func (bs Blogs) Response() response.Blogs {
	var blogs response.Blogs
	for _, b := range bs {
		blogs = append(blogs, b.Response())
	}
	return blogs
}
