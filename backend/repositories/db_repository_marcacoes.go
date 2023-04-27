package repositories

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) UpdateMappaMarcacoes(marcacoes []*entities.MappaMarcacao) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(marcacoes, 20)
	return res.Error
}

func (r *DBRepository) GetMarcacoes(codigoSecao int) (marcacoes []*entities.MappaMarcacao, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()
	res := db.Where("codigo_secao = ?", codigoSecao).Find(&marcacoes)
	return marcacoes, res.Error
}
