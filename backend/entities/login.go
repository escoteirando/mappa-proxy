package entities

import "time"

type Login struct {
	NoIdModel
	UserName        string    `gorm:"column:user_name;primaryKey"`
	LoginResponse   string    `gorm:"column:login_response"`
	PasswordHash    uint64    `gorm:"column:password_hash"`
	LastLogin       time.Time `gorm:"column:last_login"`
	MappaUserId     int       `gorm:"column:mappa_user_id"`
	MappaAuth       string    `gorm:"column:mappa_auth"`
	MappaValidUntil time.Time `gorm:"column:mappa_valid_until"`
}

// CREATE TABLE IF NOT EXISTS logins (
// 	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
// 	username TEXT NOT NULL,
// 	login_response TEXT NOT NULL,
// 	password_hash INT NOT NULL,
// 	last_login TEXT NOT NULL,
// 	mappa_userId INT NOT NULL,
// 	mappa_auth TEXT NOT NULL,
// 	mappa_validUntil INT NOT NULL,
// 	UNIQUE(username) ON CONFLICT REPLACE
// );
// CREATE INDEX IF NOT EXISTS idx_last_login ON logins (last_login);
