package bootstrap

import (
	"github.com/centrex/webcore/core/database"
	"github.com/centrex/webcore/core/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rochi88/goapi/app/providers"
)

func Application() *fiber.App {
	// Setup basic environment variables
	env.SetupEnvFile()
	database.SetupDatabase()

	// Create new Fiber instance
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	// Register routes
	providers.RegisterRoutes(app)

	// Return Fiber instance
	return app
}
