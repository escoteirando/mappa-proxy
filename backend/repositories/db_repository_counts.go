package repositories

import (
	"log"

	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
)

func (r *DBRepository) GetCounts() (counts map[string]int, err error) {
	counts = make(map[string]int, 0)
	r.DBLock()
	defer r.DBUnlock()
	db := r.GetDBFunc()	
	for tableName := range entities.GetEntities() {
		var count int64
		err = db.Table(tableName).Count(&count).Error
		if err == nil {
			counts[tableName] = int(count)
		} else {
			log.Printf("Error getting count for table %s: %s", tableName, err)
			counts[tableName] = -1
		}
	}
	return counts, err
}
