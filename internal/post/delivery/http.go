//go:generate mockgen -source http.go --destination mock/delivery_mock.go
package delivery

import (
	"blog/internal/post/usecase"
	utils "blog/utils/http"
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
