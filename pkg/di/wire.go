//go:generate wire
//go:build wireinject

package di

import (
	"github.com/goalistdeveloper/go-template/api"
	"github.com/goalistdeveloper/go-template/ent"
	"github.com/goalistdeveloper/go-template/internal/config"
	"github.com/goalistdeveloper/go-template/internal/db"
	"github.com/goalistdeveloper/go-template/internal/tracing"
	domainProvider "github.com/goalistdeveloper/go-template/pkg/domain/provider"
	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/sdk/trace"

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

// handlers
func InjectHandlers() (*api.Handlers, error) {
	wire.Build(concreteSet, dbSet)
	return &api.Handlers{}, nil
}

func InjectMockHandlers(ctrl *gomock.Controller) (*api.Handlers, error) {
	wire.Build(mockSet, dbSet)
	return &api.Handlers{}, nil
}

// services
func InjectServices() (*domainProvider.Services, error) {
	wire.Build(concreteSet, dbSet)
	return &domainProvider.Services{}, nil
}

func InjectMockServices(ctrl *gomock.Controller) (*domainProvider.Services, error) {
	wire.Build(mockSet, dbSet)
	return &domainProvider.Services{}, nil
}

// repositories
func InjectRepository() (*domainProvider.Repositories, error) {
	wire.Build(concreteSet, dbSet)
	return &domainProvider.Repositories{}, nil
}

func InjectMockRepository(ctrl *gomock.Controller) (*domainProvider.Repositories, error) {
	wire.Build(mockSet, dbSet)
	return &domainProvider.Repositories{}, nil
}

// misc
func InjectDb() (*ent.Client, error) {
	wire.Build(dbSet)
	return &ent.Client{}, nil
}

func InjectTracerProvider() (*trace.TracerProvider, error) {
	wire.Build(tracing.ProvideTracerProvider)
	return &trace.TracerProvider{}, nil
}
