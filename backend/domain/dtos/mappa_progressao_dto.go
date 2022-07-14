package dtos

import (
	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
)

func MappaProgressaoToEntity(p *responses.MappaProgressaoResponse) *entities.MappaProgressao {
	return &entities.MappaProgressao{
		Codigo:                p.Codigo,
		Descricao:             p.Descricao,
		CodigoUEB:             p.CodigoUEB,
		Ordenacao:             p.Ordenacao,
		CodigoCaminho:         p.CodigoCaminho,
		CodigoDesenvolvimento: p.CodigoDesenvolvimento,
		Segmento:              p.Segmento,
		NumeroGrupo:           p.NumeroGrupo,
		CodigoRegiao:          p.CodigoRegiao,
		CodigoCompetencia:     p.CodigoCompetencia,
		Ramo:                  domain.CaminhoToRamo(p.CodigoCaminho),
	}
}

func MappaProgressaoToResponse(p *entities.MappaProgressao) *responses.MappaProgressaoResponse {
	return &responses.MappaProgressaoResponse{
		Codigo:                p.Codigo,
		Descricao:             p.Descricao,
		CodigoUEB:             p.CodigoUEB,
		Ordenacao:             p.Ordenacao,
		CodigoCaminho:         p.CodigoCaminho,
		CodigoCompetencia:     p.CodigoCompetencia,
		CodigoDesenvolvimento: p.CodigoDesenvolvimento,
		Segmento:              p.Segmento,
		NumeroGrupo:           p.NumeroGrupo,
		CodigoRegiao:          p.CodigoRegiao,
	}
}
