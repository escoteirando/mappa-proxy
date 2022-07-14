package entities

import "gorm.io/gorm"

type (
	MappaEspecialidade struct {
		NoIdModel
		Codigo           int                           `gorm:"column:codigo;primary_key"`
		Descricao        string                        `gorm:"column:descricao"`
		RamoConhecimento string                        `gorm:"column:ramo_conhecimento"`
		PreRequisitos    []MappaEspecialidadeRequisito `gorm:"foreignKey:CodigoEspecialidade;references:Codigo;OnDelete:CASCADE"`
		Items            []MappaEspecialidadeItem      `gorm:"foreignKey:CodigoEspecialidade;references:Codigo;OnDelete:CASCADE"`
	}
	MappaEspecialidadeItem struct {
		NoIdModel
		CodigoEspecialidade int    `gorm:"column:codigo;primary_key"`
		Numero              int    `gorm:"column:numero;primary_key"`
		Descricao           string `gorm:"column:descricao"`
	}
	MappaEspecialidadeRequisito struct {
		gorm.Model
		CodigoEspecialidade int    `gorm:"column:codigo;primary_key"`
		Requisito           string `gorm:"column:requisito"`
	}
)
