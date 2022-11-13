package dtos

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
)

func MappaEspecialidadeToEntity(p *responses.MappaEspecialidadeResponse) *entities.MappaEspecialidade {

	return &entities.MappaEspecialidade{
		Codigo:           p.Codigo,
		Descricao:        p.Descricao,
		RamoConhecimento: p.RamoConhecimento,
		PreRequisitos:    p.PreRequisito,
		Itens:            MappaEspecialidadeItensToEntities(p.Itens),
	}
}

func MappaEspecialidadeToResponse(p *entities.MappaEspecialidade) *responses.MappaEspecialidadeResponse {

	return &responses.MappaEspecialidadeResponse{
		Codigo:           p.Codigo,
		Descricao:        p.Descricao,
		RamoConhecimento: p.RamoConhecimento,
		PreRequisito:     p.PreRequisitos,
		Itens:            MappaEspecialidadeItensToResponse(p.Itens),
	}

}

func MappaEspecialidadeItensToEntities(items []responses.MappaEspecialidadeItemResponse) []entities.MappaEspecialidadeItem {
	result := make([]entities.MappaEspecialidadeItem, len(items))
	for i, item := range items {
		result[i] = entities.MappaEspecialidadeItem{
			CodigoEspecialidade: item.CodigoEspecialidade,
			Numero:              item.Numero,
			Descricao:           item.Descricao,
		}
	}
	return result
}

func MappaEspecialidadeItensToResponse(items []entities.MappaEspecialidadeItem) []responses.MappaEspecialidadeItemResponse {
	result := make([]responses.MappaEspecialidadeItemResponse, len(items))
	for i, item := range items {
		result[i] = responses.MappaEspecialidadeItemResponse{
			CodigoEspecialidade: item.CodigoEspecialidade,
			Numero:              item.Numero,
			Descricao:           item.Descricao,
		}
	}
	return result
}
