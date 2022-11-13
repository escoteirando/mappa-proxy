package mappa

import (
	"fmt"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/repositories"
	"github.com/escoteirando/mappa-proxy/backend/tools"
)

func (svc *MappaService) GetProgressoes(ramo string) (progressoes []*responses.MappaProgressaoResponse, err error) {

	lastFetch := svc.Cache.GetLastEventTime("progressoes")
	if lastFetch.Before(tools.DaysAgo(PROGRESSOES_UPDATE_INTERVAL)) {
		err = svc.updateMappaProgressoes(svc.Repository)
		if err != nil {
			return
		}
		svc.Cache.SetLastEventTime("progressoes", time.Now())
	}
	ramoProgressoes := domain.ParseRamo(ramo)

	progressoes, err = svc.Repository.GetProgressoes(ramoProgressoes)
	return
}

func (svc *MappaService) updateMappaProgressoes(repository repositories.IRepository) (err error) {
	progressoes, err := svc.API.GetProgressoes()
	if err != nil || len(progressoes) == 0 {
		return fmt.Errorf("Não foi possível obter as progressoes do MAPPA")
	}
	return repository.UpdateMappaProgressoes(progressoes)
}