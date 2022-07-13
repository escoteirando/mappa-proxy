package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func NewMappaAuthMiddleware(config MappaAuthMiddlewareConfig) fiber.Handler {
	cfg := configDefault(config)

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Runs only when method is GET
		if c.Method() != "GET" {
			return c.Next()
		}

		// Get authorization header
		auth := string(c.Request().Header.Peek("Authorization"))
		if !cfg.AuthValidator(auth) {
			return cfg.Forbidden(c)
		}
		
		return c.Next()
	}
}
