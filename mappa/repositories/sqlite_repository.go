package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/guionardo/mappa_proxy/mappa/domain"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	Repository
	filename string
}

func CreateSQLiteRepository(filename string) (*SQLiteRepository, error) {
	r := &SQLiteRepository{filename: filename}

	err := r.setup()
	if err == nil {
		err = r.loadData()
	}
	return r, err
}

func (r *SQLiteRepository) setup() error {
	r.Lock()
	defer r.Unlock()
	db, err := sql.Open("sqlite3", r.filename)
	if err != nil {
		return err
	}
	sqlStmt := `
CREATE TABLE IF NOT EXISTS logins (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	login_response TEXT NOT NULL,
	password_hash TEXT NOT NULL,
	last_login TEXT NOT NULL,
	UNIQUE(username) ON CONFLICT REPLACE
);
CREATE INDEX IF NOT EXISTS idx_last_login ON logins (last_login);
`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return fmt.Errorf("%q: %s", err, sqlStmt)
	}
	defer db.Close()
	return nil
}

// {
//     "guionardo": {
//         "LoginResponse": {
//             "id": "G1BxTHTnwuSNZfJkWKaIgSz19cnAEE3EZAt9tavibzMySXLwoQr5MgBUEBXOijUr",
//             "ttl": 1209600,
//             "created": "2021-10-26T11:15:19.739Z",
//             "userId": 50442
//         },
//         "UserName": "guionardo",
//         "PasswordHash": 17660074743493757612
//     }
// }
func (r *SQLiteRepository) loadData() error {
	r.Lock()
	defer r.Unlock()
	r.logins = make(map[string]domain.LoginData)
	db, err := sql.Open("sqlite3", r.filename)
	if err != nil {
		return err
	}
	defer db.Close()
	//TODO: Implementar

	rows, err := db.Query("select id,username,login_response,password_hash,last_login from logins")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		var login_response string
		var password_hash int64
		var last_login_str string

		if err = rows.Scan(&id, &username, &login_response, &password_hash, &last_login_str); err != nil {
			return err
		}
		var loginResponse domain.LoginResponse
		if err = json.Unmarshal([]byte(login_response), &loginResponse); err != nil {
			return err
		}
		last_login, _ := time.Parse(time.RFC3339, last_login_str)
		r.logins[username] = domain.LoginData{UserName: username,
			LoginResponse: loginResponse,
			PasswordHash:  uint64(password_hash),
			LastLogin:     last_login}
	}

	return nil
}

// func (r *SQLiteRepository) SetLogin(username string, password string, loginResponse domain.LoginResponse) {
// 	r.Lock()

// 	r.logins[username] = domain.LoginData{
// 		LoginResponse: loginResponse,
// 		UserName:      username,
// 		PasswordHash:  tools.GetPasswordHash(password),
// 	}
// 	r.Unlock()
// 	r.saveLogins()
// }
// func (r *SQLiteRepository) GetLogin(username string, password string) (loginResponse domain.LoginResponse, err bool) {
// 	r.loadData()
// 	r.RLock()
// 	login, ok := r.logins[username]
// 	r.RUnlock()
// 	if !ok {
// 		return domain.LoginResponse{}, false
// 	}
// 	var validUntil = login.LoginResponse.Created.Add(time.Second * time.Duration(login.LoginResponse.TTL))
// 	if !validUntil.After(time.Now()) {
// 		log.Printf("Invalidate login from user %s\n", username)
// 		r.Lock()
// 		delete(r.logins, username)
// 		r.Unlock()
// 		r.saveLogins()
// 		return domain.LoginResponse{}, false
// 	}
// 	if tools.GetPasswordHash(password) != login.PasswordHash {
// 		log.Printf("Password doesn't matches cached data for user %s\n", username)
// 		return domain.LoginResponse{}, true
// 	}

// 	r.Lock()
// 	r.lastLogin = time.Now()
// 	r.lastUserLogin = username
// 	r.Unlock()
// 	return login.LoginResponse, true
// }

func (r *SQLiteRepository) SaveData() error {
	r.Lock()
	defer r.Unlock()

	db, err := sql.Open("sqlite3", r.filename)
	if err != nil {
		return err
	}
	defer db.Close()

	for _, logins := range r.logins {
		loginResponse, err := json.Marshal(logins)
		if err != nil {
			return err
		}

		result, err := db.Exec("insert into logins (username,login_response,password_hash,last_login) values ($1,$2,$3,$4)",
			logins.UserName,
			string(loginResponse),
			int64(logins.PasswordHash),
			logins.LastLogin.Format(time.RFC3339))
		if err != nil {
			return err
		}
		if rows, err := result.RowsAffected(); rows == 0 || err != nil {
			return fmt.Errorf("insert failed (rows=%d, err=%v)", rows, err)
		}

	}
	return nil
}
