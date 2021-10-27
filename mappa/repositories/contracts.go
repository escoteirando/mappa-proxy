package repositories

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/guionardo/mappa_proxy/mappa/domain"
	"github.com/guionardo/mappa_proxy/mappa/tools"
)

type Repository struct {
	sync.RWMutex
	lastLogin     time.Time
	lastUserLogin string
	logins        map[string]domain.LoginData
}

type Repositorie interface {
	SetLogin(username string, password string, loginResponse domain.LoginResponse, last_login time.Time)
	GetLogin(username string, password string) (loginResponse domain.LoginResponse, err bool)
	DeleteLogin(username string) error
	SaveData() error
	loadData() error

	GetLastLogin() (username string, timestamp time.Time)
	SetLastLogin(username string, timestamp time.Time) error
	GetUserCount() int
}

var defaultRepository Repositorie

func SetRepository(connection_string string) error {
	if !strings.Contains(connection_string, ":") {
		return fmt.Errorf("invalid repository '%s' (missing schema JSON or SQLITE)", connection_string)
	}
	words := strings.SplitN(connection_string, ":", 2)
	var err error
	switch words[0] {
	case "JSON":
		defaultRepository, err = CreateJSONRepository(words[1])

	case "SQLITE":
		defaultRepository, err = CreateSQLiteRepository(words[1])
	default:
		err = fmt.Errorf("invalid schema. Expected JSON or SQLITE")
	}
	return err
}

func GetRepository() Repositorie {
	return defaultRepository
}

func (r *Repository) SetLogin(username string, password string, loginResponse domain.LoginResponse, last_login time.Time) {
	r.Lock()

	r.logins[username] = domain.LoginData{
		LoginResponse: loginResponse,
		UserName:      username,
		PasswordHash:  tools.GetPasswordHash(password),
		LastLogin:     last_login,
	}
	r.Unlock()
}

func (r *Repository) GetLogin(username string, password string) (loginResponse domain.LoginResponse, err bool) {
	r.RLock()
	login, ok := r.logins[username]
	r.RUnlock()
	if !ok {
		return domain.LoginResponse{}, false
	}
	var validUntil = login.LoginResponse.Created.Add(time.Second * time.Duration(login.LoginResponse.TTL))
	if !validUntil.After(time.Now()) {
		log.Printf("Invalidate login from user %s\n", username)
		r.Lock()
		delete(r.logins, username)
		r.Unlock()
		return domain.LoginResponse{}, false
	}
	if tools.GetPasswordHash(password) != login.PasswordHash {
		log.Printf("Password doesn't matches cached data for user %s\n", username)
		return domain.LoginResponse{}, true
	}

	r.Lock()
	r.lastLogin = time.Now()
	r.lastUserLogin = username
	r.Unlock()
	return login.LoginResponse, true
}

func (r *Repository) GetLastLogin() (username string, timestamp time.Time) {
	r.RLock()
	defer r.RUnlock()
	lastLogin := time.UnixMilli(0)
	lastUser := ""
	for user, login := range r.logins {
		if login.LastLogin.After(lastLogin) {
			lastLogin = login.LastLogin
			lastUser = user
		}
	}
	return lastUser, lastLogin
}

func (r *Repository) SetLastLogin(username string, timestamp time.Time) error {
	r.Lock()
	defer r.Unlock()
	login, ok := r.logins[username]
	if !ok {
		return fmt.Errorf("user not found %s", username)
	}
	login.LastLogin = timestamp
	r.logins[username] = login
	return nil
}

func (r *Repository) GetUserCount() int {
	r.RLock()
	defer r.RUnlock()
	return len(r.logins)
}

func (r *Repository) DeleteLogin(username string) error {
	r.Lock()
	defer r.Unlock()
	_, ok := r.logins[username]
	if !ok {
		return fmt.Errorf("unexistent user %s", username)
	}
	delete(r.logins, username)
	return nil
}
