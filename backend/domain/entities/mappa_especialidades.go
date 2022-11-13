package entities

type (
	MappaEspecialidade struct {
		NoIdModel
		Codigo           int                      `gorm:"column:codigo;primary_key"`
		Descricao        string                   `gorm:"column:descricao"`
		RamoConhecimento string                   `gorm:"column:ramo_conhecimento"`
		PreRequisitos    string                   `gorm:"column:pre_requisitos"`
		Itens            []MappaEspecialidadeItem `gorm:"foreignKey:CodigoEspecialidade;references:Codigo;OnDelete:CASCADE"`
	}
	MappaEspecialidadeItem struct {
		NoIdModel
		CodigoEspecialidade int    `gorm:"column:codigo;primary_key"`
		Numero              int    `gorm:"column:numero;primary_key"`
		Descricao           string `gorm:"column:descricao"`
	}
)
