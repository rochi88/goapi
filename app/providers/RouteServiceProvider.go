package providers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rochi88/goapi/routes"
)

func RegisterRoutes(app *fiber.App) {
	// Import route files
	importRoutes := []func(*fiber.App){
		routes.RegisterApiRoutes,
	}
	// Register routes
	for _, importRoute := range importRoutes {
		importRoute(app)
	}
}

// NotFoundRoute func for describe 404 Error route.
func NotFoundHandler(c *fiber.Ctx) error {
	// Custom 404 not found message
	url := c.OriginalURL()
	message := "Oops! The requested route '" + url + "' was not found."

	// Return a custom 404 response
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}
