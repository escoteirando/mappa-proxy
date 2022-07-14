package entities

import "time"

type KeyValue struct {
	NoIdModel
	Key        string    `gorm:"column:key;primary_key"`
	Value      string    `gorm:"column:value"`
	ValidUntil time.Time `gorm:"column:valid_until"`
}
