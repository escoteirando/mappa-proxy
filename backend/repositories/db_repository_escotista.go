package repositories

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"gorm.io/gorm/clause"
)

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
