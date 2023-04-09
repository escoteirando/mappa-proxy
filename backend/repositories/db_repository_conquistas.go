package repositories

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) UpdateMappaConquistas(conquistas []*entities.MappaConquista) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(conquistas)
	return res.Error
}

func (r *DBRepository) GetConquistas(codigoSecao int, since time.Time) (conquistas []*entities.MappaConquista, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	err = db.Where("data>=?", since).Find(&conquistas).Error
	return
}
