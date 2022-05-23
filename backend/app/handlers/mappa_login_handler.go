package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/escoteirando/mappa-proxy/backend/domain/requests"
	"github.com/escoteirando/mappa-proxy/backend/mappa"

	"github.com/gofiber/fiber/v2"
)

// MappaLogin godoc
// @Summary      Mappa Login handler
// @Description  User login
// @Tags         mappa-proxy
// @Accept       json
// @Param request body requests.LoginRequest true "Login request"
// @Produce      json
// @Success      200  {object}  responses.MappaLoginResponse
// @Failure	  	 400  {object}  handlers.ReplyMessage
// @Failure	  	 403  {object}  handlers.ReplyMessage
// @Router       /mappa/login [post]
func MappaLoginHandler(c *fiber.Ctx) error {
	requestBody := c.Body()
	var loginRequest requests.LoginRequest
	err := json.Unmarshal(requestBody, &loginRequest)
	if err != nil {
		return reply_error(c, 400, "mAPPa request error", err)
	}
	contextData := getUserContext(c)
	loginData := contextData.Cache.GetLogin(loginRequest.UserName)
	if loginData != nil {
		if !loginData.IsValidPassword(loginRequest.Password) {
			return reply_error(c, 403, "UNAUTHORIZED", fmt.Errorf("Invalid user or password"))
		}
		return c.JSON(loginData.LoginResponse)
	}

	loginResponse, err := mappa.MappaLoginRequest(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		return reply_error(c, 403, "UNAUTHORIZED", err)
	}

	contextData.Cache.SetLogin(loginRequest.UserName, loginRequest.Password, loginResponse)
	return c.JSON(loginResponse)
}

// MappaLogin godoc
// @Summary      MappaEscotista handler
// @Description  Detalhes do escotista
// @Tags         mappa-proxy
// @Accept       json
// @Param userId path int true "User ID"
// @Param Authorization header string true "Authorization"
// @Produce      json
// @Success      200  {object}  responses.MappaDetalhesResponse
// @Failure	  	 400  {object}  handlers.ReplyMessage
// @Failure	  	 403  {object}  handlers.ReplyMessage
// @Router       /mappa/escotista/{userId} [get]
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
	contextData := getUserContext(c)
	detalhes, err := contextData.Cache.GetDetalheEscotista(userId)
	if err != nil && detalhes == nil {
		detalhes = mappa.MappaEscotistaRequest(userId, authorization)
		if detalhes == nil {
			return reply_error(c, 404, "NOT FOUND", fmt.Errorf("Escotista %d not found", userId))
		}
		contextData.Cache.SetDetalheEscotista(userId, *detalhes)
	}
	return c.JSON(detalhes)

}
