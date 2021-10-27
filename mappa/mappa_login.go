package mappa

import (
	"bytes"
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/guionardo/mappa_proxy/mappa/domain"
	"github.com/guionardo/mappa_proxy/mappa/repositories"
	"github.com/guionardo/mappa_proxy/mappa/tools"
)

var logins = struct {
	sync.RWMutex
	lastLogin     time.Time
	lastUserLogin string
	logins        map[string]domain.LoginData
}{logins: make(map[string]domain.LoginData), lastLogin: time.Time{}}
var StartedTime = time.Now()

// func SetLogin(username string, password string, loginResponse domain.LoginResponse) {
// 	logins.Lock()
// 	logins.logins[username] = domain.LoginData{
// 		loginResponse,
// 		username,
// 		tools.GetPasswordHash(password),
// 	}
// 	saveLogins()
// 	logins.Unlock()
// 	//2021-05-23T11:10:30.950Z
// 	// var validUntil:=loginResponse.validUntil.
// 	// var created,err:=time.Parse("2006-01-02T15:04:05.000Z",loginResponse.)
// }

func GetLogin(username string, password string) (loginResponse domain.LoginResponse, err bool) {
	repository := repositories.GetRepository()
	loadData()
	logins.RLock()
	login, ok := logins.logins[username]
	logins.RUnlock()
	if !ok {
		return domain.LoginResponse{}, false
	}
	var validUntil = login.LoginResponse.Created.Add(time.Second * time.Duration(login.LoginResponse.TTL))
	if !validUntil.After(time.Now()) {
		log.Printf("Invalidate login from user %s\n", username)
		if repository.DeleteLogin(username) == nil {
			repository.SaveData()
		}

		return domain.LoginResponse{}, false
	}
	if tools.GetPasswordHash(password) != login.PasswordHash {
		log.Printf("Password doesn't matches cached data for user %s\n", username)
		return domain.LoginResponse{}, true
	}

	logins.Lock()
	logins.lastLogin = time.Now()
	logins.lastUserLogin = username
	logins.Unlock()
	return login.LoginResponse, true

}

func PostLogin(username string, password string) (loginResponse domain.LoginResponse, ok bool) {
	loginRequest := &domain.LoginRequest{
		Type:     "LOGIN_REQUEST",
		UserName: username,
		Password: password,
	}
	b, err := json.Marshal(loginRequest)
	if err != nil {
		log.Printf("Failed to serialize login request %s\n", err)
		return domain.LoginResponse{}, false
	}

	url := UrlMappa + "/api/escotistas/login"

	log.Printf("Login request %v: %v", url, username)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create request %s\n", err)
		return domain.LoginResponse{}, false
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("User-Agent", "okhttp/3.4.1")
	resp, err := HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		if resp.StatusCode > 0 && resp.StatusCode < 300 {
			var loginResponse domain.LoginResponse
			json.Unmarshal(body, &loginResponse)
			log.Printf("MAPPA login ok: %s", username)
			return loginResponse, true
		} else {
			log.Printf("Fail on MAPPA login: StatusCode = %d Body = %v", resp.StatusCode, string(body))
		}
	}
	return domain.LoginResponse{}, false
}

var notLoadedYet = true

func loadData() {
	if !notLoadedYet {
		return
	}
	jsonFile, err := os.Open("./mappa_login.json")
	if err != nil {
		log.Printf("Failed to open file: %s", err)
		return
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Printf("Failed to read file %s\n", err)
		return
	}
	var fileLogins map[string]domain.LoginData
	err = json.Unmarshal(byteValue, &fileLogins)
	if err != nil {
		log.Printf("Failed to unmarshal file data %s\n", err)
		return
	}
	logins.Lock()
	logins.logins = fileLogins
	logins.Unlock()
	notLoadedYet = false
}

// func saveLogins() {
// 	jsonData, err := json.Marshal(logins.logins)
// 	if err != nil {
// 		log.Printf("Failed to serialize logins: %s", err)
// 		return
// 	}
// 	jsonFile, err := os.Create("./mappa_login.json")
// 	if err != nil {
// 		log.Printf("Failed to create file %s", err)
// 		return
// 	}
// 	defer jsonFile.Close()
// 	jsonFile.Write(jsonData)
// }

func StartMappa() {
	loadData()
}
