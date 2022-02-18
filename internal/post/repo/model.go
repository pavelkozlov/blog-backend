package repo

import (
	"database/sql"
	"sync"
	"time"
)

type blog struct {
	ID          string       `sql:"id"`
	Title       string       `sql:"title"`
	Slug        string       `sql:"slug"`
	Body        string       `sql:"body"`
	Description string       `sql:"description"`
	IsDraft     bool         `sql:"IsDraft"`
	AuthorId    string       `sql:"authorId"`
	CreatedAt   time.Time    `sql:"createdAt"`
	UpdatedAt   time.Time    `sql:"updatedAt"`
	DeletedAt   sql.NullTime `sql:"deletedAt"`
}

type blogPool struct {
	pool *sync.Pool
}

func (b blogPool) Get() *blog {
	return b.pool.Get().(*blog)
}

func (b blogPool) Put(old *blog) {
	*old = blog{}
	b.pool.Put(old)
}

func newBlogPool() *blogPool {
	pool := &sync.Pool{New: func() interface{} {
		return new(blog)
	}}
	return &blogPool{pool: pool}
}
