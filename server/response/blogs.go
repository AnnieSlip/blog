package response

import (
	"blogs/server/common"
	"time"
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

type Blogs []Blog
