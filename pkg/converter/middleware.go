package converter

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

var apiKey string

func init() {
	apiKey = os.Getenv("API_KEY")
}

// API Key Middleware
func apiKeyMiddleware(c *fiber.Ctx) error {
	// Get configured API key from environment
	if apiKey != "" {
		// Get API key from request
		reqAPIKey := c.Get("X-API-Key")
		if reqAPIKey == "" {
			reqAPIKey = c.FormValue("apiKey")
		}

		// Validate API key
		if reqAPIKey != apiKey {
			return c.Status(fiber.StatusUnauthorized).
				SendString("Invalid or missing API key")
		}
	}

	// Proceed to next handler if key is valid
	return c.Next()
}
