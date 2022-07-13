package domain

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/infra"
)

// LoginData represents a login read from response.MappaLoginResponse
type LoginData struct {	
	UserName      string
	UserId        int
	Authorization string
	PasswordHash  uint32
	LastLogin     time.Time
	ValidUntil    time.Time
	Deleted       bool
}

func (loginData *LoginData) IsValidPassword(password string) bool {
	return infra.GetPasswordHash(password) == loginData.PasswordHash
}

func (loginData *LoginData) GetLoginResponse() responses.MappaLoginResponse {
	return responses.MappaLoginResponse{
		ID:      loginData.Authorization,
		TTL:     int(loginData.ValidUntil.Sub(loginData.LastLogin).Seconds()),
		Created: loginData.LastLogin,
		Userid:  loginData.UserId,
	}
}

func NewLoginData(loginResponse responses.MappaLoginResponse, userName string, password string) *LoginData {
	return &LoginData{
		// LoginResponse: loginResponse,
		UserName:      userName,
		UserId:        loginResponse.Userid,
		Authorization: loginResponse.ID,
		PasswordHash:  infra.GetPasswordHash(password),
		LastLogin:     loginResponse.Created,
		ValidUntil:    loginResponse.Created.Add(time.Duration(loginResponse.TTL) * time.Second),
	}
}

func (loginData *LoginData) IsValid() bool {
	return loginData.ValidUntil.After(time.Now())
}
