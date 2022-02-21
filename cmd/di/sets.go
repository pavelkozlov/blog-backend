//go:build wireinject
// +build wireinject

package di

import (
	"blog/internal/post/delivery"
	"blog/internal/post/repo"
	"blog/internal/post/usecase"
	"blog/pkg/config"
	"blog/pkg/database"

	"github.com/google/wire"
)

func NewContainer() (*Application, error) {
	wire.Build(
		NewApplication,
		PkgSet,
		RepoSet,
		UseCaseSet,
		DeliverySet,
	)
	return &Application{}, nil
}

var PkgSet = wire.NewSet(
	database.NewPostgresDB,
	config.NewConfig,
)

var RepoSet = wire.NewSet(
	repo.NewPostRepo,
)

var UseCaseSet = wire.NewSet(
	usecase.NewBlogUsecase,
)

var DeliverySet = wire.NewSet(
	delivery.NewBlogHandlers,
)
