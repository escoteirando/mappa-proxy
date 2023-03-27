package mappa

import (
	"os"
	"testing"

	"github.com/escoteirando/mappa-proxy/backend/domain/requests"
	"github.com/escoteirando/mappa-proxy/backend/infra"
)

type AuthData struct {
	UserName string
	Password string
}

var authData *AuthData

const mappa = "http://mappa.escoteiros.org.br"

// Authorization data for tests
func getAuthData(t *testing.T) *AuthData {
	if authData != nil {
		return authData
	}

	// Create a new login data
	username := os.Getenv("AUTH_USERNAME")
	password := os.Getenv("AUTH_PASSWORD")
	if len(username) == 0 || len(password) == 0 {
		t.Errorf("Missing AUTH_USERNAME or AUTH_PASSWORD environment variables")
	}
	authData = &AuthData{
		UserName: username,
		Password: password,
	}

	return authData
}

func TestMappaAPI_Login(t *testing.T) {

	t.Run("Login", func(t *testing.T) {
		requester := infra.NewHttpRequestMock()

		requester.Setup(infra.RequestMock{
			Url:                mappa + "/api/escotistas/login",
			Data:               requests.CreateMappaLoginRequest("test", "test"),
			Method:             "POST",
			ResponseStatusCode: 201,
			ResponseBody:       []byte(`{"ID":"1234","ttl":1209600,"created":"2099-01-01T00:00:00.000Z","userId":1234}`),
		}).Setup(infra.RequestMock{
			Url:                mappa + "/api/escotistas/1234",
			Method:             "GET",
			ResponseStatusCode: 200,
			ResponseBody:       []byte(`{"codigoAssociado":1234,"codigoGrupo":12,"codigoRegiao":"ts","codigo":1234}`),
		}).Setup(infra.RequestMock{
			Url:                mappa + "/api/associados/1234",
			Method:             "GET",
			ResponseStatusCode: 200,
			ResponseBody:       []byte(`{}`),
		}).Setup(infra.RequestMock{
			Url:                mappa + "/api/grupos",
			Method:             "GET",
			ResponseStatusCode: 200,
			ResponseBody:       []byte(`[{"codigo": 32,"codigoRegiao": "SC","nome": "LEÕES DE BLUMENAU","codigoModalidade": 1}]`),
		}).Setup(infra.RequestMock{
			Url:                mappa + "/api/escotistas/1234/secoes",
			Method:             "GET",
			ResponseStatusCode: 200,
			ResponseBody:       []byte(`[{"codigo": 1,"nome": "Escoteiros","codigoModalidade": 1,"codigoGrupo": 32}]`),
		}).Setup(infra.RequestMock{
			Url:                mappa + "/api/escotistas/1234/secoes/1/equipes",
			Method:             "GET",
			ResponseStatusCode: 200,
			ResponseBody:       []byte(`[{"codigo": 1,"nome": "Escoteiros","codigoModalidade": 1,"codigoGrupo": 32}]`),
		})

		infra.ResetHttpRequester(requester)
		auth := &AuthData{
			UserName: "test",
			Password: "test",
		}
		api := NewMappaAPI()
		loginResponse, err := api.Login(auth.UserName, auth.Password)
		if err != nil || loginResponse == nil {
			t.Errorf("Login failed: %v", err)
		}
		if len(loginResponse.ID) == 0 || !loginResponse.IsValid() {
			t.Errorf("Login failed: %v", loginResponse)
		}

		escotista := api.GetEscotista(loginResponse.Userid, loginResponse.ID)
		if escotista == nil {
			t.Errorf("GetEscotista failed: %v", escotista)
		}

		associado := api.GetAssociado(escotista.CodigoAssociado, loginResponse.ID)
		if associado == nil {
			t.Errorf("GetAssociado failed: %v", associado)
		}

		grupo := api.GetGrupo(escotista.CodigoGrupo, escotista.CodigoRegiao, loginResponse.ID)
		if grupo == nil {
			t.Errorf("GetGrupo failed: %v", grupo)
		}

		secoes := api.GetEscotistaSecoes(escotista.UserId, loginResponse.ID)
		if len(secoes) == 0 {
			t.Errorf("GetEscotistaSecoes failed: %v", secoes)
		}

	})

}

func TestMappaAPI_MappaGetProgressoes(t *testing.T) {

	t.Run("Progressoes", func(t *testing.T) {
		requester := infra.NewHttpRequestMock()

		requester.Setup(infra.RequestMock{
			Url:                mappa + "/api/progressao-atividades",
			Method:             "GET",
			ResponseStatusCode: 200,
			ResponseBody:       []byte(`[{"ID":"1234","ttl":1209600,"created":"2099-01-01T00:00:00.000Z","userId":1234}]`),
			//TODO: Implementar resposta válida
		})
		infra.ResetHttpRequester(requester)
		api := NewMappaAPI()
		progressoes, err := api.GetProgressoes()
		if err != nil {
			t.Errorf("GetProgressoes failed: %v", err)
		}
		if len(progressoes) == 0 {
			t.Errorf("GetProgressoes failed: %v", progressoes)
		}
	})

}

func TestMappaAPI_MappaGetEspecialidades(t *testing.T) {

	t.Run("Especialidades", func(t *testing.T) {
		requester := infra.NewHttpRequestMock()

		requester.Setup(infra.RequestMock{
			Url:                mappa + "/api/especialidades",
			Method:             "GET",
			ResponseStatusCode: 200,
			ResponseBody:       []byte(`[{"ID":"1234","ttl":1209600,"created":"2099-01-01T00:00:00.000Z","userId":1234}]`),
			//TODO: Implementar resposta válida
		})
		infra.ResetHttpRequester(requester)
		api := NewMappaAPI()
		especialidades, err := api.GetEspecialidades()
		if err != nil {
			t.Errorf("GetEspecialidades failed: %v", err)
		}
		if len(especialidades) == 0 {
			t.Errorf("GetEspecialidades failed: %v", especialidades)
		}
	})
}
