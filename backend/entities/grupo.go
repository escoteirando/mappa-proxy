package entities

type Grupo struct {
	NoIdModel
	Codigo           int    `gorm:"field:codigo;primary_key"`
	CodigoRegiao     string `gorm:"field:codigoRegiao;primary_key"`
	Nome             string `gorm:"field:nome"`
	CodigoModalidade int    `gorm:"field:codigoModalidade"`
}
