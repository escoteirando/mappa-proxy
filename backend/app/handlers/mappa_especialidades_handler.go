package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func init() {
	Routes["/especialidades"] = RouteData{
		Name:      "Especialidades",		
		Handler:   MappaEspecialidadesHandler,
		CacheTime: 7 * 24 * time.Hour,
		Mappa:     true,
	}
}

// MappaEspecialidades godoc
//	@Summary	Lista de especialidades e items
//	@Tags		db
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string	true	"Authorization"
//	@Success	200				{object}	responses.MappaEspecialidadeResponse
//	@Failure	400				{object}	handlers.ReplyMessage
//	@Router		/mappa/especialidades [get]
func MappaEspecialidadesHandler(c *fiber.Ctx) error {
	contextData := GetUserContext(c)
	especialidades, err := contextData.MappaService.GetEspecialidades()
	if err != nil {
		return reply_error(c, 404, "Falha ao obter especialidades", err)
	}
	return c.JSON(especialidades)
}
