package repositories

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"gorm.io/gorm/clause"
)

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
