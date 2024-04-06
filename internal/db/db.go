package db

import (
	"fmt"

	"github.com/goalistdeveloper/go-template/ent"
	"github.com/goalistdeveloper/go-template/internal/config"
)

var (
	db  *ent.Client
	err error
)

func ProvideDB(conf *config.Config) (*ent.Client, error) {
	if db != nil {
		return db, nil
	}

	switch conf.DatabaseDriver {
	case "sqlite3":
		db, err = newSqliteDB(conf)
	default:
		err = fmt.Errorf("unknown db type: %s", conf.DatabaseDriver)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
