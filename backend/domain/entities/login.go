package entities

import "time"

type Login struct {
	NoIdModel
	UserName        string    `gorm:"column:user_name;primaryKey"`
	PasswordHash    uint32    `gorm:"column:password_hash"`
	LastLogin       time.Time `gorm:"column:last_login"`
	MappaUserId     int       `gorm:"column:mappa_user_id"`
	MappaAuth       string    `gorm:"column:mappa_auth"`
	MappaValidUntil time.Time `gorm:"column:mappa_valid_until"`
}

const loginTable = "logins"

func init() {
	RegisterEntity(Entity{EntityType: Login{}, TableName: loginTable})
}

func (Login) TableName() string {
	return loginTable
}
