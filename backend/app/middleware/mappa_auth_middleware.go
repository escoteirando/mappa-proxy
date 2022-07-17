package middleware

import (
	"log"
	"strings"

	"github.com/escoteirando/mappa-proxy/backend/app/handlers"
	"github.com/gofiber/fiber/v2"
)

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

	// IgnoredPaths defines a list of paths to ignore (starting with the value)
	IgnoredPaths []string
}

// ConfigDefault is the default config
var ConfigDefault = MappaAuthMiddlewareConfig{
	Next:          nil,
	AuthChecker:   nil,
	AuthValidator: nil,
	Unauthorized:  nil,
	Forbidden:     nil,
	IgnoredPaths:  nil,
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

	if cfg.IgnoredPaths == nil || len(cfg.IgnoredPaths) == 0 {
		cfg.IgnoredPaths = []string{
			"/mappa/progressoes/"}
	}
	return cfg
}

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

		// Check if path is ignored
		for _, ignore := range cfg.IgnoredPaths {
			if strings.HasPrefix(c.Path(), ignore) {
				return c.Next()
			}
		}

		// Get authorization header
		auth := string(c.Request().Header.Peek("Authorization"))
		if !cfg.AuthValidator(auth) {
			return cfg.Forbidden(c)
		}

		// Use auto Authorization header
		if auth == "auto" {
			contextData := handlers.GetUserContext(c)
			lastLogin := contextData.Cache.GetLastLogin()
			if lastLogin != nil && lastLogin.IsValid() {
				newAuth := lastLogin.Authorization
				if cfg.AuthValidator(newAuth) {
					c.Request().Header.Set("Authorization", newAuth)
					contextData.Authorization = newAuth
					log.Printf("[INFO] MappaAuthMiddleware: Auto authorization changed to %s (%s)", newAuth, lastLogin.UserName)
				} else {
					cfg.Forbidden(c)
				}
			}

		}

		return c.Next()
	}
}
