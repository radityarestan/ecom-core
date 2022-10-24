package di

import (
	"github.com/radityarestan/ecom-authentication/internal/repository"
	"github.com/radityarestan/ecom-authentication/internal/service"
	"github.com/radityarestan/ecom-authentication/internal/shared/config"
	"go.uber.org/dig"
)

var (
	Container = dig.New()
)

func init() {
	if err := Container.Provide(config.NewConfig); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewPostgres); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewLogger); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewEcho); err != nil {
		panic(err)
	}

	if err := repository.Register(Container); err != nil {
		panic(err)
	}

	if err := service.Register(Container); err != nil {
		panic(err)
	}
}
