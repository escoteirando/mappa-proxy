package repositories

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/entities"
)

type IRepository interface {
	IsValidConnectionString(connectionString string) bool
	CreateRepository(connectionString string) (IRepository, error)
	GetName() string

	GetLogins() (logins []*domain.LoginData, err error) //

	SetLogin(username string, password string, loginResponse responses.MappaLoginResponse, last_login time.Time) error //
	// GetLogin(username string, password string) (loginResponse *responses.MappaLoginResponse, err error)
	DeleteLogin(username string) error //
	// SaveData() error
	// loadData() error

	// GetLastLogin() (username string, timestamp time.Time)
	// SetLastLogin(username string, timestamp time.Time) error
	// GetUserCount() int

	SetEscotista(escotista *entities.Escotista) error
	GetEscotista(userId int) (escotista *entities.Escotista, err error)

	SetAssociado(associado *entities.Associado) error
	GetAssociado(codigoAssociado int) (associado *entities.Associado, err error)

	SetGrupo(grupo *entities.Grupo) error
	GetGrupo(codigoGrupo int, codigoRegiao string) (grupo *entities.Grupo, err error)

	SetDetalheEscotista(userId uint, detalheEscotista responses.MappaDetalhesResponse) error
	GetDetalheEscotista(userId uint) (*responses.MappaDetalhesResponse, error)

	SetKeyValue(key, value string, timeToLive time.Duration) error
	GetKeyValue(key, defaultValue string) string

	UpdateMappaProgressoes(progressoesResponse []*responses.MappaProgressaoResponse) error
	GetProgressoes(domain.MappaRamoEnum) ([]*responses.MappaProgressaoResponse, error)
}
