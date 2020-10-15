package endpoints

import (
	"github.com/go-pg/pg/v10"
	"net/http"

	"github.com/SphericalKat/go-template/crud"
	"github.com/SphericalKat/go-template/schemas"
	"github.com/gofiber/fiber/v2"
)

func createUser(c *fiber.Ctx) error {
	u := new(schemas.UserCreate)

	if err := c.BodyParser(u); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"msg": "invalid json",
			"err": err.Error(),
		})
	}

	created, err := crud.CreateUser(u)
	if err != nil {
		if err == pg.ErrNoRows {
			return c.Status(http.StatusConflict).JSON(fiber.Map{
				"msg": "user with this username already exists",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"msg": "something went wrong",
			"err": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":  "user created successfully",
		"user": created,
	})
}

func updateUser(c *fiber.Ctx) error {
	u := new(schemas.UserUpdate)

	if err := c.BodyParser(u); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"msg": "invalid json",
			"err": err.Error(),
		})
	}

	updated, err := crud.UpdateUser(u)
	if err != nil {
		if err == pg.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"msg": "user not found",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"msg": "something went wrong",
			"err": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":  "user updated successfully",
		"user": updated,
	})
}

func getUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	user, err := crud.GetUser(userID)
	if err != nil {
		if err == pg.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"msg": "user not found",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"msg": "something went wrong",
			"err": err.Error(),
		})
	}

	return c.JSON(user)
}

func deleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	deleted, err := crud.DeleteUser(userID)
	if err != nil {
		if err == pg.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"msg": "user not found",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"msg": "something went wrong",
			"err": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg": "successfully deleted user",
		"user": deleted,
	})
}

// MountRoutes mounts all routes declared here
func MountRoutes(app *fiber.App) {
	app.Post("/api/user", createUser)
	app.Patch("/api/user", updateUser)
	app.Get("/api/user/:id", getUser)
}
