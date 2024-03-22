package db

import (
	"fmt"

	"sample/ent"
	"sample/internal/config"
)

var (
	db  *ent.Client
	err error
)

func ProvideDb(conf *config.Config) (*ent.Client, error) {
	if db != nil {
		return db, nil
	}

	switch conf.DatabaseDriver {
	case "sqlite":
		db, err = newSqliteDb(conf)
	default:
		err = fmt.Errorf("unknown db type: %s", conf.DatabaseDriver)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
