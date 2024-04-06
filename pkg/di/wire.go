//go:generate wire
//go:build wireinject

package di

import (
	"github.com/goalistdeveloper/go-template/api"
	"github.com/goalistdeveloper/go-template/ent"
	"github.com/goalistdeveloper/go-template/internal/config"
	"github.com/goalistdeveloper/go-template/internal/db"
	domainProvider "github.com/goalistdeveloper/go-template/pkg/domain/provider"

	"github.com/google/wire"
)

func InjectHandlers() (*api.Handlers, error) {
	wire.Build(
		api.HandlerSet,
		domainProvider.RepositorySet,
		domainProvider.ServiceSet,

		db.ProvideDB,
		config.ProvideConfig,
	)

	return &api.Handlers{}, nil
}

func InjectDb() (*ent.Client, error) {
	wire.Build(config.ProvideConfig, db.ProvideDB)
	return &ent.Client{}, nil
}
