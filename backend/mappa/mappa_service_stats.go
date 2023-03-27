package mappa

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
)

func (svc *MappaService) GetStats() (stats responses.StatsResponse, err error) {
	var counts map[string]int
	if counts, err = svc.Repository.GetCounts(); err != nil {
		return
	}

	tableName := entities.GetTableName(entities.Escotista{})
	stats.NumeroEscotistas, _ = counts[tableName]
	tableName = entities.GetTableName(entities.Associado{})
	stats.NumeroAssociados, _ = counts[tableName]
	tableName = entities.GetTableName(entities.Grupo{})
	stats.NumeroGrupos, _ = counts[tableName]
	tableName = entities.GetTableName(entities.Secao{})
	stats.NumeroSecoes, _ = counts[tableName]

	return
}
