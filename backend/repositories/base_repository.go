package repositories

import (
	"sync"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain"
)

type BaseRepository struct {
	sync.RWMutex
	useLocking    bool
	lastLogin     time.Time
	lastUserLogin string
	logins        map[string]domain.LoginData
}

func (repository *BaseRepository) DBLock() {
	if repository.useLocking {
		repository.Lock()
	}
}

func (repository *BaseRepository) DBUnlock() {
	if repository.useLocking {
		repository.Unlock()
	}
}
