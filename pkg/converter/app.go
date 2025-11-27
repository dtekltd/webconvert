package converter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewApp() *fiber.App {
	// Create Fiber app
	app := fiber.New()

	// Add middleware
	app.Use(logger.New())

	// Routes
	app.Get("/", homeHandler)
	app.Post("/convert", apiKeyMiddleware, convertHandler)

	return app
}
