package repo

import (
	"blog/internal/post"
	"sync"
	"time"
)

type Blog struct {
	ID          string     `sql:"id"`
	Title       string     `sql:"title"`
	Slug        string     `sql:"slug"`
	Body        string     `sql:"body"`
	Description string     `sql:"description"`
	IsDraft     bool       `sql:"IsDraft"`
	AuthorId    string     `sql:"authorId"`
	CreatedAt   time.Time  `sql:"createdAt"`
	UpdatedAt   time.Time  `sql:"updatedAt"`
	DeletedAt   *time.Time `sql:"deletedAt"`
}

func (b Blog) asDomainModel() *post.Blog {
	return &post.Blog{
		Title:       b.Title,
		Slug:        b.Slug,
		Body:        b.Body,
		Description: b.Description,
		IsDraft:     b.IsDraft,
		AuthorId:    b.AuthorId,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}

func (b *Blog) fromDomainModel(input *post.NewBlog) {
	b.ID = input.ID
	b.Title = input.Title
	b.Slug = input.Slug
	b.Body = input.Body
	if input.Description == nil {
		b.Description = ""
	} else {
		b.Description = *input.Description
	}
	b.IsDraft = input.IsDraft
	b.AuthorId = input.AuthorId
	b.CreatedAt = input.CreatedAt
	b.UpdatedAt = input.UpdatedAt
	b.DeletedAt = input.DeletedAt
}

type blogPool struct {
	pool *sync.Pool
}

func (b blogPool) Get() *Blog {
	return b.pool.Get().(*Blog)
}

func (b blogPool) Put(old *Blog) {
	*old = Blog{}
	b.pool.Put(old)
}

func newBlogPool() *blogPool {
	pool := &sync.Pool{New: func() interface{} {
		return new(Blog)
	}}
	return &blogPool{pool: pool}
}
