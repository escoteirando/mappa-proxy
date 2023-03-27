package entities

type Secao struct {
	NoIdModel
	Codigo          uint   `gorm:"field:codigo;primary_key"`
	CodigoRegiao    string `gorm:"field:codigo_regiao;primary_key"`
	Nome            string `gorm:"field:nome"`
	CodigoTipoSecao uint   `gorm:"field:codigo_tipo_secao"`
	CodigoGrupo     uint   `gorm:"field:codigo_grupo"`
}

const secaoTable = "secoes"

func init() {
	RegisterEntity(Entity{EntityType: Secao{}, TableName: secaoTable})
}

func (Secao) TableName() string {
	return secaoTable
}
