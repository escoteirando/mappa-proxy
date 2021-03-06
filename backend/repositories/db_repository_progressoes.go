package repositories

import (
	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/dtos"
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) UpdateMappaProgressoes(progressoes []*responses.MappaProgressaoResponse) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	// Desabilitar todas as progressões atuais
	db.Delete(&entities.MappaProgressao{}, "deleted_at IS NULL")
	for _, progressao := range progressoes {
		res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(dtos.MappaProgressaoToEntity(progressao))
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func (r *DBRepository) GetProgressoes(ramo domain.MappaRamoEnum) ([]*responses.MappaProgressaoResponse, error) {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	var progressoes []*entities.MappaProgressao
	res := db.Order("codigo_caminho ASC, ordenacao ASC").Where("codigo_caminho in (?) and numero_grupo = 0", ramo.Caminhos()).Find(&progressoes)
	if res.Error != nil {
		return nil, res.Error
	}
	rsp := make([]*responses.MappaProgressaoResponse, len(progressoes))
	for index, progresso := range progressoes {
		rsp[index] = dtos.MappaProgressaoToResponse(progresso)
	}
	return rsp, nil
}
