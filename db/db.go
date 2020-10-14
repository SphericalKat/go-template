package db

import (
	"context"
	"log"

	"github.com/SphericalKat/go-template/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/spf13/viper"
)

var db *pg.DB = nil

// GetDB returns a singleton reference to the database
func GetDB() *pg.DB {
	if db != nil {
		return db
	}

	opt, err := pg.ParseURL(viper.GetString("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db = pg.Connect(opt)
	return db
}

// Migrate runs database migrations
func Migrate() error {
	models := []interface{}{
		(*models.User)(nil),
	}

	ctx := context.Background()

	// Check if DB connection is up and running
	if err := GetDB().Ping(ctx); err != nil {
		log.Panic(err)
	}

	// Create UUID generation extensions
	if _, err := GetDB().ExecContext(ctx, "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"); err != nil {
		log.Panic(err)
	}

	for _, model := range models {
		err := GetDB().Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
