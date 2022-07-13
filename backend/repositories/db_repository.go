package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/entities"
	"github.com/escoteirando/mappa-proxy/backend/infra"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	DBRepository struct {
		BaseRepository
		schema           string
		connectionString string
		getDBFunc        func() *gorm.DB
	}
)

func init() {
	RepositoryFactory.Register(&DBRepository{})
}

func (r *DBRepository) IsValidConnectionString(connectionString string) bool {
	cs, err := infra.CreateConnectionString(connectionString)	
	return err == nil && (cs.Schema == "sqlite" || cs.Schema == "postgres")
}

func (repository *DBRepository) CreateRepository(connectionString string) (IRepository, error) {
	conn, _ := infra.CreateConnectionString(connectionString)
	r := &DBRepository{
		schema:           conn.Schema,
		connectionString: connectionString,		
	}
	r.SetLocking(conn.Schema == "sqlite")
	switch conn.Schema {
	case "sqlite":
		r.getDBFunc = func() *gorm.DB {
			db, err := gorm.Open(sqlite.Open(conn.ConnectionData), &gorm.Config{})
			if err != nil {
				log.Fatal(err)
			}
			return db
		}
	case "postgres":
		r.getDBFunc = func() *gorm.DB {
			db, err := gorm.Open(postgres.Open(conn.ConnectionData), &gorm.Config{})
			if err != nil {
				log.Fatal(err)
			}
			return db
		}
	}

	err := r.setup()
	if err != nil {
		r = nil
	}
	return r, err
}

func (r *DBRepository) setup() error {
	db := r.getDBFunc()
	return db.AutoMigrate(
		&entities.Login{},
		&entities.Escotista{},
		&entities.Associado{},
		&entities.Grupo{},
		&entities.DetalhesEscotista{},
		&entities.MappaEspecialidadeItem{},
		&entities.MappaEspecialidadeRequisito{},
		&entities.MappaEspecialidade{},
		&entities.MappaProgressao{},
		&entities.KeyValue{},
	)
}

func (r *DBRepository) GetName() string {
	return fmt.Sprintf("DB: %s", r.schema)
}

func (r *DBRepository) GetLogins() (logins []*domain.LoginData, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	var entityLogins []entities.Login
	res := db.Find(&entityLogins)
	if res.Error != nil {
		return nil, res.Error
	}
	logins = make([]*domain.LoginData, 0)
	toDeleteLogins := make([]int, 0)
	for _, login := range entityLogins {
		logins = append(logins, &domain.LoginData{
			UserName:      login.UserName,
			UserId:        login.MappaUserId,
			PasswordHash:  login.PasswordHash,
			LastLogin:     login.LastLogin,
			Authorization: login.MappaAuth,
			ValidUntil:    login.MappaValidUntil,
		})
	}
	if len(toDeleteLogins) > 0 {
		res := db.Delete(&entities.Login{}, "mappa_user_id IN (?)", toDeleteLogins)

		if res.Error != nil {
			log.Printf("Error deleting invalid logins: %s", res.Error)
		} else {
			log.Printf("Deleted %d invalid logins", res.RowsAffected)
		}
	}
	return logins, err
}

func (r *DBRepository) SetLogin(username string, password string, loginResponse responses.MappaLoginResponse, last_login time.Time) error {
	r.DBLock()
	defer r.DBUnlock()

	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entities.Login{
		UserName:        username,
		PasswordHash:    infra.GetPasswordHash(password),
		LastLogin:       last_login,
		MappaUserId:     loginResponse.Userid,
		MappaAuth:       loginResponse.ID,
		MappaValidUntil: loginResponse.ValidUntil(),
	})

	return res.Error
}

func (r *DBRepository) DeleteLogin(username string) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	deletedLogin := entities.Login{
		UserName: username,
	}
	res := db.Delete(&deletedLogin)
	return res.Error
}
func (r *DBRepository) SetDetalheEscotista(userId uint, detalheEscotista responses.MappaDetalhesResponse) error {
	detalheEscotistaJson, err := json.Marshal(detalheEscotista)
	if err != nil {
		return err
	}
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entities.DetalhesEscotista{
		UserId:   userId,
		Detalhes: string(detalheEscotistaJson),
	})

	return res.Error
}
func (r *DBRepository) GetDetalheEscotista(userId uint) (*responses.MappaDetalhesResponse, error) {
	r.DBLock()
	defer r.DBUnlock()
	detalheEscotista := entities.DetalhesEscotista{
		UserId: userId,
	}
	db := r.getDBFunc()
	res := db.First(&detalheEscotista, userId)
	if res.Error != nil {
		return nil, res.Error
	}
	var detalhe responses.MappaDetalhesResponse
	if err := json.Unmarshal([]byte(detalheEscotista.Detalhes), &detalhe); err != nil {
		log.Printf("Erro ao deserializar detalhe escotista: %s %s", err, detalheEscotista.Detalhes)
		return nil, err
	}
	return &detalhe, nil
}

func GetDatabase(dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func GetDatabaseFromConnectionString(connectionString string) (*gorm.DB, error) {
	cs, err := infra.CreateConnectionString(connectionString)
	if err != nil {
		return nil, err
	}
	var dialector gorm.Dialector
	switch cs.Schema {
	case "sqlite":
		dialector = sqlite.Open(cs.ConnectionData)
	case "postgres":
		dialector = postgres.Open(cs.ConnectionData)
	default:
		return nil, fmt.Errorf("Unexpected database schema: %s", connectionString)
	}
	db, err := gorm.Open(dialector, &gorm.Config{})
	return db, err
}

func (r *DBRepository) SetKeyValue(key, value string, timeToLive time.Duration) error {
	r.DBLock()
	defer r.DBUnlock()
	var validUntil time.Time
	if timeToLive == time.Duration(0) {
		validUntil = time.Date(2099, 12, 31, 23, 59, 59, 0, time.UTC)
	} else {
		validUntil = time.Now().Add(timeToLive)
	}

	keyValue := entities.KeyValue{
		Key:        key,
		Value:      value,
		ValidUntil: validUntil,
	}
	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&keyValue)
	return res.Error
}

func (r *DBRepository) GetKeyValue(key, defaultValue string) string {
	r.DBLock()
	defer r.DBUnlock()
	keyValue := entities.KeyValue{
		Key: key,
	}
	db := r.getDBFunc()
	res := db.First(&keyValue)
	if res.Error != nil {
		return defaultValue
	}
	if keyValue.ValidUntil.Before(time.Now()) {
		db.Delete(&keyValue)
		return defaultValue
	}
	return keyValue.Value
}

func (r *DBRepository) SetEscotista(escotista *entities.Escotista) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(escotista)
	return res.Error
}

func (r *DBRepository) GetEscotista(userId int) (escotista *entities.Escotista, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.First(&escotista, userId)
	err = res.Error
	return
}
func (r *DBRepository) SetAssociado(associado *entities.Associado) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(associado)
	return res.Error
}

func (r *DBRepository) GetAssociado(codigoAssociado int) (associado *entities.Associado, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.First(&associado, codigoAssociado)
	err = res.Error
	return
}

func (r *DBRepository) SetGrupo(grupo *entities.Grupo) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(grupo)
	return res.Error
}
func (r *DBRepository) GetGrupo(codigoGrupo int, codigoRegiao string) (grupo *entities.Grupo, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.First(&grupo, codigoGrupo, codigoRegiao)
	err = res.Error
	return
}
