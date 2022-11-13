package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func init() {
	Routes["/escotista/:userId"] = RouteData{
		Name:      "Escotista",		
		Handler:   MappaEscotistaHandler,
		CacheTime: 24 * time.Hour,
		Mappa:     true,
	}
}

// MappaLogin godoc
// @Summary     Detalhes do escotista
// @Description Informações do escotista, associado e grupo
// @Tags        mappa-proxy
// @Accept      json
// @Param       userId        path   int    true "User ID"
// @Param       Authorization header string true "Authorization"
// @Produce     json
// @Success     200 {object} responses.MappaDetalhesResponse
// @Failure     400 {object} handlers.ReplyMessage
// @Failure     403 {object} handlers.ReplyMessage
// @Router      /mappa/escotista/{userId} [get]
func MappaEscotistaHandler(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId", 0)
	if err == nil {
		if userId <= 0 {
			err = fmt.Errorf("Invalid userId %d", userId)
		}
	}
	authorization, ok := c.GetReqHeaders()["Authorization"]
	if !ok {
		return reply_error(c, 403, "UNAUTHORIZED", fmt.Errorf("Authorization header not found"))
	}
	if err != nil {
		return reply_error(c, 400, "BAD REQUEST", err)
	}
	contextData := GetUserContext(c)

	detalhes := contextData.MappaService.GetEscotistaDetalhes(userId, authorization)
	if detalhes == nil {
		return reply_error(c, 404, "NOT FOUND", fmt.Errorf("Escotista %d not found", userId))
	}

	return c.JSON(detalhes)

}
