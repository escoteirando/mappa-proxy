package entities

type Grupo struct {
	NoIdModel
	Codigo           int    `gorm:"field:codigo;primary_key"`
	CodigoRegiao     string `gorm:"field:codigoRegiao"`
	Nome             string `gorm:"field:nome"`
	CodigoModalidade int    `gorm:"field:codigoModalidade"`
}

const grupoTable = "grupos"

func init() {
	RegisterEntity(Entity{EntityType: Grupo{}, TableName: grupoTable})
}

func (Grupo) TableName() string {
	return grupoTable
}
