package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func init() {
	Routes["/conquistas/:cod_secao"] = RouteData{
		Name:      "Conquistas Seção",		
		Handler:   MappaConquistasHandler,
		CacheTime: time.Hour,
		Mappa:     true,
	}

}

// MappaConquistas godoc
// @Summary Lista de conquistas da secão
// @Tags    mappa-proxy
// @Accept  json
// @Param   cod_secao     path   int64  true "Código Seção"
// @Param   Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} responses.MappaMarcacaoResponse
// @Failure 400 {object} handlers.ReplyMessage
// @Router  /mappa/conquistas/{cod_secao} [get]
func MappaConquistasHandler(c *fiber.Ctx) error {
	codSecao, err := strconv.Atoi(strings.ReplaceAll(c.Params("cod_secao", "0"), "%22", ""))
	if err != nil || codSecao <= 0 {
		return reply_error(c, 400, "mAPPa request error", fmt.Errorf("Invalid codSecao"))
	}

	contextData := GetUserContext(c)
	if err = contextData.NeedsAuth(c); err != nil {
		return err
	}
	marcacoes, err := contextData.MappaService.GetConquistas(codSecao, contextData.Authorization)
	if err != nil {
		return reply_error(c, 404, "Falha ao obter marcações", err)
	}
	return c.JSON(marcacoes)
	// TODO: Implementar dados agregados das conquistas, incluindo os associados e a especialidade
}