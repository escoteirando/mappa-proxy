package handlers

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/configuration"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/gofiber/fiber/v2"
)

// Index godoc
// @Summary      Index handler
// @Description  route: /
// @Tags         mappa-proxy
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.IndexResponse
// @Router       / [get]
func IndexHandler(c *fiber.Ctx) error {
	return c.JSON(responses.IndexResponse{
		App:       configuration.APP_NAME,
		Version:   configuration.APP_VERSION,
		RunningBy: time.Since(configuration.StartupTime).String(),
	})
}
