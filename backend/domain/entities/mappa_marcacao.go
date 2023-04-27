package entities

import (
	"time"

	"gorm.io/gorm"
)

type MappaMarcacao struct {
	gorm.Model
	CodigoAtividade       int             `gorm:"field:cod_atividade;unique_index:idx_marcacao"`
	CodigoAssociado       int             `gorm:"field:cod_associado;unique_index:idx_marcacao"`
	CodigoSecao           int             `gorm:"field:cod_secao;unique_index:idx_marcacao"`
	DataAtividade         time.Time       `gorm:"field:data_atividade"`
	DataStatusEscotista   time.Time       `gorm:"field:data_status_escotista"`
	DataHoraAtualizacao   time.Time       `gorm:"field:data_hora_atualizacao"`
	CodigoUltimoEscotista int             `gorm:"field:cod_ultimo_escotista"`
	Segmento              string          `gorm:"field:segmento"`
	Atividade             MappaProgressao `gorm:"foreignKey:codigo;association_foreignKey:cod_atividade"`
	Associado             Associado       `gorm:"foreignKey:codigo;association_foreignKey:cod_associado"`
	Secao                 Secao           `gorm:"foreignKey:codigo;association_foreignKey:cod_secao"`
}

const mappaMarcacaoTable = "mappa_marcacoes"

func init() {
	RegisterEntity(Entity{EntityType: MappaMarcacao{}, TableName: mappaMarcacaoTable})
}

func (MappaMarcacao) TableName() string {
	return mappaMarcacaoTable
}
