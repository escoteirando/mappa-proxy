package repositories

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) SetSecao(secao *entities.Secao) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(secao)
	return res.Error
}
func (r *DBRepository) GetSecao(codigoSecao int, codigoRegiao string) (secao *entities.Secao, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	res := db.First(&secao, codigoSecao, codigoRegiao)
	err = res.Error
	return
}
