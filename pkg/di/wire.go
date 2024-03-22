//go:generate wire
//go:build wireinject

package di

import (
	"sample/api"
	"sample/ent"
	"sample/internal/config"
	"sample/internal/db"
	domainProvider "sample/pkg/domain/provider"

	"github.com/google/wire"
)

func InjectHandlers() (*api.Handlers, error) {
	wire.Build(
		api.HandlerSet,
		domainProvider.RepositorySet,
		domainProvider.ServiceSet,

		db.ProvideDb,
		config.ProvideConfig,
	)

	return &api.Handlers{}, nil
}

func InjectDb() (*ent.Client, error) {
	wire.Build(config.ProvideConfig, db.ProvideDb)
	return &ent.Client{}, nil
}
