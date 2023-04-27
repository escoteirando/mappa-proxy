package repositories

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"gorm.io/gorm/clause"
)

func (r *DBRepository) SetKeyValue(key, value string, timeToLive time.Duration) error {
	r.DBLock()
	defer r.DBUnlock()
	var validUntil time.Time
	if timeToLive == time.Duration(0) {
		validUntil = time.Date(2099, 12, 31, 23, 59, 59, 0, time.UTC)
	} else {
		validUntil = time.Now().Add(timeToLive)
	}

	keyValue := entities.KeyValue{
		Key:        key,
		Value:      value,
		ValidUntil: validUntil,
	}
	db := r.GetDBFunc()
	res := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&keyValue)
	return res.Error
}

func (r *DBRepository) GetKeyValue(key, defaultValue string) string {
	r.DBLock()
	defer r.DBUnlock()
	keyValue := entities.KeyValue{
		Key: key,
	}
	db := r.GetDBFunc()
	res := db.First(&keyValue)
	if res.Error != nil {
		return defaultValue
	}
	if keyValue.ValidUntil.Before(time.Now()) {
		db.Delete(&keyValue)
		return defaultValue
	}
	return keyValue.Value
}

func (r *DBRepository) DeleteKey(key string) error {
	r.DBLock()
	defer r.DBUnlock()
	keyValue := entities.KeyValue{
		Key: key,
	}
	db := r.GetDBFunc()
	res := db.First(&keyValue)
	if res.Error != nil {
		return nil
	}
	return db.Delete(&keyValue).Error
}

func (r *DBRepository) DeleteKeys(keyPrefix string) error {
	r.DBLock()
	defer r.DBUnlock()
	var keys []entities.KeyValue
	db := r.GetDBFunc()
	result := db.Where("key like ?", keyPrefix+"%").Find(&keys)
	if result.Error != nil {
		return result.Error
	}
	for _, key := range keys {
		db.Delete(&key)
	}

	return nil
}
