package usecase

import (
	"blog/internal/post"
	"blog/internal/post/repo"
	"context"
)

type Blog interface {
	ListAllBlogs(ctx context.Context, limit, offset int) ([]*post.Blog, error)
}

type blog struct {
	postRepo repo.PostRepo
}

func (b blog) ListAllBlogs(ctx context.Context, limit, offset int) ([]*post.Blog, error) {
	return b.postRepo.List(limit, offset)
}

func NewBlogUsecase(postRepo repo.PostRepo) Blog {
	return &blog{postRepo: postRepo}
}
