//go:generate mockgen -source http.go --destination mock/delivery_mock.go
package delivery

import (
	"blog/internal/post"
	"blog/internal/post/usecase"
	"net/http"
)

type BlogHandlers interface {
	AllPosts() http.HandlerFunc
}

type blogHandlers struct {
	blogUsecase usecase.Blog
}

func NewBlogHandlers(blogUsecase usecase.Blog) BlogHandlers {
	return &blogHandlers{blogUsecase: blogUsecase}
}

func (b blogHandlers) AllPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		blogs, _ := b.blogUsecase.ListAllBlogs(r.Context(), 100, 0)
		js, err := post.Blogs(blogs).MarshalJSON()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		//nolint:errcheck
		w.Write(js)
	}
}
