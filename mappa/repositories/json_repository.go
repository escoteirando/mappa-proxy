package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/guionardo/mappa_proxy/mappa/domain"
	"github.com/guionardo/mappa_proxy/mappa/tools"
)

type JSONRepository struct {
	Repository
	filename string
}

func CreateJSONRepository(filename string) (*JSONRepository, error) {
	r := &JSONRepository{filename: filename}	
	err := r.loadData()
	return r, err
}

func (r *JSONRepository) SetLogin(username string, password string, loginResponse domain.LoginResponse, last_login time.Time) {
	r.Lock()

	r.logins[username] = domain.LoginData{
		LoginResponse: loginResponse,
		UserName:      username,
		PasswordHash:  tools.GetPasswordHash(password),
	}
	r.Unlock()
	r.SaveData()

	//2021-05-23T11:10:30.950Z
	// var validUntil:=loginResponse.validUntil.
	// var created,err:=time.Parse("2006-01-02T15:04:05.000Z",loginResponse.)
}
func (r *JSONRepository) GetLogin(username string, password string) (loginResponse domain.LoginResponse, err bool) {
	r.loadData()
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
		r.SaveData()
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

func (r *JSONRepository) loadData() error {
	r.Lock()
	defer r.Unlock()

	jsonFile, err := os.Open(r.filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to read file %s", err)
	}

	var fileLogins map[string]domain.LoginData
	err = json.Unmarshal(byteValue, &fileLogins)
	if err != nil {
		return fmt.Errorf("failed to unmarshal file data %s", err)
	}

	r.logins = fileLogins
	return nil
}

func (r *JSONRepository) SaveData() error {
	r.Lock()
	defer r.Unlock()
	jsonData, err := json.Marshal(r.logins)
	if err != nil {
		return fmt.Errorf("failed to serialize logins: %s", err)

	}
	jsonFile, err := os.Create(r.filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s", err)
	}
	defer jsonFile.Close()
	_, err = jsonFile.Write(jsonData)
	return err
}
