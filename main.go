package main

import "github.com/gofiber/fiber/v2"

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	app := fiber.New()
	app.Get("/", healthCheck)

	app.Listen(":3000")
}