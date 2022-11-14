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

func (svc *MappaService) GetConquistasFull(codigoSecao int, authorization string) (conquistas []*responses.FullConquistaResponse, err error) {
	c, err := svc.GetConquistas(codigoSecao, authorization)
	if err != nil {
		return nil, err
	}
	associados := make(map[int]*responses.MappaAssociadoResponse)
	getAssoc := func(codAssociado int) *responses.MappaAssociadoResponse {
		if a, ok := associados[codAssociado]; ok {
			return a
		}
		a := svc.GetAssociado(codAssociado, authorization)
		associados[codAssociado] = a
		return a
	}

	especialidades := make(map[int]*responses.MappaEspecialidadeResponse)
	getEspec := func(codEspecialidade int) *responses.MappaEspecialidadeResponse {
		if e, ok := especialidades[codEspecialidade]; ok {
			return e
		}
		var especialidade *responses.MappaEspecialidadeResponse
		if esp, err := svc.Repository.GetEspecialidade(codEspecialidade); err == nil {
			especialidade = dtos.MappaEspecialidadeToResponse(esp)
		}
		especialidades[codEspecialidade] = especialidade
		return especialidade
	}

	conquistas = make([]*responses.FullConquistaResponse, len(c))
	for i, conquista := range c {
		conquistas[i] = &responses.FullConquistaResponse{
			Type:             conquista.Type,
			Data:             conquista.Data.Time,
			CodAssociado:     conquista.CodAssociado,
			CodEscotista:     conquista.CodEscotista,
			NumeroNivel:      conquista.NumeroNivel,
			CodEspecialidade: conquista.CodEspecialidade,
			Associado:        getAssoc(conquista.CodAssociado),
			Especialidade:    getEspec(conquista.CodEspecialidade),
		}
	}
	return
}
