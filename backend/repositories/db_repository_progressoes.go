package repositories

import (
	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/entities"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) UpdateMappaProgressoes(progressoes []*responses.MappaProgressaoResponse) error {
	r.DBLock()
	defer r.DBUnlock()
	db := r.getDBFunc()
	// Desabilitar todas as progressões atuais
	db.Delete(&entities.MappaProgressao{}, "deleted_at IS NULL")
	for _, progressao := range progressoes {
		res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entities.MappaProgressao{
			Codigo:                progressao.Codigo,
			Descricao:             progressao.Descricao,
			CodigoUEB:             progressao.CodigoUEB,
			Ordenacao:             progressao.Ordenacao,
			CodigoCaminho:         progressao.CodigoCaminho,
			CodigoCompetencia:     progressao.CodigoCompetencia,
			CodigoDesenvolvimento: progressao.CodigoDesenvolvimento,
			Segmento:              progressao.Segmento,
			NumeroGrupo:           progressao.NumeroGrupo,
			CodigoRegiao:          progressao.CodigoRegiao,
			Ramo:                  domain.CaminhoToRamo(progressao.CodigoCaminho),
		})
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
		rsp[index] = &responses.MappaProgressaoResponse{
			Codigo:                progresso.Codigo,
			Descricao:             progresso.Descricao,
			CodigoUEB:             progresso.CodigoUEB,
			Ordenacao:             progresso.Ordenacao,
			CodigoCaminho:         progresso.CodigoCaminho,
			CodigoCompetencia:     progresso.CodigoCompetencia,
			CodigoDesenvolvimento: progresso.CodigoDesenvolvimento,
			Segmento:              progresso.Segmento,
			NumeroGrupo:           progresso.NumeroGrupo,
			CodigoRegiao:          progresso.CodigoRegiao,
		}
	}
	return rsp, nil
}
