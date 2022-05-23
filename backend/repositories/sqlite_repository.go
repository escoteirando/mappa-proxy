package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/entities"
	"github.com/escoteirando/mappa-proxy/backend/infra"
	_ "github.com/mattn/go-sqlite3"
	gosqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteRepository struct {
	BaseRepository
	filename string
}

func init() {
	RepositoryFactory.Register(&SQLiteRepository{})
}

func (r *SQLiteRepository) GetName() string {
	return "SQLite"
}

func (repository *SQLiteRepository) IsValidConnectionString(connectionString string) bool {
	conn, err := infra.CreateConnectionString(connectionString)
	return err == nil && conn.Schema == "sqlite"
}

func (repository *SQLiteRepository) CreateRepository(connectionString string) (IRepository, error) {
	conn, _ := infra.CreateConnectionString(connectionString)
	r := &SQLiteRepository{filename: conn.ConnectionData}

	err := r.setup()
	return r, err
}
func (r *SQLiteRepository) GetLogins() (logins []*domain.LoginData, err error) {
	r.Lock()
	defer r.Unlock()
	var db *sql.DB
	db, err = r.openDb()
	if err != nil {
		return
	}
	defer db.Close()
	var rows *sql.Rows
	rows, err = db.Query("select id,username,login_response,password_hash,last_login from logins")
	if err != nil {
		return
	}
	defer rows.Close()
	logins = make([]*domain.LoginData, 0)
	toDeleteLogins := make([]int, 0)
	for rows.Next() {
		var id int
		var username string
		var login_response string
		var password_hash_s string
		var last_login_str string

		if err = rows.Scan(&id, &username, &login_response, &password_hash_s, &last_login_str); err != nil {
			continue
		}
		var loginResponse responses.MappaLoginResponse
		if err = json.Unmarshal([]byte(login_response), &loginResponse); err != nil {
			toDeleteLogins = append(toDeleteLogins, id)
			continue
		}
		if !loginResponse.IsValid() {
			toDeleteLogins = append(toDeleteLogins, id)
			continue
		}
		last_login, _ := time.Parse(time.RFC3339, last_login_str)
		password_hash, _ := strconv.ParseInt(password_hash_s, 10, 64)
		logins = append(logins, &domain.LoginData{UserName: username,
			LoginResponse: loginResponse,
			PasswordHash:  uint64(password_hash),
			LastLogin:     last_login})
	}

	return logins, err
}
func (r *SQLiteRepository) SetLogin(username string, password string, loginResponse responses.MappaLoginResponse, last_login time.Time) error {
	r.Lock()
	defer r.Unlock()
	loginResponseJson, err := json.Marshal(loginResponse)
	if err != nil {
		return err
	}
	db, err := r.openDb()
	if err != nil {
		return err
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO logins (username,login_response,password_hash,last_login) values (?, ?, ?, ?)",
		username,
		loginResponseJson,
		fmt.Sprintf("%d", infra.GetPasswordHash(password)),
		last_login.Format(time.RFC3339),
	)
	if err == nil {
		rowsAffected, err := result.RowsAffected()
		if err == nil && rowsAffected == 0 {
			err = fmt.Errorf("Zero rows affected")
		}
	}
	return err
}

func (r *SQLiteRepository) DeleteLogin(username string) error {
	r.Lock()
	defer r.Unlock()
	db, err := r.openDb()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM logins WHERE username = ?", username)
	return err
}

func (r *SQLiteRepository) openDb() (*sql.DB, error) {
	return sql.Open("sqlite3", r.filename)
}

func setupTesting(r *SQLiteRepository) {
	db, err := gorm.Open(gosqlite.Open(r.filename), &gorm.Config{})
	if err == nil {
		err = db.AutoMigrate(&entities.MappaEspecialidadeItem{})
		if err == nil {
			db.AutoMigrate(&entities.MappaEspecialidade{})
		}
	}
	if err != nil {
		log.Printf("Error on AutoMigrate: %s", err.Error())
	}
}
func (r *SQLiteRepository) setup() error {
	/*{
	  "id": "tCfxxjJagSFXU15M9lOXLCXMdaHYItdAsNhfSXfvoUj1VrrhIcn8sgRUY4QrMiIe",
	  "ttl": 1209600,
	  "created": "2022-04-25T11:58:24.219Z",
	  "userId": 50442
	}
	*/
	r.Lock()
	defer r.Unlock()

	setupTesting(r)
	db, err := sql.Open("sqlite3", r.filename)
	if err != nil {
		return err
	}
	sqlStmt := `
CREATE TABLE IF NOT EXISTS logins (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	login_response TEXT NOT NULL,
	password_hash INT NOT NULL,
	last_login TEXT NOT NULL,
	mappa_userId INT NOT NULL,
	mappa_auth TEXT NOT NULL,
	mappa_validUntil INT NOT NULL,
	UNIQUE(username) ON CONFLICT REPLACE
);
CREATE INDEX IF NOT EXISTS idx_last_login ON logins (last_login);
CREATE TABLE IF NOT EXISTS detalhes_escotista (
	user_id integer NOT NULL PRIMARY KEY,
	creation_date DATE NOT NULL,
	detalhes TEXT NOT NULL,
	UNIQUE(user_id) ON CONFLICT REPLACE
);
CREATE INDEX IF NOT EXISTS idx_user_id ON detalhes_escotista (user_id);
`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return fmt.Errorf("%q: %s", err, sqlStmt)
	}
	defer db.Close()
	return nil
}

func (r *SQLiteRepository) loadData() error {
	r.Lock()
	defer r.Unlock()
	r.logins = make(map[string]domain.LoginData)
	db, err := sql.Open("sqlite3", r.filename)
	if err != nil {
		return err
	}
	defer db.Close()

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
		var loginResponse responses.MappaLoginResponse
		if err = json.Unmarshal([]byte(login_response), &loginResponse); err != nil {
			return err
		}
		last_login, _ := time.Parse(time.RFC3339, last_login_str)
		r.logins[username] = domain.LoginData{
			UserName:      username,
			LoginResponse: loginResponse,
			PasswordHash:  uint64(password_hash),
			LastLogin:     last_login}
	}

	return nil
}

func (r *SQLiteRepository) SetDetalheEscotista(userId uint, detalheEscotista responses.MappaDetalhesResponse) error {
	r.Lock()
	defer r.Unlock()
	db, err := r.openDb()
	if err != nil {
		return err
	}
	defer db.Close()
	detalheEscotistaJson, err := json.Marshal(detalheEscotista)
	if err != nil {
		return err
	}
	result, err := db.Exec("INSERT INTO detalhes_escotista (user_id,detalhes,creation_date) values (?, ?, ?)",
		userId,
		detalheEscotistaJson,
		time.Now().Format(time.RFC3339),
	)
	if err == nil {
		rowsAffected, err := result.RowsAffected()
		if err == nil && rowsAffected == 0 {
			err = fmt.Errorf("Zero rows affected")
		}
	}
	return err
}

func (r *SQLiteRepository) GetDetalheEscotista(userId uint) (*responses.MappaDetalhesResponse, error) {
	r.Lock()
	defer r.Unlock()
	db, err := r.openDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var detalheEscotistaJson string
	var creationDate time.Time
	err = db.QueryRow("select creation_date,detalhes from detalhes_escotista where user_id = ?", userId).Scan(&creationDate, &detalheEscotistaJson)
	if err != nil {
		return nil, err
	}
	if time.Now().Sub(creationDate) > time.Hour*24 {
		// Mais de 24 horas, limpa o detalhe
		result, err := db.Exec("delete from detalhes_escotista where user_id = ?", userId)
		if err == nil {
			rowsAffected, err := result.RowsAffected()
			if err == nil && rowsAffected == 0 {
				err = fmt.Errorf("Zero rows affected")
			}
		}
		return nil, nil
	}
	var detalheEscotista responses.MappaDetalhesResponse
	if err = json.Unmarshal([]byte(detalheEscotistaJson), &detalheEscotista); err != nil {
		log.Printf("Erro ao deserializar detalhe escotista: %s %s", err, detalheEscotistaJson)
		return nil, err
	}
	return &detalheEscotista, nil
}

func (r *SQLiteRepository) SetKeyValue(key, value string, timeToLive time.Duration) error {
	// TODO: Implement
	return nil
}

func (r *SQLiteRepository) GetKeyValue(key, defaultValue string) string {
	// TODO: IMPLEMENT
	return ""
}

func (r *SQLiteRepository) UpdateMappaProgressoes(progressoes []*responses.MappaProgressaoResponse) error {
	// TODO IMPLEMENT
	return nil
}

func (r *SQLiteRepository) GetProgressoes(ramo domain.MappaRamoEnum) ([]*responses.MappaProgressaoResponse, error) {
	// TODO IMPLEMENT
	return nil, nil
}