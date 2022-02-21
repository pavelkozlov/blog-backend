package di

import "blog/internal/post/delivery"

type Application struct {
	BlogHandlers delivery.BlogHandlers
}

func NewApplication(blogHandlers delivery.BlogHandlers) *Application {
	return &Application{
		BlogHandlers: blogHandlers,
	}
}
