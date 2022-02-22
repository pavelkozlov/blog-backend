//go:generate easyjson -all ./post.go
package post

import (
	"time"
)

//easyjson:json
type Blogs []*Blog

type Blog struct {
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Body        string    `json:"body"`
	Description string    `json:"description"`
	IsDraft     bool      `json:"is_draft"`
	AuthorId    string    `json:"author_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewBlog struct {
	Title       string     `json:"title" validate:"required"`
	Body        string     `json:"body" validate:"required"`
	IsDraft     bool       `json:"is_draft" validate:"required"`
	Description *string    `json:"description"`
	Slug        string     `json:"-"`
	ID          string     `json:"-"`
	AuthorId    string     `json:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}
