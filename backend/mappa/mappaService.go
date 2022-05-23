package mappa

import (
	"fmt"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/cache"
	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/repositories"
)

type MappaService struct {
	Cache      *cache.MappaCache
	Repository repositories.IRepository
}

func (svc *MappaService) GetProgressoes(ramo string) (progressoes []*responses.MappaProgressaoResponse, err error) {

	lastFetch := svc.Cache.GetLastEventTime("progressoes")
	if lastFetch.Before(time.Now().Add(-24 * time.Hour)) {
		err = updateMappaProgressoes(svc.Repository)
		if err != nil {
			return
		}
		svc.Cache.SetLastEventTime("progressoes", time.Now())
	}
	ramoProgressoes := domain.ParseRamo(ramo)

	progressoes, err = svc.Repository.GetProgressoes(ramoProgressoes)
	return
}

func updateMappaProgressoes(repository repositories.IRepository) (err error) {
	progressoesResponse := MappaProgressoesRequest()
	if len(progressoesResponse) == 0 {
		return fmt.Errorf("Não foi possível obter os dados da Mappa - Progressões")
	}
	return repository.UpdateMappaProgressoes(progressoesResponse)
}
