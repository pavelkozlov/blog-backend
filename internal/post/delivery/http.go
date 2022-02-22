//go:generate mockgen -source http.go --destination mock/delivery_mock.go
package delivery

import (
	"blog/internal/post"
	"blog/internal/post/usecase"
	utils "blog/utils/http"
	"net/http"
)

type BlogHandlers interface {
	AllPosts() http.HandlerFunc
	CreatePost() http.HandlerFunc
}

type blogHandlers struct {
	blogUsecase usecase.Blog
}

func (b blogHandlers) CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newBlog := new(post.NewBlog)
		if err := utils.ParseAndValidateBody(newBlog, newBlog.UnmarshalJSON, r); err != nil {
			utils.WriteErrResponse(err, w, 400)
			return
		}
		blog, err := b.blogUsecase.CreateBlog(r.Context(), newBlog)
		if err != nil {
			utils.WriteErrResponse(err, w, 500)
			return
		}
		utils.WriteOkResponse(blog, w)
	}
}

func NewBlogHandlers(blogUsecase usecase.Blog) BlogHandlers {
	return &blogHandlers{blogUsecase: blogUsecase}
}

func (b blogHandlers) AllPosts() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		limit, offset, err := utils.ParseLimitAndOffset(r)
		if err != nil {
			utils.WriteErrResponse(err, w, 400)
			return
		}
		blogs, err := b.blogUsecase.ListAllBlogs(r.Context(), limit, offset)
		if err != nil {
			utils.WriteErrResponse(err, w, 500)
			return
		}
		utils.WriteOkResponse(blogs, w)
	}
}
