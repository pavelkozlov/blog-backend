// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"blog/internal/post/delivery"
	"blog/internal/post/repo"
	"blog/internal/post/usecase"
	"blog/pkg/config"
	"blog/pkg/database"

	"github.com/google/wire"
)

// Injectors from sets.go:

func NewContainer() (*Application, error) {
	configConfig := config.NewConfig()
	postgres := database.NewPostgresDB(configConfig)
	postRepo := repo.NewPostRepo(postgres)
	blog := usecase.NewBlogUsecase(postRepo)
	blogHandlers := delivery.NewBlogHandlers(blog)
	application := NewApplication(blogHandlers)
	return application, nil
}

// sets.go:

var PkgSet = wire.NewSet(database.NewPostgresDB, config.NewConfig)

var RepoSet = wire.NewSet(repo.NewPostRepo)

var UseCaseSet = wire.NewSet(usecase.NewBlogUsecase)

var DeliverySet = wire.NewSet(delivery.NewBlogHandlers)
