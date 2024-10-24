package bootstrap

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rochi88/goapi/app/helpers"
	"github.com/rochi88/goapi/app/middlewares"
	"github.com/rochi88/goapi/app/providers"
)

func Application() *fiber.App {
	// Setup basic environment variables
	helpers.SetupEnvFile()
	helpers.SetupDatabase()

	// Create new Fiber instance
	app := fiber.New()

	// Register middlewares
	middlewares.DefaultMiddleware(app)

	// Register routes
	providers.RegisterRoutes(app)

	// Return Fiber instance
	return app
}
