package entities

import "time"

type Associado struct {
	NoIdModel
	Codigo                  int       `gorm:"column:codigo;primaryKey"`
	Nome                    string    `gorm:"column:nome"`
	CodigoFoto              int       `gorm:"column:codigo_foto"`
	CodigoEquipe            int       `gorm:"column:codigo_equipe"`
	UserName                int       `gorm:"column:username"`
	NumeroDigito            int       `gorm:"column:numero_digito"`
	DataNascimento          time.Time `gorm:"column:data_nascimento"`
	DataValidade            time.Time `gorm:"column:data_validade"`
	NomeAbreviado           string    `gorm:"column:nome_abreviado"`
	Sexo                    string    `gorm:"column:sexo"`
	CodigoRamo              int       `gorm:"column:codigo_ramo"`
	CodigoCategoria         int       `gorm:"column:codigo_categoria"`
	CodigoSegundaCategoria  int       `gorm:"column:codigo_segunda_categoria"`
	CodigoTerceiraCategoria int       `gorm:"column:codigo_terceira_categoria"`
	LinhaFormacao           string    `gorm:"column:linha_formacao"`
	CodigoRamoAdulto        int       `gorm:"column:codigo_ramo_adulto"`
	DataAcompanhamento      time.Time `gorm:"column:data_acompanhamento"`
	CodigoSecao             int       `gorm:"column:codigo_secao"`
}

const associadoTable = "associados"

func init() {
	RegisterEntity(Entity{EntityType: Associado{}, TableName: associadoTable})
}

func (Associado) TableName() string {
	return associadoTable
}
