package domain

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/infra"
)

type LoginData struct {
	LoginResponse responses.MappaLoginResponse
	UserName      string
	PasswordHash  uint64
	LastLogin     time.Time
	Deleted       bool
}

func (loginData *LoginData) IsValidPassword(password string) bool {
	return infra.GetPasswordHash(password) == loginData.PasswordHash
}


