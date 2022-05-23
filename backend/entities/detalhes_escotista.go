package entities

type DetalhesEscotista struct {
	NoIdModel
	UserId   uint   `gorm:"column:user_id;primary_key"`
	Detalhes string `gorm:"column:detalhes"`
}

// CREATE TABLE IF NOT EXISTS detalhes_escotista (
// 	user_id integer NOT NULL PRIMARY KEY,
// 	creation_date DATE NOT NULL,
// 	detalhes TEXT NOT NULL,
// 	UNIQUE(user_id) ON CONFLICT REPLACE
// );
// CREATE INDEX IF NOT EXISTS idx_user_id ON detalhes_escotista (user_id);
