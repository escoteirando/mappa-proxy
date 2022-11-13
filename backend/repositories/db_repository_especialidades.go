package repositories

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) UpdateMappaEspecialidades(especialidades []*entities.MappaEspecialidade) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(especialidades)
	return res.Error
}

func (r *DBRepository) GetEspecialidades() (especialidades []*entities.MappaEspecialidade, err error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	if err = db.Find(&especialidades).Error; err != nil {
		return
	}
	for _, especialidade := range especialidades {
		var items []entities.MappaEspecialidadeItem
		if err = db.Where("codigo = ?", especialidade.Codigo).Find(&items).Error; err == nil {
			especialidade.Itens = items
		}
	}

	return
}
