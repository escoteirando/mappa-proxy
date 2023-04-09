package repositories

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) SetSubSecao(subsecao *entities.SubSecao) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(subsecao)
	return res.Error
}
func (r *DBRepository) GetSubSecao(codigoSubSecao int) (subsecao *entities.SubSecao, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	res := db.First(&subsecao, codigoSubSecao)
	err = res.Error
	return
}

func (r *DBRepository) GetSubSecoes(codigoSecao int) (subsecoes []*entities.SubSecao, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	res := db.Where("codigo_secao = ?", codigoSecao).Find(&subsecoes)
	err = res.Error
	return
}

func (r *DBRepository) GetSubSecaoAssociados(codigoSubSecao int) (associados []entities.Associado, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()

	res := db.Where("codigo_equipe = ?", codigoSubSecao).Find(&associados)
	err = res.Error
	return
}
