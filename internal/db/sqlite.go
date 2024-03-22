package db

import (
	"log/slog"

	"sample/ent"

	"sample/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

func newSqliteDb(conf *config.Config) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", conf.DatabaseUrl)
	if err != nil {
		slog.Error("error connecting to sqlite database", "err", err)
		return nil, err
	}

	return client, nil
}
