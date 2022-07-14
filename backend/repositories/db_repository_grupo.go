package repositories

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) SetGrupo(grupo *entities.Grupo) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(grupo)
	return res.Error
}
func (r *DBRepository) GetGrupo(codigoGrupo int) (grupo *entities.Grupo, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.First(&grupo, codigoGrupo)
	err = res.Error
	return
}
