package entities

import "gorm.io/gorm"

type AssociadoSecao struct {
	gorm.Model
	CodigoAssociado int  `gorm:"column:cod_associado;unique_index:idx_associado_secao"`
	CodigoSecao     int  `gorm:"column:cod_secao;unique_index:idx_associado_secao"`
	SubSecao        bool `gorm:"column:subsecao"`
}

const associadoSecaoTable = "associado_secoes"

func init() {
	RegisterEntity(Entity{EntityType: AssociadoSecao{}, TableName: associadoSecaoTable})
}

func (AssociadoSecao) TableName() string {
	return associadoSecaoTable
}
