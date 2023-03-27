package entities

import (
	"fmt"
	"sync"
)

type Entity struct {
	EntityType interface{}
	TableName  string
}

var (
	entities   = make(map[string]Entity)
	entityLock sync.RWMutex
)

func RegisterEntity(entity Entity) {
	entityLock.Lock()
	entities[entity.TableName] = entity
	entityLock.Unlock()
}

func GetEntities() map[string]Entity {
	entityLock.RLock()
	defer entityLock.RUnlock()
	return entities
}

func GetTableName(entity interface{}) string {
	entityLock.RLock()
	defer entityLock.RUnlock()
	for _, e := range entities {
		if fmt.Sprintf("%T", e.EntityType) == fmt.Sprintf("%T", entity) {
			return e.TableName
		}
	}
	return ""
}
