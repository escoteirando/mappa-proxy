package entities

import "github.com/escoteirando/mappa-proxy/backend/domain"

type MappaProgressao struct {
	NoIdModel
	Codigo                int                  `gorm:"column:codigo;primary_key"`
	Descricao             string               `gorm:"column:descricao"`
	CodigoUEB             string               `gorm:"column:codigo_ueb"`
	Ordenacao             int                  `gorm:"column:ordenacao"`
	CodigoCaminho         int                  `gorm:"column:codigo_caminho"`
	CodigoDesenvolvimento int                  `gorm:"column:codigo_desenvolvimento"`
	NumeroGrupo           int                  `gorm:"column:numero_grupo"`
	CodigoRegiao          string               `gorm:"column:codigo_regiao"`
	CodigoCompetencia     int                  `gorm:"column:codigo_competencia"`
	Segmento              string               `gorm:"column:segmento"`
	Ramo                  domain.MappaRamoEnum `gorm:"column:ramo"`
}

const mappaProgressaoTable = "mappa_progressoes"

func init() {
	RegisterEntity(Entity{EntityType: MappaProgressao{}, TableName: mappaProgressaoTable})
}

func (MappaProgressao) TableName() string {
	return mappaProgressaoTable
}
