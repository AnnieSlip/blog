package request

import (
	"blogs/server/common"
	"blogs/server/db"
	"time"
)

type BlogCreate struct {
	Author      string         `json:"author"`
	AuthorEmail string         `json:"author_email"`
	Title       string         `json:"title"`
	Content     string         `json:"content"`
	Categories  common.Strings `json:"categories"`
	CreatedAt   *time.Time     `json:"created_at"`
}

func (bc BlogCreate) DB() db.BlogDraft {
	return db.BlogDraft{
		Author:      bc.Author,
		AuthorEmail: bc.AuthorEmail,
		Title:       bc.Title,
		Content:     bc.Content,
		Categories:  bc.Categories,
		CreatedAt:   bc.CreatedAt,
	}

}
