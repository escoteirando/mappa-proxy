package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/danwakefield/fnmatch"
	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/entities"
	"github.com/escoteirando/mappa-proxy/backend/infra"
)

type JsonRepository struct {
	BaseRepository
	repositoryFolder string
}

func init() {
	RepositoryFactory.Register(&JsonRepository{})
}

func (r *JsonRepository) GetName() string {
	return "JSON"
}

func (repository *JsonRepository) IsValidConnectionString(connectionString string) bool {
	conn, err := infra.CreateConnectionString(connectionString)
	return err == nil && conn.Schema == "json"
}

func (repository *JsonRepository) CreateRepository(connectionString string) (IRepository, error) {
	conn, _ := infra.CreateConnectionString(connectionString)
	repositoryFolder, err := filepath.Abs(conn.ConnectionData)
	if err != nil {
		return nil, err
	}
	stat, err := os.Stat(repositoryFolder)
	if err != nil {
		err = os.MkdirAll(repositoryFolder, 0666)
	} else {
		if !stat.IsDir() {
			err = fmt.Errorf("Repository folder is a file: %s", repositoryFolder)
		}
	}

	repo := &JsonRepository{
		repositoryFolder: repositoryFolder,
	}
	return repo, nil
}

func loadLoginData(filename string) *domain.LoginData {
	if stat, err := os.Stat(filename); err != nil || stat.IsDir() {
		return nil
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read %s - %v", filename, err)
		return nil
	}
	var loginData *domain.LoginData
	if err = json.Unmarshal(content, &loginData); err != nil {
		log.Printf("Failed to parse json %s - %v", filename, err)
		return nil
	}
	if !loginData.IsValid() {
		log.Printf("Removed invalidated login data %s", filename)
		os.Remove(filename)
		return nil
	}
	return loginData
}

func (r *JsonRepository) GetLogins() (logins []*domain.LoginData, err error) {
	r.Lock()
	defer r.Unlock()

	// Get all logins files
	loginsFiles, err := ioutil.ReadDir(r.repositoryFolder)
	logins = make([]*domain.LoginData, 0)
	if err != nil {
		return nil, err
	}
	for _, file := range loginsFiles {
		if file.IsDir() || !fnmatch.Match("login_*.json", file.Name(), fnmatch.FNM_FILE_NAME) {
			continue
		}
		loginData := loadLoginData(file.Name())
		if loginData != nil {
			logins = append(logins, loginData)
		}
	}
	return logins, err
}

func loginFile(repositoryFolder string, username string) (fileName string, exists bool) {
	fileName = path.Join(repositoryFolder, fmt.Sprintf("login_%s.json", username))
	stat, err := os.Stat(fileName)
	exists = err != nil || stat.IsDir()
	return
}

func detalheFile(repositoryFolder string, userId int) (fileName string, exists bool) {
	fileName = path.Join(repositoryFolder, fmt.Sprintf("detalhe_%v.json", userId))
	stat, err := os.Stat(fileName)
	exists = err != nil || stat.IsDir()
	return
}

func (r *JsonRepository) SetLogin(username string, password string, loginResponse responses.MappaLoginResponse, last_login time.Time) error {
	r.Lock()
	defer r.Unlock()
	loginFile, _ := loginFile(r.repositoryFolder, username)
	loginData := domain.LoginData{
		UserName:     username,
		PasswordHash: infra.GetPasswordHash(password),
		LastLogin:    last_login,
		Deleted:      false,
	}

	bytes, err := json.Marshal(loginData)
	if err != nil {
		return err
	}
	return os.WriteFile(loginFile, bytes, 0666)
}

func (r *JsonRepository) DeleteLogin(username string) error {
	r.Lock()
	defer r.Unlock()
	loginFile, exists := loginFile(r.repositoryFolder, username)
	if !exists {
		return nil
	}
	return os.Remove(loginFile)
}

func (r *JsonRepository) SetDetalheEscotista(userId uint, detalheEscotista responses.MappaDetalhesResponse) error {
	r.Lock()
	defer r.Unlock()
	detalheEscotistaJson, err := json.Marshal(detalheEscotista)
	if err != nil {
		return err
	}
	detalheFile, _ := detalheFile(r.repositoryFolder, int(userId))
	err = os.WriteFile(detalheFile, detalheEscotistaJson, 0666)

	return err
}

func (r *JsonRepository) GetDetalheEscotista(userId uint) (*responses.MappaDetalhesResponse, error) {
	r.Lock()
	defer r.Unlock()
	detalheFile, exists := detalheFile(r.repositoryFolder, int(userId))
	if !exists {
		return nil, fmt.Errorf("Detalhe escotista não encontrado")
	}
	if stat, err := os.Stat(detalheFile); stat.ModTime().Before(time.Now().Add(-24*time.Hour)) || err != nil {
		return nil, fmt.Errorf("Detalhe escotista expirado")
	}
	content, err := os.ReadFile(detalheFile)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler detalhe escotista - %v", err)
	}
	var detalheEscotista *responses.MappaDetalhesResponse
	if err = json.Unmarshal(content, detalheEscotista); err != nil {
		return nil, err
	}
	return detalheEscotista, nil

}

func (r *JsonRepository) SetKeyValue(key, value string, timeToLive time.Duration) error {
	// TODO: Implement
	return nil
}

func (r *JsonRepository) GetKeyValue(key, defaultValue string) string {
	// TODO: IMPLEMENT
	return ""
}

func (r *JsonRepository) UpdateMappaProgressoes(progressoes []*responses.MappaProgressaoResponse) error {
	// TODO IMPLEMENT
	return nil
}

func (r *JsonRepository) GetProgressoes(ramo domain.MappaRamoEnum) ([]*responses.MappaProgressaoResponse, error) {
	// TODO IMPLEMENT
	return nil, nil
}

func (r *JsonRepository) SetEscotista(escotista *entities.Escotista) error {
	return nil
}
func (r *JsonRepository) GetEscotista(userId int) (escotista *entities.Escotista, err error) {
	return
}

func (r *JsonRepository) SetAssociado(associado *entities.Associado) error {
	return nil
}

func (r *JsonRepository) GetAssociado(codigoAssociado int) (associado *entities.Associado, err error) {
	return
}
func (r *JsonRepository) SetGrupo(grupo *entities.Grupo) error {
	return nil
}
func (r *JsonRepository) GetGrupo(codigoGrupo int, codigoRegiao string) (grupo *entities.Grupo, err error) {
	return
}
