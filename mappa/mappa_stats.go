package mappa

import (
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guionardo/mappa_proxy/mappa/repositories"
)

type Stats struct {
	RunningSince  time.Time `json:"running_since"`
	Users         int       `json:"users"`
	LastLogin     time.Time `json:"last_login"`
	LastUserLogin string    `json:"last_user_login"`
}

func GetStats(repository repositories.Repositorie) Stats {

	lastUserLogin, lastLogin := repository.GetLastLogin()

	return Stats{
		RunningSince:  StartedTime,
		Users:         repository.GetUserCount(),
		LastLogin:     lastLogin,
		LastUserLogin: lastUserLogin,
	}
}

func MemoryStatus() gin.H {
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	return gin.H{"alloc": m1.Alloc, "total_alloc": m1.TotalAlloc, "heap_alloc": m1.HeapAlloc}
}
