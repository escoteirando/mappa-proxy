package mappa

import (
	"fmt"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/dtos"
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/tools"
)

func (svc *MappaService) GetMarcacoes(codigoSecao int, authorization string) (marcacoes []*responses.MappaMarcacaoResponse, err error) {
	keyMarcacoes := fmt.Sprintf("marcacoes_%d", codigoSecao)
	lastFetch := svc.Cache.GetLastEventTime(keyMarcacoes)
	if lastFetch.Before(tools.DaysAgo(MARCACOES_UPDATE_INTERVAL)) {
		objMarcacoes, err := svc.API.GetMarcacoes(codigoSecao, lastFetch, authorization)
		if err == nil && objMarcacoes != nil && objMarcacoes.Values != nil && len(objMarcacoes.Values) > 0 {
			marcacoes = objMarcacoes.Values
			if err = svc.updateMarcacoes(codigoSecao, marcacoes); err == nil {
				svc.Cache.SetLastEventTime(keyMarcacoes, time.Now())
			}
		}

		err = svc.updateMappaProgressoes(svc.Repository)
		if err != nil {
			return nil, err
		}
	} else {
		objMarcacoes, err := svc.Repository.GetMarcacoes(codigoSecao)
		if err == nil && objMarcacoes != nil && len(objMarcacoes) > 0 {
			marcacoes = make([]*responses.MappaMarcacaoResponse, len(objMarcacoes))
			for i, marcacao := range objMarcacoes {
				marcacoes[i] = dtos.MappaMarcacaoToResponse(marcacao)
			}
		}

	}
	return
}

func (svc *MappaService) updateMarcacoes(codigoSecao int, marcacoes []*responses.MappaMarcacaoResponse) error {
	eMarcacoes := make([]*entities.MappaMarcacao, len(marcacoes))
	for i, marcacao := range marcacoes {
		eMarcacoes[i] = dtos.MappaMarcacaoToEntity(marcacao, codigoSecao)
	}
	return svc.Repository.UpdateMappaMarcacoes(eMarcacoes)
}
