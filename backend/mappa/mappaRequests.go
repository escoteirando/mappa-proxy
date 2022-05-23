package mappa

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/escoteirando/mappa-proxy/backend/domain/requests"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"github.com/escoteirando/mappa-proxy/backend/infra"
)

func MappaLoginRequest(username string, password string) (loginResponse responses.MappaLoginResponse, err error) {
	loginRequest := requests.CreateMappaLoginRequest(username, password)

	url := fmt.Sprintf("%s/api/escotistas/login", URL)
	log.Printf("Login request %v: %v", url, username)
	statusCode, responseBody, err := infra.HttpPost(url, loginRequest, "Login "+username)
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

func MappaEscotistaRequest(userId int, authorization string) (response *responses.MappaDetalhesResponse) {

	escotista := mappaGetEscotista(userId, authorization)
	if escotista == nil {
		return
	}
	associado := mappaGetAssociado(escotista.CodigoAssociado, authorization)
	if associado == nil {
		return
	}

	grupo := mappaGetGrupo(escotista.CodigoGrupo, escotista.CodigoRegiao, authorization)
	if grupo == nil {
		return
	}
	return &responses.MappaDetalhesResponse{
		Escotista: escotista,
		Associado: associado,
		Grupos:    grupo,
	}
}

func mappaGet(url string, authorization string, logMessage string) (body []byte, err error) {
	statusCode := 0
	log.Printf("mappaGet %s", url)
	statusCode, body, err = infra.HttpGet(url, map[string]string{
		"Authorization": authorization,
	}, logMessage)
	if err != nil {
		return
	}
	if statusCode >= 400 {
		err = fmt.Errorf("%s: %s - %d %s", logMessage, url, statusCode, string(body))
	}
	return
}

func mappaGetEscotista(userId int, authorization string) (escotista *responses.MappaEscotistaResponse) {
	body, err := mappaGet(fmt.Sprintf("%s/api/escotistas/%d", URL, userId), authorization, "Get Escotista")
	if err != nil {
		log.Printf("Error on mappaGetEscotista: %v", err)
	} else {
		escotista = responses.GetMappaEscotistaResponseFromJSON(body)
		if escotista == nil {
			log.Printf("Error on mappaGetEscotista: %s", string(body))
		}
	}
	return
}

func mappaGetAssociado(codigoAssociado int, authorization string) (associado *responses.MappaAssociadoResponse) {

	body, err := mappaGet(fmt.Sprintf("%s/api/associados/%d", URL, codigoAssociado), authorization, "Get Associado")
	if err != nil {
		log.Printf("Error on mappaGetEscotista: %v", err)
	} else {
		associado = responses.GetMappaAssociadoResponseFromJSON(body)
		if associado == nil {
			log.Printf("Error on mappaGetEscotista: %s", string(body))
		}
	}

	return
}

func mappaGetGrupo(codigoGrupo int, codigoRegiao string, authorization string) (grupo []*responses.MappaGrupoResponse) {
	body, err := mappaGet(fmt.Sprintf("%s/api/grupos?filter={\"where\":{\"codigo\":%d,\"codigoRegiao\":\"%s\"}}", URL, codigoGrupo, codigoRegiao), authorization, "Get Grupo")
	if err != nil {
		log.Printf("Error on mappaGetGrupo: %v", err)
	} else {
		grupo = responses.GetMappaGruposResponseFromJSON(body)
		if grupo == nil {
			log.Printf("Error on mappaGetGrupo: %s", string(body))
		}
	}

	return
}

func MappaEspecialidadesRequest(authorization string) (response []*responses.MappaEspecialidadeResponse) {
	body, err := mappaGet(fmt.Sprintf("%s/api/especialidades?filter={\"filter[include]\":\"itens\"}", URL), authorization, "Get Especialidades")
	if err != nil {

		log.Printf("Error on MappaEspecialidadesRequest: %v", err)
	} else {
		especialidades, err := responses.GetMappaEspecialidadesResponseFromJSON(body)
		if err != nil {
			log.Printf("Error on MappaEspecialidadesRequest: %v", err)
		} else if especialidades == nil {
			log.Printf("Error on MappaEspecialidadesRequest: %s", string(body))
		}
	}
	return
}

func MappaProgressoesRequest() (response []*responses.MappaProgressaoResponse) {
	body, err := mappaGet(fmt.Sprintf("%s/api/progressao-atividades", URL), "authorization", "Get Progressoes")
	if err != nil {
		log.Printf("Error on MappaProgressoesRequest: %v", err)
	} else {
		response, err := responses.GetMappaProgressaoResponsesFromJSON(body)
		if err != nil {
			log.Printf("Error on MappaProgressoesRequest: %v", err)
		} else if response == nil {
			log.Printf("Error on MappaProgressoesRequest: %s", string(body))
		}
		return response
	}
	return
}
