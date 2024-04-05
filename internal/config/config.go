package config

import (
	"fmt"
	"log/slog"
	"os"
	"sample/utils"
	"slices"
	"strings"

	"github.com/go-errors/errors"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	// used by the repositories (required)
	DatabaseUrl    string `koanf:"DATABASE_URL"`
	DatabaseDriver string `koanf:"DATABASE_DRIVER"`
}

var Conf *Config

var k = koanf.New(".")

var allowedDbTypes = []string{
	"sqlite",
	"postgres",
}

var defaultConfigProvider = confmap.Provider(map[string]any{
	"DATABASE_URL":    "file:sqlite.db?mode=rwc&cache=shared&_fk=1",
	"DATABASE_DRIVER": "sqlite",
}, "")

func ProvideConfig() (*Config, error) {
	if Conf == nil {
		err := k.Load(defaultConfigProvider, nil)
		if err != nil {
			slog.Error("error loading default config", "err", err)
			os.Exit(1)
		}

		if err := k.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
			slog.Info("unable to find env file", "err", err)
			slog.Info("falling back to env variables")
		}

		if err := k.Load(env.Provider("", ".", nil), nil); err != nil {
			slog.Error("error loading config", "err", err)
			os.Exit(1)
		}

		Conf = &Config{}
		err = k.Unmarshal("", Conf)
		if err != nil {
			slog.Error("error loading config", "err", err)
			os.Exit(1)
		}
	}

	if !slices.Contains(allowedDbTypes, Conf.DatabaseDriver) {
		return nil, unknownOptionError("database driver", Conf.DatabaseDriver, allowedDbTypes)
	}

	return Conf, nil
}

func unknownOptionError(valueType string, providedValue string, allowedValues []string) error {
	return errors.New(
		fmt.Sprintf("Unknown %s provided to config: '%s'\nMust be one of: %s", valueType, providedValue,
			strings.Join(
				utils.MapOver(allowedValues, func(s string) string { return fmt.Sprintf("'%s'", s) }),
				", ",
			)),
	)
}
