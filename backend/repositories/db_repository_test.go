package repositories

import (
	"testing"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/types"

	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"gorm.io/gorm"
)

func TestDBRepository_IsValidConnectionString(t *testing.T) {
	type fields struct {
		BaseRepository   BaseRepository
		schema           string
		connectionString string
		getDBFunc        func() *gorm.DB
	}
	type args struct {
		connectionString string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DBRepository{
				BaseRepository:   tt.fields.BaseRepository,
				schema:           tt.fields.schema,
				connectionString: tt.fields.connectionString,
				getDBFunc:        tt.fields.getDBFunc,
			}
			if got := r.IsValidConnectionString(tt.args.connectionString); got != tt.want {
				t.Errorf("DBRepository.IsValidConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBRepository_CreateRepository(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		var r *DBRepository
		repository, err := r.CreateRepository("sqlite:./testing.db")
		if err != nil {
			t.Errorf("DBRepository.CreateRepository() error = %v", err)
		}
		if repository.GetName() != "DB: sqlite" {
			t.Errorf("DBRepository.CreateRepository() error = %v", err)
		}
		logins, err := repository.GetLogins()
		if err != nil {
			t.Errorf("DBRepository.GetLogins() error = %v", err)
		}
		for _, login := range logins {
			if err = repository.DeleteLogin(login.UserName); err != nil {
				t.Errorf("DBRepository.DeleteLogin() error = %v", err)
			}
		}
		if err = repository.SetLogin("test_user", "ABCD", responses.MappaLoginResponse{
			ID:      "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			TTL:     86400,
			Created: time.Now(),
			Userid:  1234,
		}, time.Date(2022, 5, 22, 0, 0, 0, 0, time.UTC)); err != nil {
			t.Errorf("DBRepository.SetLogin() error = %v", err)
		}
		if err = repository.SetLogin("old_user", "ABCD", responses.MappaLoginResponse{
			ID:      "1234567890",
			TTL:     86400,
			Created: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			Userid:  8888,
		}, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)); err != nil {
			t.Errorf("DBRepository.SetLogin() error = %v", err)
		}

		logins, err = repository.GetLogins()
		if err != nil {
			t.Errorf("DBRepository.GetLogins() error = %v", err)
		}
		if len(logins) != 1 {
			t.Errorf("DBRepository.GetLogins() expected only one login = %v", logins)
		}
		escotista := responses.MappaEscotistaResponse{
			UserId:          1234,
			CodigoAssociado: 2222,
			UserName:        "test_user",
			NomeCompleto:    "TEST USER",
			Ativo:           types.TrueBool,
			CodigoGrupo:     32,
			CodigoRegiao:    "SC",
			CodigoFoto:      0,
		}
		associado := responses.MappaAssociadoResponse{
			Codigo:                  2222,
			Nome:                    "TEST",
			CodigoFoto:              0,
			CodigoEquipe:            12,
			UserName:                1234,
			NumeroDigito:            5,
			DataNascimento:          types.Date(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)),
			DataValidade:            types.Date(time.Date(2022, 5, 22, 0, 0, 0, 0, time.UTC)),
			NomeAbreviado:           "TESTE",
			Sexo:                    "M",
			CodigoRamo:              1,
			CodigoCategoria:         0,
			CodigoSegundaCategoria:  0,
			CodigoTerceiraCategoria: 0,
			LinhaFormacao:           "TESTE",
			CodigoRamoAdulto:        0,
			DataAcompanhamento:      types.Date(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)),
		}
		grupo := responses.MappaGrupoResponse{
			Codigo:           32,
			CodigoRegiao:     "SC",
			Nome:             "GELB",
			CodigoModalidade: 1,
		}
		if err = repository.SetDetalheEscotista(7777, responses.MappaDetalhesResponse{
			Escotista: &escotista,
			Associado: &associado,
			Grupos:    []*responses.MappaGrupoResponse{&grupo},
		}); err != nil {
			t.Errorf("DBRepository.SetDetalheEscotista() error = %v", err)
		}

		detalhes, err := repository.GetDetalheEscotista(7777)
		if err != nil {
			t.Errorf("DBRepository.GetDetalheEscotista() error = %v", err)
		}
		if detalhes.Escotista.UserName != "test_user" {
			t.Errorf("DBRepository.GetDetalheEscotista() error = %v", err)
		}

	})

}
