package mappa

import (
	"bytes"
	"encoding/json"
	"hash/crc64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type LoginResponse struct {
	ID      string    `json:"id"`
	TTL     int       `json:"ttl"`
	Created time.Time `json:"created"`
	Userid  int       `json:"userId"`
}

type LoginData struct {
	LoginResponse LoginResponse
	UserName      string
	PasswordHash  uint64
}

// {"type":"LOGIN_REQUEST","username":"guionardo","password":"A1GU"}
type LoginRequest struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Stats struct {
	RunningSince  time.Time `json:"running_since"`
	Users         int       `json:"users"`
	LastLogin     time.Time `json:"last_login"`
	LastUserLogin string    `json:"last_user_login"`
}

var logins = struct {
	sync.RWMutex
	lastLogin     time.Time
	lastUserLogin string
	logins        map[string]LoginData
}{logins: make(map[string]LoginData), lastLogin: time.Time{}}
var startedTime = time.Now()

func SetLogin(username string, password string, loginResponse LoginResponse) {
	logins.Lock()
	logins.logins[username] = LoginData{
		loginResponse,
		username,
		getPasswordHash(password),
	}
	saveLogins()
	logins.Unlock()
	//2021-05-23T11:10:30.950Z
	// var validUntil:=loginResponse.validUntil.
	// var created,err:=time.Parse("2006-01-02T15:04:05.000Z",loginResponse.)
}

func GetLogin(username string, password string) (loginResponse LoginResponse, err bool) {
	loadData()
	logins.RLock()
	login, ok := logins.logins[username]
	logins.RUnlock()
	if !ok {
		return LoginResponse{}, false
	}
	var validUntil = login.LoginResponse.Created.Add(time.Second * time.Duration(login.LoginResponse.TTL))
	if !validUntil.After(time.Now()) {
		log.Printf("Invalidate login from user %s\n", username)
		logins.Lock()
		delete(logins.logins, username)
		saveLogins()
		logins.Unlock()
		return LoginResponse{}, false
	}
	if getPasswordHash(password) != login.PasswordHash {
		log.Printf("Password doesn't matches cached data for user %s\n", username)
		return LoginResponse{}, true
	}

	logins.Lock()
	logins.lastLogin = time.Now()
	logins.lastUserLogin = username
	logins.Unlock()
	return login.LoginResponse, true

}

func PostLogin(username string, password string) (loginResponse LoginResponse, ok bool) {
	loginRequest := &LoginRequest{
		Type:     "LOGIN_REQUEST",
		UserName: username,
		Password: password,
	}
	b, err := json.Marshal(loginRequest)
	if err != nil {
		log.Printf("Failed to serialize login request %s\n", err)
		return LoginResponse{}, false
	}

	url := UrlMappa + "/api/escotistas/login"

	log.Printf("Login request %v: %v", url, username)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to create request %s\n", err)
		return LoginResponse{}, false
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
			var loginResponse LoginResponse
			json.Unmarshal(body, &loginResponse)
			log.Printf("MAPPA login ok: %s", username)
			return loginResponse, true
		} else {
			log.Printf("Fail on MAPPA login: StatusCode = %d Body = %v", resp.StatusCode, string(body))
		}
	}
	return LoginResponse{}, false
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
	var fileLogins map[string]LoginData
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

func saveLogins() {
	jsonData, err := json.Marshal(logins.logins)
	if err != nil {
		log.Printf("Failed to serialize logins: %s", err)
		return
	}
	jsonFile, err := os.Create("./mappa_login.json")
	if err != nil {
		log.Printf("Failed to create file %s", err)
		return
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
}

var crc64Table = crc64.MakeTable(crc64.ECMA)

func getPasswordHash(password string) uint64 {
	return crc64.Checksum([]byte(password), crc64Table)
}

func StartMappa() {
	loadData()
}

func GetStats() Stats {
	logins.RLock()
	usersCount := len(logins.logins)
	lastLogin := logins.lastLogin
	lastUserLogin := logins.lastUserLogin
	logins.RUnlock()
	return Stats{
		RunningSince:  startedTime,
		Users:         usersCount,
		LastLogin:     lastLogin,
		LastUserLogin: lastUserLogin,
	}
}
