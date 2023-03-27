package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	cacheStorage fiber.Storage
)

func AddAdminHandlers(app *fiber.App, storage fiber.Storage) {
	admin := app.Group("/admin")
	cacheStorage = storage
	admin.Post("/cache/reset", ClearCache)
	admin.Get("/routes", ListRoutes)
	admin.Get("/stats", GetStats)

}

// ClearCache godoc
// @Summary     Clear cache
// @Description Publica mensagem em chat do Telegram
// @Tags        admin
// @Produce     json
// @Success     200 {object} handlers.ReplyMessage
// @Failure     400 {object} handlers.ReplyMessage
// @Router      /admin/cache/reset [post]
func ClearCache(c *fiber.Ctx) error {
	if err := cacheStorage.Reset(); err != nil {
		return reply_error(c, fiber.StatusExpectationFailed, "Failed on reset cache", err)
	}
	return reply_status(c, fiber.StatusOK, "Cache reseted")
}

// ListRoutes godoc
// @Summary List registered routes
// @Tags    admin
// @Produce json
// @Success 200 {object} []string
// @Failure 400 {object} handlers.ReplyMessage
// @Router  /admin/routes [get]
func ListRoutes(c *fiber.Ctx) error {
	allRoutes := c.App().GetRoutes()
	routes := make([]string, 0, len(allRoutes))
	for _, route := range allRoutes {
		if route.Handlers == nil || len(route.Handlers) == 0 {
			continue
		}
		routes = append(routes, fmt.Sprintf("[%s] %s (%v)", route.Method, route.Path, route.Params))
	}

	return c.JSON(routes)
}

// GetStats godoc
// @Summary Data statistics
// @Tags    admin
// @Produce json
// @Success 200 {object} responses.StatsResponse
// @Failure 400 {object} handlers.ReplyMessage
// @Router  /admin/stats [get]
func GetStats(c *fiber.Ctx) error {
	contextData := GetUserContext(c)
	stats, err := contextData.MappaService.GetStats()
	if err != nil {
		return reply_error(c, 500, "Failed to get stats", err)
	}
	return c.JSON(stats)

}
