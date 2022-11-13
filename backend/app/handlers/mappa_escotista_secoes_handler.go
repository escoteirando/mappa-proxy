package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func init() {
	Routes["/escotista/:userId/secoes"] = RouteData{
		Name:      "Seções do Escotista",		
		Handler:   MappaEscotistaSecoesHandler,
		CacheTime: 24 * time.Hour,
		Mappa:     true,
	}
}

// MappaEscotistaSecoesHandler godoc
// @Summary Seções do escotista
// @Tags    mappa-proxy
// @Accept  json
// @Param   userId        path   int    true "User ID"
// @Param   Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} responses.MappaSecaoResponse
// @Failure 400 {object} handlers.ReplyMessage
// @Failure 403 {object} handlers.ReplyMessage
// @Router  /mappa/escotista/{userId}/secoes [get]
func MappaEscotistaSecoesHandler(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId", 0)
	if err == nil && userId <= 0 {
		err = fmt.Errorf("Invalid userId %d", userId)
	}

	if err != nil {
		return reply_error(c, 400, "BAD REQUEST", err)
	}
	contextData := GetUserContext(c)
	if err = contextData.NeedsAuth(c); err != nil {
		return err
	}

	secoes, err := contextData.MappaService.GetEscotistaSecoes(userId, contextData.Authorization)
	if err == nil && len(secoes) == 0 {
		err = fmt.Errorf("Não foram encontradas seções para o usuário %d", userId)
	}
	if err != nil {
		return reply_error(c, 400, "BAD REQUEST", err)
	}

	return c.JSON(secoes)

}
