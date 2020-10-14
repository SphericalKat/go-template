package main

import (
	"fmt"
	"log"

	"github.com/SphericalKat/go-template/api/endpoints"
	"github.com/SphericalKat/go-template/db"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	// Set global configuration
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicln(fmt.Errorf("fatal error config file: %s", err))
	}

	app := fiber.New()
	app.Get("/", healthCheck)

	// Run db migrations
	log.Println("Running database migrations")
	if err := db.Migrate(); err != nil {
		log.Panic(err)
	}

	endpoints.MountRoutes(app)

	app.Listen(":3000")
}
