package dtos

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
)

func MappaSecaoToEntity(s *responses.MappaSecaoResponse) *entities.Secao {
	return &entities.Secao{
		Codigo:          s.Codigo,
		CodigoRegiao:    s.CodigoRegiao,
		Nome:            s.Nome,
		CodigoTipoSecao: s.CodigoTipoSecao,
		CodigoGrupo:     s.CodigoGrupo,
		// SubSecoes:       MappaSubSecoesToEntity(s.Subsecoes),
	}
}

func MappaSecaoToResponse(s *entities.Secao) *responses.MappaSecaoResponse {
	return &responses.MappaSecaoResponse{
		Codigo:          s.Codigo,
		Nome:            s.Nome,
		CodigoRegiao:    s.CodigoRegiao,
		CodigoTipoSecao: s.CodigoTipoSecao,
		CodigoGrupo:     s.CodigoGrupo,
		// Subsecoes:       MappaSubSecoesToResponse(s.SubSecoes),
	}
}
