package entities

type Escotista struct {
	NoIdModel
	UserId          uint   `gorm:"column:user_id;primary_key"`
	CodigoAssociado uint   `gorm:"column:cod_associado"`
	UserName        string `gorm:"column:username"`
	NomeCompleto    string `gorm:"column:nome_completo"`
	Ativo           bool   `gorm:"column:ativo"`
	CodigoGrupo     uint   `gorm:"column:codigo_grupo"`
	CodigoRegiao    string `gorm:"column:codigo_regiao"`
	CodigoFoto      uint   `gorm:"column:codigo_foto"`
}
