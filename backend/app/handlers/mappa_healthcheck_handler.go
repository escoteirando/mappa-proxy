package handlers

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/infra"
	"github.com/escoteirando/mappa-proxy/backend/mappa"
	"github.com/gofiber/fiber/v2"
)

func init() {
	Routes["/hc"] = RouteData{
		Name:      "Health Check",
		Handler:   MappaHealthCheckHandler,
		CacheTime: 10 * time.Second,
		Mappa:     false,
	}
}

// Healthcheck godoc
//	@Summary		Healthcheck handler
//	@Description	Service healthcheck
//	@Tags			mappa-proxy
//	@Produce		json
//	@Success		200	{object}	responses.HealthCheckResponse
//	@Router			/hc [get]
func MappaHealthCheckHandler(c *fiber.Ctx) error {

	statusCode, status, err := infra.Ping(mappa.URL)
	statusHealthy := "HEALTHY"
	if err != nil || statusCode < 1 || statusCode >= 400 {
		statusHealthy = "UNHEALTHY"
	}
	response := responses.HealthCheckResponse{
		Status: statusHealthy,
		MappaServer: responses.MappaServerResponse{
			MappaServerUrl: mappa.URL,
			Status:         status,
			StatusCode:     statusCode,
		},
		Memory: *infra.GetMemoryStatus(),
	}
	c.JSON(response)
	return nil
}
