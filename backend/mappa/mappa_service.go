package mappa

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/cache"
	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/dtos"
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/repositories"
	"github.com/escoteirando/mappa-proxy/backend/tools"
	"github.com/escoteirando/mappa-proxy/backend/types"
)

type MappaService struct {
	Cache      *cache.MappaCache
	Repository repositories.IRepository
	API        *MappaAPI
}

func (svc *MappaService) Login(username string, password string) (response *responses.MappaLoginResponse, err error) {
	response, err = svc.API.Login(username, password)
	if err != nil {
		return nil, fmt.Errorf("UNAUTHORIZED")
	}
	if err == nil {
		if errCache := svc.Cache.SetLogin(username, password, *response); errCache != nil {
			log.Printf("Erro ao salvar login em cache: %s", errCache.Error())
		}
		// Atualizar dados do escotista
		svc.GetEscotistaDetalhes(response.Userid, response.ID)

	}
	return
}

func (svc *MappaService) GetEscotistaDetalhes(userId int, authorization string) (response *responses.MappaDetalhesResponse) {
	escotista := svc.GetEscotista(userId, authorization)
	associado := svc.GetAssociado(escotista.CodigoAssociado, authorization)
	grupo := svc.GetGrupo(escotista.CodigoGrupo, escotista.CodigoRegiao, authorization)
	response = &responses.MappaDetalhesResponse{
		Associado: associado,
		Grupos:    grupo,
		Escotista: escotista,
	}

	return
}

func (svc *MappaService) GetEscotista(userId int, authorization string) (response *responses.MappaEscotistaResponse) {
	escotista, err := svc.Repository.GetEscotista(userId)
	if err != nil || escotista == nil || escotista.UpdatedAt.Before(time.Now().Add(-7*24*time.Hour)) {
		response = svc.API.GetEscotista(userId, authorization)
		if response != nil {
			escotista = &entities.Escotista{
				UserId:          uint(response.UserId),
				CodigoAssociado: uint(response.CodigoAssociado),
				CodigoGrupo:     uint(response.CodigoGrupo),
				UserName:        response.UserName,
				NomeCompleto:    response.NomeCompleto,
				Ativo:           response.Ativo.IsTrue(),
				CodigoRegiao:    response.CodigoRegiao,
				CodigoFoto:      uint(response.CodigoFoto),
			}
			if err := svc.Repository.SetEscotista(escotista); err != nil {
				log.Printf("Erro ao salvar escotista: %s", err.Error())
			}
		}
	}
	if escotista != nil {
		response = &responses.MappaEscotistaResponse{
			UserId:          int(escotista.UserId),
			CodigoAssociado: int(escotista.CodigoAssociado),
			CodigoGrupo:     int(escotista.CodigoGrupo),
			UserName:        escotista.UserName,
			NomeCompleto:    escotista.NomeCompleto,
			Ativo:           types.FromNativeBool(escotista.Ativo),
			CodigoRegiao:    escotista.CodigoRegiao,
			CodigoFoto:      int(escotista.CodigoFoto),
		}
	}
	return
}

func (svc *MappaService) GetAssociado(codigoAssociado int, authorization string) (response *responses.MappaAssociadoResponse) {
	associado, err := svc.Repository.GetAssociado(codigoAssociado)
	if err != nil || associado == nil || associado.UpdatedAt.Before(time.Now().Add(-7*24*time.Hour)) {
		response = svc.API.GetAssociado(codigoAssociado, authorization)
		if response != nil {
			associado = &entities.Associado{
				Codigo:                  response.Codigo,
				CodigoEquipe:            response.CodigoEquipe,
				Nome:                    response.Nome,
				CodigoFoto:              response.CodigoFoto,
				UserName:                response.UserName,
				NumeroDigito:            response.NumeroDigito,
				DataNascimento:          response.DataNascimento.Time,
				DataValidade:            response.DataValidade.Time,
				NomeAbreviado:           response.NomeAbreviado,
				Sexo:                    response.Sexo,
				CodigoRamo:              response.CodigoRamo,
				CodigoCategoria:         response.CodigoCategoria,
				CodigoSegundaCategoria:  response.CodigoSegundaCategoria,
				CodigoTerceiraCategoria: response.CodigoTerceiraCategoria,
				LinhaFormacao:           response.LinhaFormacao,
				CodigoRamoAdulto:        response.CodigoRamoAdulto,
				DataAcompanhamento:      response.DataAcompanhamento.Time,
			}
			if err := svc.Repository.SetAssociado(associado); err != nil {
				log.Printf("Erro ao salvar associado: %s", err.Error())
			}
		}
	}
	if associado != nil {
		response = &responses.MappaAssociadoResponse{
			Codigo:                  associado.Codigo,
			CodigoEquipe:            associado.CodigoEquipe,
			Nome:                    associado.Nome,
			CodigoFoto:              associado.CodigoFoto,
			UserName:                associado.UserName,
			NumeroDigito:            associado.NumeroDigito,
			DataNascimento:          types.Date(associado.DataNascimento),
			DataValidade:            types.Date(associado.DataValidade),
			NomeAbreviado:           associado.NomeAbreviado,
			Sexo:                    associado.Sexo,
			CodigoRamo:              associado.CodigoRamo,
			CodigoCategoria:         associado.CodigoCategoria,
			CodigoSegundaCategoria:  associado.CodigoSegundaCategoria,
			CodigoTerceiraCategoria: associado.CodigoTerceiraCategoria,
			LinhaFormacao:           associado.LinhaFormacao,
			CodigoRamoAdulto:        associado.CodigoRamoAdulto,
			DataAcompanhamento:      types.Date(associado.DataAcompanhamento),
		}
	}
	return
}

func (svc *MappaService) GetGrupo(codigoGrupo int, codigoRegiao string, authorization string) (response *responses.MappaGrupoResponse) {
	grupo, err := svc.Repository.GetGrupo(codigoGrupo, codigoRegiao)
	if err != nil || grupo == nil || grupo.UpdatedAt.Before(tools.DaysAgo(GRUPOS_UPDATE_INTERVAL)) {
		response = svc.API.GetGrupo(codigoGrupo, codigoRegiao, authorization)
		if response != nil {
			grupo = &entities.Grupo{
				Codigo:           response.Codigo,
				Nome:             response.Nome,
				CodigoModalidade: response.CodigoModalidade,
				CodigoRegiao:     response.CodigoRegiao,
			}
			if err := svc.Repository.SetGrupo(grupo); err != nil {
				log.Printf("Erro ao salvar grupo: %s", err.Error())
			}
		}
	}
	return
}

func (svc *MappaService) GetValidAuthorization() (authorization string, err error) {
	loginData := svc.Cache.GetLastLogin()
	if loginData == nil {

		return "", fmt.Errorf("N??o foi poss??vel obter o ??ltimo login")
	}
	if loginData.IsValid() {
		return loginData.Authorization, nil
	}
	return "", fmt.Errorf("O ??ltimo login n??o ?? v??lido")
}

func (svc *MappaService) GetProgressoes(ramo string) (progressoes []*responses.MappaProgressaoResponse, err error) {

	lastFetch := svc.Cache.GetLastEventTime("progressoes")
	if lastFetch.Before(tools.DaysAgo(PROGRESSOES_UPDATE_INTERVAL)) {
		err = svc.updateMappaProgressoes(svc.Repository)
		if err != nil {
			return
		}
		svc.Cache.SetLastEventTime("progressoes", time.Now())
	}
	ramoProgressoes := domain.ParseRamo(ramo)

	progressoes, err = svc.Repository.GetProgressoes(ramoProgressoes)
	return
}

func (svc *MappaService) updateMappaProgressoes(repository repositories.IRepository) (err error) {
	progressoes, err := svc.API.GetProgressoes()
	if err != nil || len(progressoes) == 0 {
		return fmt.Errorf("N??o foi poss??vel obter as progressoes do MAPPA")
	}
	return repository.UpdateMappaProgressoes(progressoes)
}

func (svc *MappaService) GetEscotistaSecoes(userId int, authorization string) (secoes []*responses.MappaSecaoResponse, err error) {
	eventKey := "escotista_secoes_" + strconv.Itoa(userId)
	lastFetch := svc.Cache.GetLastEventTime(eventKey)
	if lastFetch.Before(tools.DaysAgo(ESCOTISTA_SECOES_UPDATE_INTERVAL)) {
		eSecoes := svc.API.GetEscotistaSecoes(userId, authorization)
		if len(eSecoes) == 0 {
			return nil, fmt.Errorf("N??o foi poss??vel obter as se????es do escotista")
		}
		secoes = make([]*responses.MappaSecaoResponse, len(eSecoes))
		for i, eSecao := range eSecoes {
			secao := dtos.MappaSecaoToEntity(eSecao)

			// for j, eSubSecao := range eSecao.Subsecoes {
			// 	secao.SubSecoes[j] = entities.SubSecao{
			// 		Codigo:          eSubSecao.Codigo,
			// 		Nome:            eSubSecao.Nome,
			// 		CodigoSecao:     eSubSecao.CodigoSecao,
			// 		CodigoLider:     eSubSecao.CodigoLider,
			// 		CodigoViceLider: eSubSecao.CodigoViceLider,
			// 		Associados:      make([]entities.Associado, len(eSubSecao.Associados)),
			// 	}
			// }

			svc.Repository.SetSecao(secao)

			
			secoes[i] = &responses.MappaSecaoResponse{
				CodigoTipoSecao: eSecao.CodigoTipoSecao,
				CodigoGrupo:     eSecao.CodigoGrupo,
				CodigoRegiao:    eSecao.CodigoRegiao,
				Subsecoes:       eSecao.Subsecoes,
				Codigo:          eSecao.Codigo,
				Nome:            eSecao.Nome,
			}
		}
		svc.Cache.SetLastEventTime(eventKey, time.Now())
	}
	return
}
