package entities

type SubSecao struct {
	NoIdModel
	Codigo          uint        `gorm:"field:codigo;primary_key"`
	Nome            string      `gorm:"field:nome"`
	CodigoSecao     uint        `gorm:"field:codigo_secao"`
	CodigoLider     uint        `gorm:"field:codigo_lider"`
	CodigoViceLider uint        `gorm:"field:codigo_vice_lider"`
	Associados      []Associado `gorm:"foreignKey:codigo_equipe;association_foreignKey:codigo"`
}
