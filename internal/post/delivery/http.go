package delivery

import (
	"blog/internal/post/usecase"
	"encoding/json"
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
		js, err := json.Marshal(blogs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
