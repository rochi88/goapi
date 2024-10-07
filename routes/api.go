package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterApiRoutes(app *fiber.App) {
	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("goapi")
	})
}
