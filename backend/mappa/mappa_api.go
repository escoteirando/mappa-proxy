package mappa

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/escoteirando/mappa-proxy/backend/domain/requests"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
)

type MappaAPI struct {
	mappaUrl string
}

func NewMappaAPI(mappaURL ...string) *MappaAPI {
	url := URL
	if len(mappaURL) > 0 {
		url = mappaURL[0]
	}

	return &MappaAPI{mappaUrl: url}
}

func (api *MappaAPI) Login(username string, password string) (loginResponse *responses.MappaLoginResponse, err error) {
	logMsg := ""
	defer func() {
		log.Printf("Login: %s", logMsg)
	}()

	if len(username) == 0 || len(password) == 0 {
		err = fmt.Errorf("Invalid username or password")
		logMsg = "Login failed: " + err.Error()
		return
	}

	loginRequest := requests.CreateMappaLoginRequest(username, password)

	statusCode, responseBody, err := api.post("/api/escotistas/login", loginRequest)
	if err != nil {
		return
	}

	if statusCode > 0 && statusCode < 300 {
		err = json.Unmarshal(responseBody, &loginResponse)
		if err == nil {
			log.Printf("MAPPA login ok: %s", username)
		} else {
			log.Printf("MAPPA login response failed: %v", err)
		}
	} else {
		log.Printf("Fail on MAPPA login: StatusCode = %d Body = %v", statusCode, string(responseBody))
	}
	return
}

func (api *MappaAPI) GetEscotista(userId int, authorization string) (escotista *responses.MappaEscotistaResponse) {
	_, body, err := api.get(fmt.Sprintf("/api/escotistas/%d", userId), authorization)
	if err != nil {
		log.Printf("HTTP Error on GetEscotista: %v", err)
	} else {
		escotista = responses.GetMappaEscotistaResponseFromJSON(body)
		if escotista == nil {
			log.Printf("Error on GetEscotista: %s", string(body))
		}
	}
	return
}

func (api *MappaAPI) GetAssociado(codigoAssociado int, authorization string) (associado *responses.MappaAssociadoResponse) {
	_, body, err := api.get(fmt.Sprintf("/api/associados/%d", codigoAssociado), authorization)
	if err != nil {
		log.Printf("HTTP Error on GetAssociado: %v", err)
	} else {
		associado = responses.GetMappaAssociadoResponseFromJSON(body)
		if associado == nil {
			log.Printf("Error on GetAssociado: %s", string(body))
		}
	}
	return
}

func (api *MappaAPI) GetGrupo(codigoGrupo int, codigoRegiao string, authorization string) (grupo *responses.MappaGrupoResponse) {
	_, body, err := api.get(fmt.Sprintf("/api/grupos?filter={\"where\":{\"codigo\":%d,\"codigoRegiao\":\"%s\"}}", codigoGrupo, codigoRegiao), authorization)
	if err != nil {
		log.Printf("HTTP Error on GetGrupo: %v", err)
	} else {
		grupos := responses.GetMappaGruposResponseFromJSON(body)
		if len(grupos) == 0 {
			log.Printf("Error on GetGrupo: %s", string(body))
		}
		grupo = grupos[0]
	}
	return
}

func (api *MappaAPI) GetEscotistaSecoes(userId int, authorization string) (response []*responses.MappaSecaoResponse) {
	_, body, err := api.get(fmt.Sprintf("/api/escotistas/%d/secoes", userId), authorization)
	if err != nil {
		log.Printf("HTTP Error on GetEscotistaSecoes: %v", err)
		return
	}
	response, err = responses.GetMappaSecaoResponsesFromJSON(body)
	if err != nil {
		log.Printf("Error on GetEscotistaSecoes: %v", err)
	} else if response == nil {
		log.Printf("Error on GetEscotistaSecoes: %s", string(body))
	}
	for i, secao := range response {
		response[i].Subsecoes = api.GetEscotistaEquipe(userId, int(secao.Codigo), authorization)
	}
	return
}

func (api *MappaAPI) GetEscotistaEquipe(userId int, codSecao int, authorization string) (response []*responses.MappaSubSecaoResponse) {
	_, body, err := api.get(fmt.Sprintf("/api/escotistas/%d/secoes/%d/equipes?filter={\"include\":\"associados\"}", userId, codSecao), authorization)
	if err != nil {
		log.Printf("HTTP Error on GetEscotistaEquipe: %v", err)
		return
	}
	response, err = responses.GetSubSecoesFromJSON(body)
	if err != nil {
		log.Printf("Error on GetEscotistaEquipe: %v", err)
	}
	return
}

func (api *MappaAPI) GetProgressoes() (response []*responses.MappaProgressaoResponse, err error) {
	_, body, err := api.get("/api/progressao-atividades", "")
	if err != nil {
		log.Printf("HTTP Error on mappaGetProgressoes: %v", err)
		return
	}
	response, err = responses.GetMappaProgressaoResponsesFromJSON(body)
	if err != nil {
		log.Printf("Parsing Error on mappaGetProgressoes: %v", err)
	} else if response == nil {
		log.Printf("Empty response error on mappaGetProgressoes: %s", string(body))
	}
	return
}

func (api *MappaAPI) GetEspecialidades() (response []*responses.MappaEspecialidadeResponse, err error) {
	_, body, err := api.get("/api/especialidades?filter={\"include\":\"itens\"}", "")
	if err != nil {
		log.Printf("HTTP Error on mappaGetEspecialidades: %v", err)
		return nil, err
	}
	response, err = responses.GetMappaEspecialidadesResponseFromJSON(body)
	if err != nil {
		log.Printf("Parsing Error on mappaGetEspecialidades: %v", err)
		return nil, err
	}
	return
}
