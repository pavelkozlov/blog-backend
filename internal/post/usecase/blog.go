//go:generate mockgen -source blog.go --destination mock/usecase_mock.go
package usecase

import (
	"blog/internal/post"
	"blog/internal/post/repo"
	"context"
	"time"

	"github.com/Machiel/slugify"
)

type Blog interface {
	ListAllBlogs(ctx context.Context, limit, offset int) ([]*post.Blog, error)
	CreateBlog(ctx context.Context, newBlog *post.NewBlog) (*post.Blog, error)
}

type blog struct {
	postRepo repo.PostRepo
}

func (b blog) CreateBlog(ctx context.Context, newBlog *post.NewBlog) (*post.Blog, error) {
	newBlog.Slug = slugify.Slugify(newBlog.Title + " " + time.Now().String())
	// ToDo fetch it from ctx
	newBlog.AuthorId = "c976c8c1-3c61-4517-8593-5d12f81e4567"
	return b.postRepo.Create(newBlog)
}

func (b blog) ListAllBlogs(ctx context.Context, limit, offset int) ([]*post.Blog, error) {
	return b.postRepo.List(limit, offset)
}

func NewBlogUsecase(postRepo repo.PostRepo) Blog {
	return &blog{postRepo: postRepo}
}
