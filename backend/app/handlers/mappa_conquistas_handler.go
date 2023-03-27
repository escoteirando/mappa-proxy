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
// @Tags    mappa
// @Accept  json
// @Param   cod_secao     path   int64  true  "Código Seção"
// @Param   Authorization header string true  "Authorization"
// @Param   desde         query  string false "Data de início do período (YYYY-MM-DD) padrão 1 ano atrás"
// @Produce json
// @Success 200 {object} []responses.FullConquistaResponse
// @Failure 400 {object} handlers.ReplyMessage
// @Router  /mappa/conquistas/{cod_secao} [get]
func MappaConquistasHandler(c *fiber.Ctx) error {
	codSecao, err := strconv.Atoi(strings.ReplaceAll(c.Params("cod_secao", "0"), "%22", ""))
	if err != nil || codSecao <= 0 {
		return reply_error(c, 400, "mAPPa request error", fmt.Errorf("Invalid codSecao"))
	}
	desdeQuery := c.Query("desde", time.Now().Add(-time.Hour*24*365).Format("2006-01-02"))
	since, err := time.Parse("2006-01-02", desdeQuery)
	if err != nil {
		return reply_error(c, 400, "mAPPa request error", fmt.Errorf("Invalid desde"))
	}
	contextData := GetUserContext(c)
	if err = contextData.NeedsAuth(c); err != nil {
		return err
	}
	marcacoes, err := contextData.MappaService.GetConquistasFull(codSecao, contextData.Authorization, since)
	if err != nil {
		return reply_error(c, 404, "Falha ao obter marcações", err)
	}
	return c.JSON(marcacoes)
}
