package mappa

import (
	"fmt"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/dtos"
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/tools"
)

func (svc *MappaService) GetConquistas(codigoSecao int, authorization string) (conquistas []*responses.MappaConquistaResponse, err error) {
	keyConquistas := fmt.Sprintf("conquistas_%d", codigoSecao)
	lastFetch := svc.Cache.GetLastEventTime(keyConquistas)
	if lastFetch.Before(tools.DaysAgo(CONQUISTAS_UPDATE_INTERVAL)) {
		objConquistas, err := svc.API.GetConquistas(codigoSecao, lastFetch, authorization)
		if err == nil && objConquistas != nil && objConquistas.Values != nil && len(objConquistas.Values) > 0 {
			conquistas = objConquistas.Values
			if err = svc.updateConquistas(codigoSecao, conquistas); err == nil {
				svc.Cache.SetLastEventTime(keyConquistas, time.Now())
			}
		}

	} else {
		objConquistas, err := svc.Repository.GetConquistas(codigoSecao)
		if err == nil && objConquistas != nil && len(objConquistas) > 0 {
			conquistas = make([]*responses.MappaConquistaResponse, len(objConquistas))
			for i, conquista := range objConquistas {
				conquistas[i] = dtos.MappaConquistaToResponse(conquista)

			}
		}
	}
	return
}

func (svc *MappaService) updateConquistas(codigoSecao int, conquistas []*responses.MappaConquistaResponse) error {
	eConquistas := make([]*entities.MappaConquista, len(conquistas))
	for i, conquista := range conquistas {
		eConquistas[i] = dtos.MappaConquistaToEntity(conquista, codigoSecao)
	}
	return svc.Repository.UpdateMappaConquistas(eConquistas)
}
