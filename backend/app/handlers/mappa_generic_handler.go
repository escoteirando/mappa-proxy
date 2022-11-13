package handlers

import (
	"log"
	"strings"

	"github.com/escoteirando/mappa-proxy/backend/infra"
	"github.com/escoteirando/mappa-proxy/backend/mappa"
	"github.com/gofiber/fiber/v2"
)

func init() {
	Routes["/*"] = RouteData{
		Name:      "Generic",		
		Handler:   MappaGenericHandler,
		CacheTime: 0,
		Mappa:     true,
	}
}

// MappaGeneric godoc
// @Summary     Requisição genérica para a API do Mappa
// @Description Adicione o caminho da requisição após /mappa/
// @Tags        mappa-proxy
// @Accept      json
// @Produce     json
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} interface{}
// @Router      /mappa [get]
func MappaGenericHandler(c *fiber.Ctx) error {
	original := c.OriginalURL()
	log.Printf("original url: %s", original)
	tudo := strings.TrimPrefix(original, "/mappa")
	url := mappa.URL + tudo
	headers := getHeaders(c)
	statusCode, body, err := infra.HttpGet(url, headers, "GET "+url)
	if err == nil {
		c.Response().Header.Add("Content-Type", "application/json")
		c.Status(statusCode)
		c.Write(body)
	} else {
		reply_error(c, statusCode, "MAPPA REQUEST ERROR", err)
	}
	return nil
}

func getHeaders(c *fiber.Ctx) (headers map[string]string) {
	headers = make(map[string]string)

	allowedHeaders := []string{"Authorization", "User-Agent", "Host"}
	for _, s := range allowedHeaders {
		headerValue := string(c.Request().Header.Peek(s))
		if len(headerValue) > 0 {
			headers[s] = headerValue
		}
	}
	return
}
