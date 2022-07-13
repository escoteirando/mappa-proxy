package repositories

import (
	"sync"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain"
)

type BaseRepository struct {
	sync.RWMutex
	lastLogin     time.Time
	lastUserLogin string
	logins        map[string]domain.LoginData
	DBLock        func()
	DBUnlock      func()
}

func (r *BaseRepository) SetLocking(enabled bool) {
	if enabled {
		r.DBLock = func() {
			r.Lock()
		}
		r.DBUnlock = func() {
			r.Unlock()
		}
	} else {
		r.DBLock = func() {}
		r.DBUnlock = func() {}
	}
}
