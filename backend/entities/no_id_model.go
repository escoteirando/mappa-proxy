package entities

import (
	"time"

	"gorm.io/gorm"
)

// Model with no primary key ID
type NoIdModel struct{
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}