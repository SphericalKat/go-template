//go:generate wire
//go:build wireinject

package di

import (
	"github.com/goalistdeveloper/go-template/api"
	"github.com/goalistdeveloper/go-template/ent"
	"github.com/goalistdeveloper/go-template/internal/config"
	"github.com/goalistdeveloper/go-template/internal/db"
	domainProvider "github.com/goalistdeveloper/go-template/pkg/domain/provider"
	"github.com/golang/mock/gomock"

	"github.com/google/wire"
)

var dbSet = wire.NewSet(
	db.ProvideDB,
	config.ProvideConfig,
)

var concreteSet = wire.NewSet(
	api.HandlerSet,
	domainProvider.RepositorySet,
	domainProvider.ServiceSet,
)

var mockSet = wire.NewSet(
	api.MockHandlerSet,
	domainProvider.MockRepositorySet,
	domainProvider.MockServiceSet,
)

func InjectHandlers() (*api.Handlers, error) {
	wire.Build(concreteSet, dbSet)
	return &api.Handlers{}, nil
}

func InjectMockHandlers(ctrl *gomock.Controller) (*api.Handlers, error) {
	wire.Build(mockSet, dbSet)
	return &api.Handlers{}, nil
}

func InjectDb() (*ent.Client, error) {
	wire.Build(dbSet)
	return &ent.Client{}, nil
}
