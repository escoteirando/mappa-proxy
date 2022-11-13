package dtos

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/types"
)

func MappaConquistaToEntity(p *responses.MappaConquistaResponse, codigoSecao int) *entities.MappaConquista {
	return &entities.MappaConquista{
		Type:             p.Type,
		Data:             p.Data.Time,
		CodAssociado:     p.CodAssociado,
		CodEscotista:     p.CodEscotista,
		NumeroNivel:      p.NumeroNivel,
		CodEspecialidade: p.CodEspecialidade,
		CodigoSecao:      codigoSecao,
	}

}

func MappaConquistaToResponse(p *entities.MappaConquista) *responses.MappaConquistaResponse {
	return &responses.MappaConquistaResponse{
		Type:             p.Type,
		Data:             types.Date(p.Data),
		CodAssociado:     p.CodAssociado,
		CodEscotista:     p.CodEscotista,
		NumeroNivel:      p.NumeroNivel,
		CodEspecialidade: p.CodEspecialidade,
	}

}
