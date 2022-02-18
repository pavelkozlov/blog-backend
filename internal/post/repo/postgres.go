package repo

import (
	"blog/internal/post"
	"blog/pkg/database"
	"context"
)

type PostRepo interface {
	List(limit, offset int) ([]*post.Blog, error)
	Create(input *blog) error
}

type postRepo struct {
	db       database.Postgres
	blogPool *blogPool
}

func NewPostRepo(db database.Postgres) PostRepo {
	return &postRepo{
		db:       db,
		blogPool: newBlogPool(),
	}
}

func (p postRepo) List(limit, offset int) ([]*post.Blog, error) {
	conn, err := p.db.Acquire()
	if err != nil {
		// ToDo log
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), SELECT_QUERY, limit, offset)
	if err != nil {
		// ToDo log
		return nil, err
	}
	defer rows.Close()

	posts := make([]*post.Blog, 0, limit)

	model := p.blogPool.Get()
	defer p.blogPool.Put(model)

	for rows.Next() {
		model = &blog{}
		if err = rows.Scan(&model.ID, &model.Title, &model.Slug,
			&model.Body, &model.Description, &model.IsDraft, &model.AuthorId,
			&model.CreatedAt, &model.UpdatedAt, &model.DeletedAt); err != nil {
			// ToDo log me
			continue
		}
		posts = append(posts, model.asDomainModel())
	}

	return posts, rows.Err()
}

func (p postRepo) Create(input *blog) error {
	//TODO implement me
	return nil
}
