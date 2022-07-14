package entities

type Grupo struct {
	NoIdModel
	Codigo           int    `gorm:"field:codigo;primary_key"`
	CodigoRegiao     string `gorm:"field:codigoRegiao"`
	Nome             string `gorm:"field:nome"`
	CodigoModalidade int    `gorm:"field:codigoModalidade"`
}
