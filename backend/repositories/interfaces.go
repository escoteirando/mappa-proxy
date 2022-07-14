package repositories

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
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
	GetGrupo(codigoGrupo int) (grupo *entities.Grupo, err error)

	SetDetalheEscotista(userId uint, detalheEscotista responses.MappaDetalhesResponse) error
	GetDetalheEscotista(userId uint) (*responses.MappaDetalhesResponse, error)

	SetSecao(secao *entities.Secao) error
	GetSecao(codigoSecao int, codigoRegiao string) (secao *entities.Secao, err error)
	
	SetKeyValue(key, value string, timeToLive time.Duration) error
	GetKeyValue(key, defaultValue string) string

	UpdateMappaProgressoes(progressoesResponse []*responses.MappaProgressaoResponse) error
	GetProgressoes(domain.MappaRamoEnum) ([]*responses.MappaProgressaoResponse, error)
}
