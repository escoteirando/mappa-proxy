package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/gofiber/fiber/v2"
)

func init() {
	Routes["/progressoes/:ramo"] = RouteData{
		Name:      "Progressoes",		
		Handler:   MappaProgressoesHandler,
		CacheTime: 7 * 24 * time.Hour,
		Mappa:     true,
	}
}

// MappaLogin godoc
// @Summary Lista de progressões do ramo
// @Tags    db
// @Accept  json
// @Param   ramo path string true "Ramo" Enums(L,E,S,P)
// @Produce json
// @Success 200 {object} responses.MappaProgressoesResponse
// @Failure 400 {object} handlers.ReplyMessage
// @Router  /mappa/progressoes/{ramo} [get]
func MappaProgressoesHandler(c *fiber.Ctx) error {
	ramo := strings.ReplaceAll(c.Params("ramo", ""), "%22", "")
	if ramo == "" {
		return reply_error(c, 400, "mAPPa request error", fmt.Errorf("Invalid ramo"))
	}

	contextData := GetUserContext(c)
	progressoes, err := contextData.MappaService.GetProgressoes(ramo)
	if err != nil {
		return reply_error(c, 404, "Falha ao obter progressoes", err)
	}
	response := make(responses.MappaProgressoesResponse, len(progressoes))
	for index, progressao := range progressoes {
		response[index] = &responses.MappaProgressaoResponse{
			Codigo:                progressao.Codigo,
			Descricao:             progressao.Descricao,
			CodigoUEB:             progressao.CodigoUEB,
			Ordenacao:             progressao.Ordenacao,
			CodigoCaminho:         progressao.CodigoCaminho,
			CodigoDesenvolvimento: progressao.CodigoDesenvolvimento,
			NumeroGrupo:           progressao.NumeroGrupo,
			CodigoRegiao:          progressao.CodigoRegiao,
			CodigoCompetencia:     progressao.CodigoCompetencia,
			Segmento:              progressao.Segmento,
		}
	}
	return c.JSON(response)
}
