package dtos

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/types"
)

func MappaMarcacaoToEntity(p *responses.MappaMarcacaoResponse, codigoSecao int) *entities.MappaMarcacao {
	return &entities.MappaMarcacao{
		CodigoAtividade:       p.CodigoAtividade,
		CodigoAssociado:       p.CodigoAssociado,
		CodigoSecao:           codigoSecao,
		DataAtividade:         p.DataAtividade.Time,
		DataStatusEscotista:   p.DataStatusEscotista.Time,
		DataHoraAtualizacao:   p.DataHoraAtualizacao.Time,
		CodigoUltimoEscotista: p.CodigoUltimoEscotista,
		Segmento:              p.Segmento,
	}
}

func MappaMarcacaoToResponse(p *entities.MappaMarcacao) *responses.MappaMarcacaoResponse {
	return &responses.MappaMarcacaoResponse{
		CodigoAtividade:       p.CodigoAtividade,
		CodigoAssociado:       p.CodigoAssociado,
		DataAtividade:         types.Date(p.DataAtividade),
		DataStatusEscotista:   types.Date(p.DataStatusEscotista),
		DataHoraAtualizacao:   types.Date(p.DataHoraAtualizacao),
		CodigoUltimoEscotista: p.CodigoUltimoEscotista,
		Segmento:              p.Segmento,
	}
}
