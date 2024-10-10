package handlers

import (
	"encoding/json"
	"errors"

	"github.com/escoteirando/mappa-proxy/backend/domain/requests"

	"github.com/gofiber/fiber/v2"
)

func init() {
	Routes["/login"] = RouteData{
		Name:      "Login",
		Method:    "POST",
		Handler:   MappaLoginHandler,
		CacheTime: 0,
		Mappa:     true,
	}
}

// MappaLogin godoc
//	@Summary		Mappa Login handler
//	@Description	User login
//	@Tags			mappa
//	@Accept			json
//	@Param			request	body	requests.LoginRequest	true	"Login request"
//	@Produce		json
//	@Success		200	{object}	responses.MappaLoginResponse
//	@Failure		400	{object}	handlers.ReplyMessage
//	@Failure		403	{object}	handlers.ReplyMessage
//	@Router			/mappa/login [post]
func MappaLoginHandler(c *fiber.Ctx) error {
	requestBody := c.Body()
	var loginRequest requests.LoginRequest
	err := json.Unmarshal(requestBody, &loginRequest)
	if err != nil {
		return reply_error(c, 400, "mAPPa request error", err)
	}
	contextData := GetUserContext(c)
	loginData := contextData.Cache.GetLogin(loginRequest.UserName)
	if loginData != nil {
		if !loginData.IsValidPassword(loginRequest.Password) {
			return reply_error(c, 403, "UNAUTHORIZED", errors.New("invalid user or password"))
		}

		return c.JSON(loginData.GetLoginResponse())
	}

	loginResponse, err := contextData.MappaService.Login(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		return reply_error(c, 403, "UNAUTHORIZED", err)
	}

	return c.JSON(loginResponse)
}
