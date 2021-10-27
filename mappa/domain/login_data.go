package domain

import "time"

type LoginData struct {
	LoginResponse LoginResponse
	UserName      string
	PasswordHash  uint64
	LastLogin     time.Time
	Deleted       bool
}
