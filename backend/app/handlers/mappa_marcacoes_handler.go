package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func init() {
	Routes["/marcacoes/:cod_secao"] = RouteData{
		Name:      "Marcações Seção",		
		Handler:   MappaMarcacoesHandler,
		CacheTime: time.Hour,
		Mappa:     true,
	}

}

// MappaMarcacoes godoc
// @Summary     MappaMarcacoes handler
// @Description Lista de marcações da sessão
// @Tags        mappa
// @Accept      json
// @Param       cod_secao     path   int64  true "Código Seção"
// @Param       Authorization header string true "Authorization"
// @Produce     json
// @Success     200 {object} responses.MappaMarcacaoResponse
// @Failure     400 {object} handlers.ReplyMessage
// @Router      /mappa/marcacoes/{cod_secao} [get]
func MappaMarcacoesHandler(c *fiber.Ctx) error {
	codSecao, err := strconv.Atoi(strings.ReplaceAll(c.Params("cod_secao", "0"), "%22", ""))
	if err != nil || codSecao <= 0 {
		return reply_error(c, 400, "mAPPa request error", fmt.Errorf("Invalid codSecao"))
	}

	contextData := GetUserContext(c)
	if err = contextData.NeedsAuth(c); err != nil {
		return err
	}
	marcacoes, err := contextData.MappaService.GetMarcacoes(codSecao, contextData.Authorization)
	if err != nil {
		return reply_error(c, 404, "Falha ao obter marcações", err)
	}
	return c.JSON(marcacoes)
}
