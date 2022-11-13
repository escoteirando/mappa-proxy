package repositories

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) UpdateMappaConquistas(conquistas []*entities.MappaConquista) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(conquistas)
	return res.Error
}

func (r *DBRepository) GetConquistas(codigoSecao int) (conquistas []*entities.MappaConquista, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	err = db.Find(&conquistas).Error
	return
}
