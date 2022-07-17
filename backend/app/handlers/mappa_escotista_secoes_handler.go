package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// MappaEscotistaSecoesHandler godoc
// @Summary      MappaEscotistaSecoes handler
// @Description  Detalhes das seções do escotista
// @Tags         mappa-proxy
// @Accept       json
// @Param userId path int true "User ID"
// @Param Authorization header string true "Authorization"
// @Produce      json
// @Success      200  {object}  responses.MappaSecaoResponse
// @Failure	  	 400  {object}  handlers.ReplyMessage
// @Failure	  	 403  {object}  handlers.ReplyMessage
// @Router       /mappa/escotista/{userId}/secoes [get]
func MappaEscotistaSecoesHandler(c *fiber.Ctx) error {
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
	secoes, err := contextData.MappaService.GetEscotistaSecoes(userId, authorization)
	if err != nil {
		return reply_error(c, 400, "BAD REQUEST", err)
	}
	if len(secoes) == 0 {
		return reply_error(c, 400, "BAD REQUEST", fmt.Errorf("No sections found"))
	}

	return c.JSON(secoes)

}
