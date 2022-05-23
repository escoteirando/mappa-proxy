package cache

import (
	"fmt"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/repositories"
	"github.com/escoteirando/mappa-proxy/backend/tools"
)

type MappaCache struct {
	repository repositories.IRepository
	logins     map[string]*domain.LoginData
	lastLogin  *domain.LoginData
}

func CreateMappaCache(repository repositories.IRepository) (cache *MappaCache, err error) {
	cache = &MappaCache{
		repository: repository,
		logins:     make(map[string]*domain.LoginData),
	}
	return
}

func (cache *MappaCache) updateLastLogin(loginData *domain.LoginData) {
	if loginData.Deleted || !loginData.LoginResponse.IsValid() {
		return
	}
	if cache.lastLogin == nil || cache.lastLogin.LastLogin.Before(loginData.LastLogin) {
		cache.lastLogin = loginData
	}
}

func (cache *MappaCache) Load() error {
	logins, err := cache.repository.GetLogins()
	if err != nil {
		return err
	}
	for _, login := range logins {
		if login.LoginResponse.IsValid() {
			cache.logins[login.UserName] = login
			cache.updateLastLogin(login)
			continue
		}
		if _, ok := cache.logins[login.UserName]; ok {
			delete(cache.logins, login.UserName)
		}
		cache.repository.DeleteLogin(login.UserName)
	}
	return nil
}

func (cache *MappaCache) GetLogin(username string) *domain.LoginData {
	if login, ok := cache.logins[username]; ok {
		if login.LoginResponse.IsValid() {
			return login
		}
		delete(cache.logins, login.UserName)
		cache.repository.DeleteLogin(login.UserName)
	}

	return nil
}

func (cache *MappaCache) SetLogin(username string, password string, login responses.MappaLoginResponse) error {
	if !login.IsValid() {
		return fmt.Errorf("Invalid login %v", login)
	}
	newLoginData := &domain.LoginData{
		LoginResponse: login,
		UserName:      username,
		Deleted:       false,
		PasswordHash:  tools.GetPasswordHash(password),
		LastLogin:     time.Now(),
	}
	cache.updateLastLogin(newLoginData)
	cache.logins[username] = newLoginData
	cache.repository.SetLogin(username, password, login, newLoginData.LastLogin)
	return nil
}

func (cache *MappaCache) GetLastLogin() *domain.LoginData {
	return cache.lastLogin
}

func (cache *MappaCache) GetUserCount() int {
	return len(cache.logins)
}

func (cache *MappaCache) SetDetalheEscotista(userId int, detalheEscotista responses.MappaDetalhesResponse) error {
	return cache.repository.SetDetalheEscotista(uint(userId), detalheEscotista)
}

func (cache *MappaCache) GetDetalheEscotista(userId int) (*responses.MappaDetalhesResponse, error) {
	return cache.repository.GetDetalheEscotista(uint(userId))
}

func (cache *MappaCache) GetLastEventTime(eventName string) time.Time {
	lastEventTime := cache.repository.GetKeyValue("event_"+eventName, "2000-01-01T00:00:00Z00:00")
	lastTime, err := time.Parse(time.RFC3339, lastEventTime)
	if err != nil {
		lastTime = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	return lastTime
}

func (cache *MappaCache) SetLastEventTime(eventName string, lastEventTime time.Time) {
	cache.repository.SetKeyValue("event_"+eventName, lastEventTime.Format(time.RFC3339), time.Duration(0))
}
