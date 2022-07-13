package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// MappaAuthMiddlewareConfig defines the config for middleware.
type MappaAuthMiddlewareConfig struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Unauthorized defines the response body for unauthorized responses.
	// By default it will return with a 401 Unauthorized and the correct WWW-Auth header
	//
	// Optional. Default: nil
	Unauthorized fiber.Handler

	Forbidden fiber.Handler

	// AuthChecker defines a function to check if the authorization is informed
	AuthChecker func(string) bool

	// AuthValidator defines a function to check if the authorization is still valid
	AuthValidator func(string) bool
}

// ConfigDefault is the default config
var ConfigDefault = MappaAuthMiddlewareConfig{
	Next:          nil,
	AuthChecker:   nil,
	AuthValidator: nil,
	Unauthorized:  nil,
	Forbidden:     nil,
}

// Helper function to set default values
func configDefault(config ...MappaAuthMiddlewareConfig) MappaAuthMiddlewareConfig {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}

	if cfg.AuthChecker == nil {
		cfg.AuthChecker = func(auth string) bool {
			return len(auth) > 0
		}
	}
	if cfg.AuthValidator == nil {
		cfg.AuthValidator = func(auth string) bool {
			return len(auth) > 0
		}
	}
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}
	if cfg.Forbidden == nil {
		cfg.Forbidden = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusForbidden)
		}
	}
	return cfg
}
