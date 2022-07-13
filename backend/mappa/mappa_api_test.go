package mappa

import (
	"os"
	"testing"
)

type AuthData struct {
	UserName string
	Password string
}

var authData *AuthData

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
		auth := getAuthData(t)
		api := NewMappaAPI()
		loginResponse, err := api.Login(auth.UserName, auth.Password)
		if err != nil {
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
