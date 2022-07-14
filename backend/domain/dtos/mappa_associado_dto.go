package dtos

import (
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/types"
)

func MappaAssociadoToEntity(a *responses.MappaAssociadoResponse) *entities.Associado {
	return &entities.Associado{
		Codigo:                  a.Codigo,
		Nome:                    a.Nome,
		CodigoFoto:              a.CodigoFoto,
		CodigoEquipe:            a.CodigoEquipe,
		UserName:                a.UserName,
		NumeroDigito:            a.NumeroDigito,
		DataNascimento:          a.DataNascimento.Time,
		DataValidade:            a.DataValidade.Time,
		NomeAbreviado:           a.NomeAbreviado,
		Sexo:                    a.Sexo,
		CodigoRamo:              a.CodigoRamo,
		CodigoCategoria:         a.CodigoCategoria,
		CodigoSegundaCategoria:  a.CodigoSegundaCategoria,
		CodigoTerceiraCategoria: a.CodigoTerceiraCategoria,
		LinhaFormacao:           a.LinhaFormacao,
		CodigoRamoAdulto:        a.CodigoRamoAdulto,
		DataAcompanhamento:      a.DataAcompanhamento.Time,
	}
}

func MappaAssociadoToResponse(a *entities.Associado) *responses.MappaAssociadoResponse {
	return &responses.MappaAssociadoResponse{
		Codigo:                  a.Codigo,
		Nome:                    a.Nome,
		CodigoFoto:              a.CodigoFoto,
		CodigoEquipe:            a.CodigoEquipe,
		UserName:                a.UserName,
		NumeroDigito:            a.NumeroDigito,
		DataNascimento:          types.Date(a.DataNascimento),
		DataValidade:            types.Date(a.DataValidade),
		NomeAbreviado:           a.NomeAbreviado,
		Sexo:                    a.Sexo,
		CodigoRamo:              a.CodigoRamo,
		CodigoCategoria:         a.CodigoCategoria,
		CodigoSegundaCategoria:  a.CodigoSegundaCategoria,
		CodigoTerceiraCategoria: a.CodigoTerceiraCategoria,
		LinhaFormacao:           a.LinhaFormacao,
		CodigoRamoAdulto:        a.CodigoRamoAdulto,
		DataAcompanhamento:      types.Date(a.DataAcompanhamento),
	}
}

func MappaAssociadosToEntity(a []*responses.MappaAssociadoResponse) []entities.Associado {
	var associados []entities.Associado
	for _, associado := range a {
		associados = append(associados, *MappaAssociadoToEntity(associado))
	}
	return associados
}

func MappaAssociadosToResponse(a []entities.Associado) []*responses.MappaAssociadoResponse {
	var associados []*responses.MappaAssociadoResponse
	for _, associado := range a {
		associados = append(associados, MappaAssociadoToResponse(&associado))
	}
	return associados
}
