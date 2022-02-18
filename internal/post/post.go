package post

import (
	"time"
)

type Blog struct {
	Title       string    `sql:"title"`
	Slug        string    `sql:"slug"`
	Body        string    `sql:"body"`
	Description string    `sql:"description"`
	IsDraft     bool      `sql:"IsDraft"`
	AuthorId    string    `sql:"authorId"`
	CreatedAt   time.Time `sql:"createdAt"`
	UpdatedAt   time.Time `sql:"updatedAt"`
}
