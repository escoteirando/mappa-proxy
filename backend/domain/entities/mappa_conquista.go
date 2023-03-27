package entities

import (
	"time"

	"gorm.io/gorm"
)

type MappaConquista struct {
	gorm.Model
	Type             string             `gorm:"column:type;unique_index:idx_conquista"`
	CodAssociado     int                `gorm:"field:cod_associado;unique_index:idx_conquista"`
	CodEscotista     int                `gorm:"field:cod_escotista"`
	NumeroNivel      int                `gorm:"field:numero_nivel;unique_index:idx_conquista"`
	CodEspecialidade int                `gorm:"field:cod_especialidade;unique_index:idx_conquista"`
	CodigoSecao      int                `gorm:"field:cod_secao:unique_index:idx_marcacao"`
	Associado        Associado          `gorm:"foreignKey:codigo;association_foreignKey:cod_associado"`
	Secao            Secao              `gorm:"foreignKey:codigo;association_foreignKey:cod_secao"`
	Especialidade    MappaEspecialidade `gorm:"foreignKey:codigo;association_foreignKey:cod_especialidade"`
	Data             time.Time          `gorm:"column:data"`
}

const mappaConquistaTable = "mappa_conquistas"

func init() {
	RegisterEntity(Entity{EntityType: MappaConquista{}, TableName: mappaConquistaTable})
}

func (MappaConquista) TableName() string {
	return mappaConquistaTable
}
