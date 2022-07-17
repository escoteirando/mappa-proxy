package repositories

import (
	"fmt"
	"log"

	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/infra"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
		&entities.SubSecao{},
		&entities.Grupo{},
		&entities.Secao{},
		&entities.DetalhesEscotista{},
		&entities.MappaEspecialidadeItem{},
		&entities.MappaEspecialidadeRequisito{},
		&entities.MappaEspecialidade{},
		&entities.MappaProgressao{},
		&entities.KeyValue{},
		&entities.AssociadoSecao{},
	)
}

func (r *DBRepository) GetName() string {
	return fmt.Sprintf("DB: %s", r.schema)
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
