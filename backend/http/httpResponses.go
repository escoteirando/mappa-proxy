package http

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/cache"
	"github.com/escoteirando/mappa-proxy/backend/configuration"
)

type Stats struct {
	RunningSince  time.Time `json:"running_since"`
	Users         int       `json:"users"`
	LastLogin     time.Time `json:"last_login"`
	LastUserLogin string    `json:"last_user_login"`
}

func GetStats(cache cache.MappaCache) Stats {

	lastLogin := cache.GetLastLogin()

	return Stats{
		RunningSince:  configuration.StartupTime,
		Users:         cache.GetUserCount(),
		LastLogin:     lastLogin.LastLogin,
		LastUserLogin: lastLogin.UserName,
	}
}
