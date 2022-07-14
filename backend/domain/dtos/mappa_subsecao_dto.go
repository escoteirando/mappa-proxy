package dtos

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
)

func MappaSubSecaoToEntity(s *responses.MappaSubSecaoResponse) entities.SubSecao {
	return entities.SubSecao{
		Codigo:          s.Codigo,
		Nome:            s.Nome,
		CodigoSecao:     s.CodigoSecao,
		CodigoLider:     s.CodigoLider,
		CodigoViceLider: s.CodigoViceLider,
		Associados:      MappaAssociadosToEntity(s.Associados),
	}
}

func MappaSubSecaoToResponse(s entities.SubSecao) *responses.MappaSubSecaoResponse {
	return &responses.MappaSubSecaoResponse{
		Codigo:          s.Codigo,
		Nome:            s.Nome,
		CodigoSecao:     s.CodigoSecao,
		CodigoLider:     s.CodigoLider,
		CodigoViceLider: s.CodigoViceLider,
		Associados:      MappaAssociadosToResponse(s.Associados),
	}
}

func MappaSubSecoesToEntity(s []*responses.MappaSubSecaoResponse) []entities.SubSecao {
	subsecoes := make([]entities.SubSecao, len(s))
	for i, subsecao := range s {
		subsecoes[i] = MappaSubSecaoToEntity(subsecao)
	}
	return subsecoes
}

func MappaSubSecoesToResponse(s []entities.SubSecao) []*responses.MappaSubSecaoResponse {
	subsecoes := make([]*responses.MappaSubSecaoResponse, len(s))
	for i, subsecao := range s {
		subsecoes[i] = MappaSubSecaoToResponse(subsecao)
	}
	return subsecoes
}
