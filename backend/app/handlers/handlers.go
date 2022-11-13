package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type RouteData struct {
	Name      string
	Method    string
	Handler   func(*fiber.Ctx) error
	CacheTime time.Duration
	Mappa     bool
}

var Routes = make(map[string]RouteData)
